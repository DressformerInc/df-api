package controllers

import (
	"df/api/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"regexp"
	"strings"
)

type Material struct {
	model *models.Material
}

func (*Material) Construct(args ...interface{}) interface{} {
	return &Material{
		model: (*models.Material).Construct(nil).(*models.Material),
	}
}

func (this *Material) Find(u *models.User, r render.Render, params martini.Params) {
	result := this.model.Find(params["id"])
	if result == nil {
		r.JSON(http.StatusOK, struct{}{})
		return
	}

	r.JSON(http.StatusOK, result)
}

func (this *Material) FindAll(opts models.URLOptionsScheme, u *models.User, r render.Render) {
	guid := regexp.MustCompile("\b[A-F0-9]{8}(?:-[A-F0-9]{4}){3}-[A-F0-9]{12}\b")
	ids := []string{}

	for _, id := range strings.Split(opts.Ids, ",") {
		if guid.MatchString(id) {
			ids = append(ids, id)
		}
	}

	if opts.Limit == 0 || opts.Limit > 100 {
		opts.Limit = 25
	}

	result := this.model.FindAll(ids, opts)

	// @todo
	// see T() - always create empty instance
	if result == nil {
		result = []models.MaterialScheme{}
	}

	r.JSON(http.StatusOK, result)
}

func (this *Material) Create(u *models.User, payload []models.MaterialScheme, r render.Render) {
	result, err := this.model.Create(payload)
	if err != nil {
		r.JSON(http.StatusBadRequest, []byte{})
		return
	}

	r.JSON(http.StatusOK, result)
}

func (this *Material) Put(u *models.User, payload models.MaterialScheme, r render.Render, p martini.Params) {

	result, err := this.model.Put(p["id"], &payload)
	if err != nil {
		r.JSON(http.StatusBadRequest, []byte{})
		return
	}

	r.JSON(http.StatusOK, result)
}
