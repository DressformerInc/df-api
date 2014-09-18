package models

import (
	r "github.com/dancannon/gorethink"
	"log"
)

type DummyScheme struct {
	Id      string `gorethink:"id,omitempty"      json:"id,omitempty"   binding:"-"`
	Name    string `gorethink:"name,omitempty"    json:"name,omitempty"`
	Default bool   `gorethink:"default,omitempty" json:"default,omitempty"`

	Assets struct {
		Geometry Source `gorethink:"geometry,omitempty" json:"geometry,omitempty"`
	} `gorethink:"assets,omitempty" json:"assets,omitempty"`

	Body struct {
		Height    float64 `gorethink:"height,omitempty"    json:"height,omitempty"`
		Chest     float64 `gorethink:"chest,omitempty"     json:"chest,omitempty"`
		Underbust float64 `gorethink:"underbust,omitempty" json:"underbust,omitempty"`
		Waist     float64 `gorethink:"waist,omitempty"     json:"waist,omitempty"`
		Hips      float64 `gorethink:"hips,omitempty"      json:"hips,omitempty"`
	} `gorethink:"body,omitempty" json:"body,omitempty"`
}

type Dummy struct {
	*Base
}

func (*Dummy) Construct(args ...interface{}) interface{} {
	return &Dummy{
		&Base{r.Db("dressformer").Table("dummies")},
	}
}

func (this *Dummy) Find(id string) *DummyScheme {
	var query r.Term

	result := &DummyScheme{}

	if id == "" {
		query = this.GetAllByIndex("default", true)
	} else {
		query = this.Get(id)
	}

	rows, err := query.Run(session())
	if err != nil {
		log.Println("Unable to fetch cursor for id:", id, "Error:", err)
		return nil
	}

	if err = rows.One(&result); err != nil {
		log.Println("Unable to get data, err:", err)
		return nil
	}

	url(&result.Assets.Geometry, "geometry")

	return result
}

func (this *Dummy) Default() *DummyScheme {
	return this.Find("")
}

func (this *Dummy) FindAll(ids []string, opts URLOptionsScheme) []DummyScheme {
	var query r.Term

	if len(ids[0]) > 0 {
		query = this.GetAll(r.Args(ids))
	} else {
		query = this.Skip(opts.Start).Limit(opts.Limit)
	}

	rows, err := query.Run(session())

	if err != nil {
		log.Println("Unable to fetch cursor for args:", ids, "Error:", err)
		return nil
	}

	result := []DummyScheme{}

	if err = rows.All(&result); err != nil {
		log.Println("Unable to get data, err:", err)
	}

	for idx, _ := range result {
		url(&result[idx].Assets.Geometry, "geometry")
	}

	return result
}

func (this *Dummy) ResetDefault() {
	_, err := this.GetAllByIndex("default", true).Update(map[string]bool{"default": false}).Run(session())
	if err != nil {
		log.Println("Unable to update. Error:", err)
	}
}
