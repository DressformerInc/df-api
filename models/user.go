package models

import (
	"log"
)

type Model interface {
	Construct(arg ...interface{}) interface{}
}

type UserScheme struct {
	Name string `json:"name"`
}

type User struct {
	Object *UserScheme
}

func (*User) Construct(args ...interface{}) interface{} {
	this := &User{
		Object: &UserScheme{Name: "xxx"},
	}
	log.Println("User model initialized,", this)

	return this
}

func (this *User) Name() string {
	return this.Object.Name
}
