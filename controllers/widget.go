package controllers

import (
	"df/api/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"time"
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

func (this *Widget) Get(u *models.User, g *models.Garment, r render.Render, p martini.Params) {
	r.HTML(http.StatusOK, this.Name, struct {
		Id      string
		User    *models.UserScheme
		Garment *models.GarmentScheme
		Version int64
	}{
		p["id"],
		u.Object,
		g.Find(p["id"]),
		time.Now().Unix(),
	})
}

func (this *Widget) Index(r render.Render, g *models.Garment) {
	ids := make([]string, 0)
	result := g.FindAll(ids, models.URLOptionsScheme{Limit: 100})

	r.HTML(200, "index", struct{ Garments []models.GarmentScheme }{result})
}
