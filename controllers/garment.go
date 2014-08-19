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
	this := &Garment{
		model: (*models.Garment).Construct(nil).(*models.Garment),
	}

	return this
}

func (this *Garment) Find(u *models.User, enc encoder.Encoder) (int, []byte) {
	return http.StatusOK, []byte{}
}

func (this *Garment) FindAll(opts models.URLOptionsScheme, u *models.User, enc encoder.Encoder) (int, []byte) {
	guid := regexp.MustCompile("\b[A-F0-9]{8}(?:-[A-F0-9]{4}){3}-[A-F0-9]{12}\b")
	ids := make([]interface{}, 0)

	for _, id := range strings.Split(opts.Ids, ",") {
		// validate uuid here
		if !guid.MatchString(id) {
			ids = append(ids, id)
		} else {
			log.Println("Wrong GUID in:", id)
		}
	}

	if result := this.model.FindAll(ids); result != nil {
		return http.StatusOK, encoder.Must(enc.Encode(result))
	}

	return http.StatusNotFound, []byte{}
}
