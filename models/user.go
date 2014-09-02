package models

import (
	r "github.com/dancannon/gorethink"
)

type UserScheme struct {
	Dummy string `gorethink:"dummy" json:"dummy"`

	Body struct {
		Height    float32 `gorethink:"height"    json:"height,omitempty"`
		Chest     float32 `gorethink:"chest"     json:"chest,omitempty"`
		Underbust float32 `gorethink:"underbust" json:"underbust,omitempty"`
		Waist     float32 `gorethink:"waist"     json:"waist,omitempty"`
		Hips      float32 `gorethink:"hips"      json:"hips,omitempty"`
	} `gorethink:"body,omitempty" json:"body,omitempty"`
}

type User struct {
	r.Term
	dummy *Dummy
}

func (*User) Construct(args ...interface{}) interface{} {
	return &User{
		r.Db("dressformer").Table("users"),
		(*Dummy).Construct(nil).(*Dummy),
	}
}

func (this *User) Find(args ...interface{}) *UserScheme {
	result := &UserScheme{
		Dummy: this.dummy.Find("").Assets.Geometry,
	}

	return result
}
