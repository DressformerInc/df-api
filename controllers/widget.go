package controllers

import (
	"df/api/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type Widget struct {
	Name string
}

func (*Widget) Construct(args ...interface{}) interface{} {
	var name string
	if len(args) == 0 {
		name = "widget"
	} else {
		name = args[0].(string)
	}

	return &Widget{name}
}

func (this *Widget) Get(u *models.User, r render.Render, p martini.Params) {
	r.HTML(200, this.Name, struct {
		Id   string
		User *models.UserScheme
	}{
		p["id"],
		u.Object,
	})
}
