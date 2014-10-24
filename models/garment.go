package models

import (
	. "df/api/utils"
	r "github.com/dancannon/gorethink"
	"log"
)

type Size struct {
	Id       string `gorethink:"id,omitempty"        json:"id,omitempty"`
	SizeName string `gorethink:"size_name,omitempty" json:"size_name,omitempty"`
}

type GarmentScheme struct {
	Id        string  `gorethink:"id,omitempty"        json:"id"   binding:"-"`
	Gid       string  `gorethink:"gid,omitempty"       json:"gid,omitempty"`
	Name      string  `gorethink:"name,omitempty"      json:"name,omitempty"`
	SizeName  string  `gorethink:"size_name,omitempty" json:"size_name,omitempty"`
	Sizes     []Size  `gorethink:"sizes,omitempty"     json:"sizes,omitempty"`
	DummyId   string  `gorethink:"dummy_id,omitempty"  json:"dummy_id,omitempty"`
	Slot      string  `gorethink:"slot,omitempty"      json:"slot,omitempty"`
	Layer     float64 `gorethink:"layer,omitempty"     json:"layer,omitempty"`
	UrlPrefix string  `gorethink:"-"                   json:"url_prefix,omitempty"`

	Assets struct {
		Geometry    Source `gorethink:"geometry,omitempty"    json:"geometry,omitempty"`
		Placeholder Source `gorethink:"placeholder,omitempty" json:"placeholder,omitempty"`
		MtlSrc      Source `gorethink:"mtl_src,omitempty"     json:"mtl_src,omitempty"`
	} `gorethink:"assets,omitempty" json:"assets,omitempty"`

	Materials interface{} `gorethink:"materials,omitempty"    json:"materials,omitempty"`
}

type Garment struct {
	*Base
	Dummy    *Dummy
	Material *Material
}

func (*Garment) Construct(args ...interface{}) interface{} {
	return &Garment{
		&Base{r.Db("dressformer").Table("garments")},
		(*Dummy).Construct(nil).(*Dummy),
		(*Material).Construct(nil).(*Material),
	}
}

func (this *Garment) Find(id string) *GarmentScheme {
	result := &GarmentScheme{
		UrlPrefix: AppConfig.AssetsUrl() + "/",
	}

	rows, err := this.Get(id).Run(session())
	if err != nil {
		log.Println("Unable to fetch cursor for id:", id, "Error:", err)
		return nil
	}

	if err = rows.One(&result); err != nil {
		log.Println("Unable to get data, err:", err)
		return nil
	}

	return this.Expand(result)
}

func (this *Garment) Expand(result *GarmentScheme) *GarmentScheme {
	var items []string

	switch result.Materials.(type) {
	case []interface{}:
		for idx, _ := range result.Materials.([]interface{}) {
			items = append(items, result.Materials.([]interface{})[idx].(string))
		}

	default:
		return result
	}

	result.Materials = this.Material.FindAll(items, URLOptionsScheme{})

	return result
}

func (this *Garment) FindAll(ids []string, opts URLOptionsScheme) []GarmentScheme {
	i, err := this.Base.FindAll(ids, opts, &[]GarmentScheme{})
	if err != nil {
		return nil
	}

	result := *(i.(*[]GarmentScheme))

	for idx, _ := range result {
		(&result[idx]).UrlPrefix = AppConfig.AssetsUrl() + "/"
		result[idx] = *this.Expand(&result[idx])
	}

	if result == nil {
		result = []GarmentScheme{}
	}

	return result
}
