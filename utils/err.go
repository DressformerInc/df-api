package utils

import (
	"github.com/3d0c/binding"
	"github.com/martini-contrib/encoder"
	"net/http"
)

type ErrorScheme struct {
	Items binding.Errors `json:"errors"`
}

func Err(s string) map[string]string {
	return map[string]string{"error_msg": s}
}

func ErrorHandler(errs binding.Errors, w http.ResponseWriter, enc encoder.Encoder) {
	if len(errs) == 0 {
		return
	}

	if errs.Has(binding.DeserializationError) {
		w.WriteHeader(http.StatusBadRequest)
	} else if errs.Has(binding.ContentTypeError) {
		w.WriteHeader(http.StatusUnsupportedMediaType)
	} else {
		w.WriteHeader(binding.StatusUnprocessableEntity)
	}

	e := ErrorScheme{Items: errs}

	w.Write(encoder.Must(enc.Encode(e)))
}
