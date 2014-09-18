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
	if u.Object == nil {
		r.JSON(http.StatusOK, struct{}{})
		return
	}

	// Temp stuff. Fancy Filter is coming
	filter(u.Object)

	r.JSON(http.StatusOK, u.Object)
}

func filter(u *models.UserScheme) {
	if len(u.History) == 0 {
		return
	}

	for idx, _ := range u.History {
		models.Url(&u.History[idx].Assets.Geometry, "geometry")
		models.Url(&u.History[idx].Assets.Placeholder, "image")
		models.Url(&u.History[idx].Assets.Diffuse, "image")
		models.Url(&u.History[idx].Assets.Specular, "image")
		models.Url(&u.History[idx].Assets.Normal, "image")
	}
}
