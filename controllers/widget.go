package controllers

import (
	"df/api/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"regexp"
	"strings"
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
	garment := g.Find(p["id"])

	u.UpdateHistory(garment)

	r.HTML(http.StatusOK, this.Name, struct {
		Ids      []string
		User     *models.UserScheme
		Garments []models.GarmentScheme
		Version  int64
	}{
		[]string{p["id"]},
		u.Object,
		[]models.GarmentScheme{*garment},
		time.Now().Unix(),
	})
}

func (this *Widget) FindAll(opts models.URLOptionsScheme, u *models.User, g *models.Garment, r render.Render) {
	guid := regexp.MustCompile("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}")
	ids := []string{}

	for _, id := range strings.Split(opts.Ids, ",") {
		if guid.MatchString(id) {
			ids = append(ids, id)
		}
	}

	garments := []models.GarmentScheme{}
	if len(ids) > 0 {
		garments = g.FindAll(ids, opts)
	}

	r.HTML(http.StatusOK, this.Name, struct {
		Ids      []string
		User     *models.UserScheme
		Garments []models.GarmentScheme
		Version  int64
	}{
		ids,
		u.Object,
		garments,
		time.Now().Unix(),
	})
}

func (this *Widget) Index(r render.Render, g *models.Garment) {
	ids := make([]string, 0)
	result := g.FindAll(ids, models.URLOptionsScheme{Limit: 100})

	r.HTML(200, "index", struct{ Garments []models.GarmentScheme }{result})
}
