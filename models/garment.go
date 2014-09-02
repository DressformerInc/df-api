package models

import (
	"code.google.com/p/go-uuid/uuid"
	"errors"
	"github.com/3d0c/oid"
	r "github.com/dancannon/gorethink"
	"log"
)

type Source struct {
	Id     string  `gorethink:"id" json:"id"`
	Weight float32 `gorethink:"id" json:"weight"`
}

type Size struct {
	Id       string `gorethink:"id"        json:"id"`
	SizeName string `gorethink:"size_name" json:"size_name"`
}

type GarmentScheme struct {
	Id       string `gorethink:"id,omitempty"    json:"id"   binding:"-"`
	Gid      string `gorethink:"gid"             json:"gid"`
	Name     string `gorethink:"name"            json:"name"`
	SizeName string `gorethink:"size_name"       json:"size_name"`
	Sizes    []Size `gorethink:"sizes"           json:"sizes"`
	Dummy    string `gorethink:"dummy,omitempty" json:"dummy,omitempty"`

	Assets struct {
		Geometry string `gorethink:"geometry" json:"geometry"`
		Diffuse  string `gorethink:"diffuse"  json:"diffuse"`
		Normal   string `gorethink:"normal"   json:"normal"`
		Specular string `gorethink:"specular" json:"specular"`
	} `gorethink:"assets" json:"assets"`
}

type Garment struct {
	r.Term
}

func (*Garment) Construct(args ...interface{}) interface{} {
	return &Garment{
		r.Db("dressformer").Table("garments"),
	}
}

// obsolete
func (this *Garment) Find(id interface{}) *GarmentScheme {
	var query r.Term

	switch t := id.(type) {
	case oid.ObjectId:
		query = this.GetAllByIndex("assetsGeometry", id.(oid.ObjectId).String())

	default:
		log.Println("Unexpected type:", t)
		return nil
	}

	rows, err := query.Run(session())
	if err != nil {
		log.Println("Unable to fetch cursor for id:", id, "Error:", err)
		return nil
	}

	var result *GarmentScheme

	if err = rows.One(&result); err != nil {
		log.Println("Unable to get data, err:", err)
		return nil
	}

	return result
}

func (this *Garment) FindAll(ids []string, opts URLOptionsScheme) []GarmentScheme {
	var query r.Term

	if len(ids[0]) > 0 {
		log.Println("!empty, ", len(ids), ",:", ids)
		query = this.GetAll(r.Args(ids))
	} else {
		query = this.Skip(opts.Start).Limit(opts.Limit)
	}

	rows, err := query.Run(session())

	if err != nil {
		log.Println("Unable to fetch cursor for args:", ids, "Error:", err)
		return nil
	}

	var result []GarmentScheme

	if err = rows.All(&result); err != nil {
		log.Println("Unable to get data, err:", err)
	}

	return result
}

func (this *Garment) Create(payload GarmentScheme) (*GarmentScheme, error) {
	if payload.Gid == "" {
		payload.Gid = uuid.New()
	}

	result, err := this.Insert(payload, r.InsertOpts{ReturnVals: true}).Run(session())
	if err != nil {
		log.Println("Error inserting data:", err)
		return nil, errors.New("Internal server error")
	}

	response := &r.WriteResponse{NewValue: &GarmentScheme{}}

	if err = result.One(response); err != nil {
		log.Println("Unable to iterate cursor:", err)
		return nil, errors.New("Internal server error")
	}

	log.Println("inserted :", response.Inserted)
	log.Println("new_val:", response.NewValue)

	return response.NewValue.(*GarmentScheme), nil
}

func (this *Garment) Put(id string, payload GarmentScheme) (*GarmentScheme, error) {
	result, err := this.Get(id).Update(payload, r.UpdateOpts{ReturnVals: true}).Run(session())
	if err != nil {
		log.Println("Error updating:", id, "with data:", payload, "error:", err)
		return nil, errors.New("Wrong data")
	}

	response := &r.WriteResponse{NewValue: &GarmentScheme{}}

	if err = result.One(response); err != nil {
		log.Println("Unable to iterate cursor:", err)
		return nil, errors.New("Internal server error")
	}

	log.Println("new_val:", response.NewValue)

	return response.NewValue.(*GarmentScheme), nil
}

func (this *Garment) Remove(id string) error {
	_, err := this.Get(id).Delete().Run(session())
	if err != nil {
		log.Println("Error deleting:", id, "error:", err)
		return errors.New("Internal server error")
	}

	return nil
}
