package models

import (
	r "github.com/dancannon/gorethink"
)

type UserScheme struct {
	Dummy *DummyScheme `json:"dummy,omitempty"`
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
	user := &UserScheme{}

	if this.dummy != nil {
		user.Dummy = this.dummy.Default()
	}

	return user
}
