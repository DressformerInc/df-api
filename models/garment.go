package models

import (
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
	Id       string `gorethink:"id,omitempty" json:"id"   binding:"-"`
	Name     string `gorethink:"name"         json:"name"`
	SizeName string `gorethink:"size_name"    json:"size"`
	Sizes    []Size `gorethink:"sizes"        json:"sizes"`

	Assets struct {
		Geometry string `gorethink:"geometry" json:"geometry"`
		Diffuse  string `gorethink:"diffuse"  json:"diffuse"`
		Normal   string `gorethink:"normal"   json:"normal"`
	} `gorethink:"assets" json:"assets"`

	Sources [][]Source `gorethink:"sources" json:"sources,omitempty"`
}

type Garment struct {
	r.Term
}

func (*Garment) Construct(args ...interface{}) interface{} {
	return &Garment{
		r.Db("dressformer"),
	}
}

func (this *Garment) Find(id interface{}) *GarmentScheme {
	var query r.Term

	switch t := id.(type) {
	case oid.ObjectId:
		query = r.Table("garments").GetAllByIndex("assetsGeometry", id.(oid.ObjectId).String())

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
	query := this.Table("garments")

	if len(ids[0]) > 0 {
		log.Println("!empty, ", len(ids), ",:", ids)
		query = query.GetAll(r.Args(ids))
	} else {
		query.Skip(opts.Skip).Limit(50)
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
	result, err := r.Table("garments").Insert(payload, r.InsertOpts{ReturnVals: true}).Run(session())
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
