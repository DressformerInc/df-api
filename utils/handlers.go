package utils

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/go-martini/martini"
	"github.com/gorilla/securecookie"
	"log"
	"net/http"
)

func CorsHandler(w http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		w.Header().Add("Access-Control-Allow-Origin", origin)
	} else {
		w.Header().Add("Access-Control-Allow-Origin", "*")
	}

	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Requested-With")
	w.Header().Add("Cache-Control", "max-age=2592000")
	w.Header().Add("Pragma", "public")
	w.Header().Add("Cache-Control", "public")
}

func LogHandler(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.RequestURI)
}

func TokenHandler(w http.ResponseWriter, r *http.Request, sc *securecookie.SecureCookie, c martini.Context) {
	this := Token{IsRestored: false}

	if cookie, err := r.Cookie("df-token"); err != nil {
		this.payload = uuid.New()
		encoded, err := sc.Encode("df-token", this.payload)
		if err != nil {
			log.Println("Encode error:", err)
			return
		}

		cookie := &http.Cookie{
			Name:  "df-token",
			Value: encoded,
			Path:  "/",
		}

		http.SetCookie(w, cookie)
	} else {
		if err := sc.Decode("df-token", cookie.Value, &this.payload); err == nil {
			this.IsRestored = true
		}
	}

	c.Map(this)
}

type Token struct {
	payload    string
	IsRestored bool
}

func (this Token) Get() string {
	return this.payload
}
