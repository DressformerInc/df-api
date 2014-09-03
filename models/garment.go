package models

import (
	"code.google.com/p/go-uuid/uuid"
	"errors"
	r "github.com/dancannon/gorethink"
	"log"
)

type Size struct {
	Id       string `gorethink:"id,omitempty"        json:"id,omitempty"`
	SizeName string `gorethink:"size_name,omitempty" json:"size_name,omitempty"`
}

type GarmentScheme struct {
	Id       string `gorethink:"id,omitempty"        json:"id"   binding:"-"`
	Gid      string `gorethink:"gid,omitempty"       json:"gid,omitempty"`
	Name     string `gorethink:"name,omitempty"      json:"name,omitempty"`
	SizeName string `gorethink:"size_name,omitempty" json:"size_name,omitempty"`
	Sizes    []Size `gorethink:"sizes,omitempty"     json:"sizes,,omitempty"`
	DummyId  string `gorethink:"dummy_id,omitempty"  json:"dummy_id,omitempty"`

	Assets struct {
		Geometry Source `gorethink:"geometry,omitempty" json:"geometry,omitempty"`
		Diffuse  Source `gorethink:"diffuse,omitempty"  json:"diffuse,omitempty"`
		Normal   Source `gorethink:"normal,omitempty"   json:"normal,omitempty"`
		Specular Source `gorethink:"specular,omitempty" json:"specular,omitempty"`
	} `gorethink:"assets,omitempty" json:"assets,omitempty"`
}

type Garment struct {
	r.Term
	dummy *Dummy
}

func (*Garment) Construct(args ...interface{}) interface{} {
	return &Garment{
		r.Db("dressformer").Table("garments"),
		(*Dummy).Construct(nil).(*Dummy),
	}
}

// obsolete
func (this *Garment) Find(id string) *GarmentScheme {
	var result *GarmentScheme

	rows, err := this.Get(id).Run(session())
	if err != nil {
		log.Println("Unable to fetch cursor for id:", id, "Error:", err)
		return nil
	}

	if err = rows.One(&result); err != nil {
		log.Println("Unable to get data, err:", err)
		return nil
	}

	url(&result.Assets.Geometry, "geometry")
	url(&result.Assets.Diffuse, "image")
	url(&result.Assets.Normal, "image")
	url(&result.Assets.Specular, "image")

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

	for idx, _ := range result {
		url(&result[idx].Assets.Geometry, "geometry")
		url(&result[idx].Assets.Diffuse, "image")
		url(&result[idx].Assets.Normal, "image")
		url(&result[idx].Assets.Specular, "image")
	}

	return result
}

func (this *Garment) Create(payload GarmentScheme) (*GarmentScheme, error) {
	if payload.Gid == "" {
		payload.Gid = uuid.New()
	}

	if payload.DummyId == "" {
		payload.DummyId = this.dummy.Default().Id
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
