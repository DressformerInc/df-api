package controllers

import (
	"df/api/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"regexp"
	"strings"
)

type Dummy struct {
	model *models.Dummy
}

func (*Dummy) Construct(args ...interface{}) interface{} {
	return &Dummy{
		model: (*models.Dummy).Construct(nil).(*models.Dummy),
	}
}

func (this *Dummy) Find(u *models.User, params martini.Params, r render.Render) {
	result := this.model.Find(params["id"])
	if result == nil {
		r.JSON(http.StatusOK, struct{}{})
		return
	}

	r.JSON(http.StatusOK, result)
}

func (this *Dummy) FindAll(opts models.URLOptionsScheme, u *models.User, r render.Render) {
	guid := regexp.MustCompile("\b[A-F0-9]{8}(?:-[A-F0-9]{4}){3}-[A-F0-9]{12}\b")
	ids := make([]string, 0)

	for _, id := range strings.Split(opts.Ids, ",") {
		if guid.MatchString(id) {
			ids = append(ids, id)
		}
	}

	if opts.Limit == 0 || opts.Limit > 100 {
		opts.Limit = 25
	}

	result := this.model.FindAll(ids, opts)

	r.JSON(http.StatusOK, result)
}

func (this *Dummy) Create(u *models.User, payload models.DummyScheme, r render.Render) {
	if payload.Default {
		this.model.ResetDefault()
	}

	result, err := this.model.Create(&payload)
	if err != nil {
		r.JSON(http.StatusBadRequest, []byte{})
		return
	}

	r.JSON(http.StatusOK, result)
}

func (this *Dummy) Put(u *models.User, payload models.DummyScheme, r render.Render, p martini.Params) {

	if payload.Default {
		this.model.ResetDefault()
	}

	result, err := this.model.Put(p["id"], &payload)
	if err != nil {
		r.JSON(http.StatusBadRequest, []byte{})
		return
	}

	r.JSON(http.StatusOK, result)
}

func (this *Dummy) Remove(u *models.User, r render.Render, p martini.Params) {
	err := this.model.Remove(p["id"])
	if err != nil {
		r.JSON(http.StatusBadRequest, []byte{})
		return
	}

	r.JSON(http.StatusOK, struct{}{})
}
