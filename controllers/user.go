package controllers

import (
	"df/api/models"
	"github.com/martini-contrib/render"
	"net/http"
)

type User struct {
}

func (*User) Construct(args ...interface{}) interface{} {
	this := &User{}
	return this
}

func (this *User) Find(u *models.User, r render.Render) {
	r.JSON(http.StatusOK, u.Object)
}
