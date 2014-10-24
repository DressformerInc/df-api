package controllers

import (
	"code.google.com/p/go-uuid/uuid"
	"df/api/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"regexp"
	"strings"
)

type Garment struct {
	model *models.Garment
}

func (*Garment) Construct(args ...interface{}) interface{} {
	return &Garment{
		model: (*models.Garment).Construct(nil).(*models.Garment),
	}
}

func (this *Garment) Find(u *models.User, r render.Render, params martini.Params) {
	result := this.model.Find(params["id"])
	if result == nil {
		r.JSON(http.StatusOK, struct{}{})
		return
	}

	u.UpdateHistory(result)

	r.JSON(http.StatusOK, result)
}

func (this *Garment) FindAll(opts models.URLOptionsScheme, u *models.User, r render.Render) {
	guid := regexp.MustCompile("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}")
	ids := []string{}

	for _, id := range strings.Split(opts.Ids, ",") {
		if guid.MatchString(id) {
			ids = append(ids, id)
		}
	}

	result := this.model.FindAll(ids, opts)
	u.UpdateHistory(result)

	r.JSON(http.StatusOK, result)
}

func (this *Garment) Create(u *models.User, payload models.GarmentScheme, r render.Render) {
	if payload.Gid == "" {
		payload.Gid = uuid.New()
	}

	if payload.DummyId == "" {
		payload.DummyId = this.model.Dummy.Default().Id
	}

	result, err := this.model.Create(&payload)
	if err != nil {
		r.JSON(http.StatusBadRequest, []byte{})
		return
	}

	r.JSON(http.StatusOK, result)
}

func (this *Garment) Put(u *models.User, payload models.GarmentScheme, r render.Render, p martini.Params) {

	result, err := this.model.Put(p["id"], &payload)
	if err != nil {
		r.JSON(http.StatusBadRequest, []byte{})
		return
	}

	r.JSON(http.StatusOK, result)
}

func (this *Garment) Remove(u *models.User, r render.Render, p martini.Params) {
	err := this.model.Remove(p["id"])
	if err != nil {
		r.JSON(http.StatusBadRequest, []byte{})
		return
	}

	r.JSON(http.StatusOK, []byte{})
}
