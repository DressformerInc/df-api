package models

import (
	"errors"
	r "github.com/dancannon/gorethink"
	"github.com/martini-contrib/binding"
	"log"
	"net/http"
)

type Pair map[string]interface{}

type GarmentScheme struct {
	Id       string `gorethink:"id,omitempty" json:"id"   binding:"-"`
	Name     string `gorethink:"name"         json:"name"`
	SizeName string `gorethink:"size"         json:"size"`
	Sizes    []Pair `gorethink:"sizes"        json:"sizes"`

	Assets struct {
		Geometry string `gorethink:"geometry" json:"geometry"`
		Diffuse  string `gorethink:"diffuse"  json:"diffuse"`
		Normal   string `gorethink:"normal"   json:"normal"`
	}

	Sources [][]Pair `gorethink:"sources" json:"sources, omitempty"`
}

func (this GarmentScheme) Validate(errors *binding.Errors, req *http.Request) {
	// E.g.:
	// if len(this.Title) == 0 {
	// 	errors.Fields["title"] = "Title can't be empty."
	// }
}

type Garment struct{}

func (*Garment) Construct(args ...interface{}) interface{} {
	this := &Garment{}
	log.Println("garment model:", this)
	return this
}

func (this *Garment) FindAll(ids interface{}) []GarmentScheme {
	rows, err := r.Table("garments").GetAll(r.Args(ids)).Run(session())
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

	log.Println("inserted:", response.Inserted)
	log.Println("new_val:", response.NewValue)

	return response.NewValue.(*GarmentScheme), nil
}
