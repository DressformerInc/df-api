package models

import (
	"errors"
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
	r.Term
}

func (*Dummy) Construct(args ...interface{}) interface{} {
	return &Dummy{
		r.Db("dressformer").Table("dummies"),
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

	result := []DummyScheme{}

	if err = rows.All(&result); err != nil {
		log.Println("Unable to get data, err:", err)
	}

	for idx, _ := range result {
		url(&result[idx].Assets.Geometry, "geometry")
	}

	return result
}

func (this *Dummy) Create(payload DummyScheme) (*DummyScheme, error) {
	if payload.Default {
		this.ResetDefault()
	}

	result, err := this.Insert(payload, r.InsertOpts{ReturnVals: true}).Run(session())
	if err != nil {
		log.Println("Error inserting data:", err)
		return nil, errors.New("Internal server error")
	}

	response := &r.WriteResponse{NewValue: &DummyScheme{}}

	if err = result.One(response); err != nil {
		log.Println("Unable to iterate cursor:", err)
		return nil, errors.New("Internal server error")
	}

	log.Println("inserted :", response.Inserted)
	log.Println("new_val:", response.NewValue)

	return response.NewValue.(*DummyScheme), nil
}

func (this *Dummy) Put(id string, payload DummyScheme) (*DummyScheme, error) {
	if payload.Default {
		this.ResetDefault()
	}

	result, err := this.Get(id).Update(payload, r.UpdateOpts{ReturnVals: true}).Run(session())
	if err != nil {
		log.Println("Error updating:", id, "with data:", payload, "error:", err)
		return nil, errors.New("Wrong data")
	}

	response := &r.WriteResponse{NewValue: &DummyScheme{}}

	if err = result.One(response); err != nil {
		log.Println("Unable to iterate cursor:", err)
		return nil, errors.New("Internal server error")
	}

	if response.NewValue == nil {
		return nil, errors.New("Wrong data")
	}

	return response.NewValue.(*DummyScheme), nil
}

func (this *Dummy) Remove(id string) error {
	_, err := this.Get(id).Delete().Run(session())
	if err != nil {
		log.Println("Error deleting:", id, "error:", err)
		return errors.New("Internal server error")
	}

	return nil
}

func (this *Dummy) ResetDefault() {
	_, err := this.GetAllByIndex("default", true).Update(map[string]bool{"default": false}).Run(session())
	if err != nil {
		log.Println("Unable to update. Error:", err)
	}
}
