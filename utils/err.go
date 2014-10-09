package utils

import (
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"net/http"
)

type ErrorScheme struct {
	Items binding.Errors `json:"errors"`
}

func Err(s string) map[string]string {
	return map[string]string{"error_msg": s}
}

func ErrorHandler(errs binding.Errors, w http.ResponseWriter, r render.Render) {
	if len(errs) == 0 {
		return
	}

	var status int

	if errs.Has(binding.DeserializationError) {
		status = http.StatusBadRequest
	} else if errs.Has(binding.ContentTypeError) {
		status = http.StatusUnsupportedMediaType
	} else {
		status = http.StatusInternalServerError
	}

	r.JSON(status, ErrorScheme{Items: errs})
}
