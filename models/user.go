package models

import (
	r "github.com/dancannon/gorethink"
	// "log"
)

type UserScheme struct {
	Dummy DummyScheme `json:"dummy"`
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
		Dummy: *this.dummy.Default(),
	}

	return result
}
