package models

import (
	r "github.com/dancannon/gorethink"
	"log"
)

type Size struct {
	Id       string `gorethink:"id,omitempty"        json:"id,omitempty"`
	SizeName string `gorethink:"size_name,omitempty" json:"size_name,omitempty"`
}

type GarmentScheme struct {
	Id       string  `gorethink:"id,omitempty"        json:"id"   binding:"-"`
	Gid      string  `gorethink:"gid,omitempty"       json:"gid,omitempty"`
	Name     string  `gorethink:"name,omitempty"      json:"name,omitempty"`
	SizeName string  `gorethink:"size_name,omitempty" json:"size_name,omitempty"`
	Sizes    []Size  `gorethink:"sizes,omitempty"     json:"sizes,omitempty"`
	DummyId  string  `gorethink:"dummy_id,omitempty"  json:"dummy_id,omitempty"`
	Slot     string  `gorethink:"slot,omitempty"      json:"slot,omitempty"`
	Layer    float64 `gorethink:"layer,omitempty"     json:"layer,omitempty"`

	Assets struct {
		Geometry    Source `gorethink:"geometry,omitempty"    json:"geometry,omitempty"`
		Mtl         Source `gorethink:"mtl,omitempty"         json:"mtl,omitempty"`
		Diffuse     Source `gorethink:"diffuse,omitempty"     json:"diffuse,omitempty"`
		Normal      Source `gorethink:"normal,omitempty"      json:"normal,omitempty"`
		Specular    Source `gorethink:"specular,omitempty"    json:"specular,omitempty"`
		Placeholder Source `gorethink:"placeholder,omitempty" json:"placeholder,omitempty"`
	} `gorethink:"assets,omitempty" json:"assets,omitempty"`
}

type Garment struct {
	*Base
	Dummy *Dummy
}

func (*Garment) Construct(args ...interface{}) interface{} {
	return &Garment{
		&Base{r.Db("dressformer").Table("garments")},
		(*Dummy).Construct(nil).(*Dummy),
	}
}

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
	url(&result.Assets.Placeholder, "image")
	url(&result.Assets.Mtl, "")

	return result
}

func (this *Garment) FindAll(ids []string, opts URLOptionsScheme) []GarmentScheme {
	i, err := this.Base.FindAll(ids, opts, &[]GarmentScheme{})
	if err != nil {
		return nil
	}

	result := *(i.(*[]GarmentScheme))

	for idx, _ := range result {
		url(&result[idx].Assets.Geometry, "geometry")
		url(&result[idx].Assets.Diffuse, "image")
		url(&result[idx].Assets.Normal, "image")
		url(&result[idx].Assets.Specular, "image")
		url(&result[idx].Assets.Placeholder, "image")
		url(&result[idx].Assets.Mtl, "")
	}

	return result
}
