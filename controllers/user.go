package controllers

import (
	"df/api/models"
	"github.com/martini-contrib/encoder"
	"net/http"
)

type User struct {
}

func (*User) Construct(args ...interface{}) interface{} {
	this := &User{}
	return this
}

func (this *User) Find(u *models.User, enc encoder.Encoder) (int, []byte) {
	return http.StatusOK, encoder.Must(enc.Encode(u.Object))
}
