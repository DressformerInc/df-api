package controllers

import (
	"df/api/models"
	"github.com/martini-contrib/encoder"
	"log"
	"net/http"
)

type Controller interface {
	Construct(arg ...interface{}) interface{}
}

type User struct {
}

func (*User) Construct(args ...interface{}) interface{} {
	this := &User{}
	log.Println("User controller initialized,", this)

	return this
}

func (this *User) Find(u *models.User, enc encoder.Encoder) (int, []byte) {
	return http.StatusOK, encoder.Must(enc.Encode(u.Object))
}
