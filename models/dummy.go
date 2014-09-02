package models

import (
	"errors"
	r "github.com/dancannon/gorethink"
	"log"
)

type DummyScheme struct {
	Id      string `gorethink:"id,omitempty" json:"id"   binding:"-"`
	Name    string `gorethink:"name"         json:"name"`
	Default bool   `gorethink:"default"      json:"default"`

	Assets struct {
		Geometry string `gorethink:"geometry" json:"geometry"`
	} `gorethink:"assets" json:"assets"`

	Body struct {
		Height    float32 `gorethink:"height"    json:"height,omitempty"`
		Chest     float32 `gorethink:"chest"     json:"chest,omitempty"`
		Underbust float32 `gorethink:"underbust" json:"underbust,omitempty"`
		Waist     float32 `gorethink:"waist"     json:"waist,omitempty"`
		Hips      float32 `gorethink:"hips"      json:"hips,omitempty"`
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
	rows, err := this.Get(id).Run(session())
	if err != nil {
		log.Println("Unable to fetch cursor for id:", id, "Error:", err)
		return nil
	}

	var result *DummyScheme

	if err = rows.One(&result); err != nil {
		log.Println("Unable to get data, err:", err)
		return nil
	}

	return result
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

	var result []DummyScheme

	if err = rows.All(&result); err != nil {
		log.Println("Unable to get data, err:", err)
	}

	return result
}

func (this *Dummy) Create(payload DummyScheme) (*DummyScheme, error) {
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

	log.Println("new_val:", response.NewValue)

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
