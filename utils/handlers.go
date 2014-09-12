package utils

import (
	"github.com/go-martini/martini"
	"github.com/gorilla/securecookie"
	uuid "github.com/nu7hatch/gouuid"
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
	this := Token{}

	if cookie, err := r.Cookie("df-token"); err != nil {
		if u4, err := uuid.NewV4(); err != nil {
			log.Println("Unable to generate u4:", err)
		} else {
			this.payload = u4.String()
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
		}
	} else {
		if err := sc.Decode("df-token", cookie.Value, &this.payload); err == nil {
			log.Println("df-token:", this.payload)
		}
	}

	c.Map(this)
}

type Token struct {
	payload string
}

func (this Token) Get() string {
	return this.payload
}
