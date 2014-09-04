package controllers

import (
	"df/api/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/encoder"
	"log"
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

func (this *Dummy) Find(u *models.User, enc encoder.Encoder, params martini.Params) (int, []byte) {
	result := this.model.Find(params["id"])
	if result == nil {
		return http.StatusOK, encoder.Must(enc.Encode(struct{}{}))
	}

	return http.StatusOK, encoder.Must(enc.Encode(result))
}

func (this *Dummy) FindAll(opts models.URLOptionsScheme, u *models.User, enc encoder.Encoder, r *http.Request) (int, []byte) {
	log.Println(r.RequestURI)
	log.Println("opts:", opts)
	guid := regexp.MustCompile("\b[A-F0-9]{8}(?:-[A-F0-9]{4}){3}-[A-F0-9]{12}\b")
	ids := make([]string, 0)

	for _, id := range strings.Split(opts.Ids, ",") {
		if !guid.MatchString(id) {
			ids = append(ids, id)
		} else {
			log.Println("Wrong GUID in:", id)
		}
	}

	if opts.Limit == 0 || opts.Limit > 100 {
		opts.Limit = 25
	}

	result := this.model.FindAll(ids, opts)

	return http.StatusOK, encoder.Must(enc.Encode(result))
}

func (this *Dummy) Create(u *models.User, payload models.DummyScheme, enc encoder.Encoder) (int, []byte) {

	result, err := this.model.Create(payload)
	if err != nil {
		return http.StatusBadRequest, []byte{}
	}

	return http.StatusOK, encoder.Must(enc.Encode(result))
}

func (this *Dummy) Put(u *models.User, payload models.DummyScheme, enc encoder.Encoder, p martini.Params) (int, []byte) {

	result, err := this.model.Put(p["id"], payload)
	if err != nil {
		return http.StatusBadRequest, []byte{}
	}

	return http.StatusOK, encoder.Must(enc.Encode(result))
}

func (this *Dummy) Remove(u *models.User, enc encoder.Encoder, p martini.Params) (int, []byte) {
	err := this.model.Remove(p["id"])
	if err != nil {
		return http.StatusBadRequest, []byte{}
	}

	return http.StatusOK, []byte{}
}
