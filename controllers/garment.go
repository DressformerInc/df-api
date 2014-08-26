package controllers

import (
	"df/api/models"
	"github.com/martini-contrib/encoder"
	"log"
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

func (this *Garment) Find(u *models.User, enc encoder.Encoder) (int, []byte) {
	return http.StatusOK, []byte{}
}

func (this *Garment) FindAll(opts models.URLOptionsScheme, u *models.User, enc encoder.Encoder) (int, []byte) {
	guid := regexp.MustCompile("\b[A-F0-9]{8}(?:-[A-F0-9]{4}){3}-[A-F0-9]{12}\b")
	ids := make([]string, 0)

	for _, id := range strings.Split(opts.Ids, ",") {
		if !guid.MatchString(id) {
			ids = append(ids, id)
		} else {
			log.Println("Wrong GUID in:", id)
		}
	}

	if opts.Skip == 0 {
		opts.Skip = 30
	}

	if result := this.model.FindAll(ids, opts); result != nil {
		return http.StatusOK, encoder.Must(enc.Encode(result))
	}

	return http.StatusNotFound, []byte{}
}

func (this *Garment) Create(u *models.User, payload models.GarmentScheme, enc encoder.Encoder) (int, []byte) {

	result, err := this.model.Create(payload)
	if err != nil {
		return http.StatusBadRequest, []byte{}
	}

	return http.StatusOK, encoder.Must(enc.Encode(result))
}
