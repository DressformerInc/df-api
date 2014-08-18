package main

import (
	ctrl "df/api/controllers"
	"df/api/models"
	. "df/api/utils"
	"fmt"
	"github.com/3d0c/martini-contrib/config"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/encoder"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	config.Init("./config.json")
	config.LoadInto(AppConfig)
}

func some(a int) martini.Handler {
	return func(ctx martini.Context, req *http.Request) {
		ctx.Map(a)
	}
}

func foo(i int) string {
	return fmt.Sprintln(i)
}

func main() {
	m := martini.New()
	route := martini.NewRouter()

	m.Use(func(c martini.Context, w http.ResponseWriter) {
		c.MapTo(encoder.JsonEncoder{}, (*encoder.Encoder)(nil))
		w.Header().Set("Content-Type", "application/json")
	})

	m.Use(func(w http.ResponseWriter, req *http.Request) {
		if origin := req.Header.Get("Origin"); origin != "" {
			w.Header().Add("Access-Control-Allow-Origin", origin)
		} else {
			w.Header().Add("Access-Control-Allow-Origin", "*")
		}

		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		w.Header().Add("Cache-Control", "max-age=2592000")
		w.Header().Add("Pragma", "public")
		w.Header().Add("Cache-Control", "public")
	})

	route.Options("/**")

	route.Get("/user",
		construct(&models.User{}),
		construct(&ctrl.User{}),
		(*ctrl.User).Find,
	)

	route.Get("/post",
		construct(&models.User{}),
		construct(&ctrl.Post{}),
		(*ctrl.Post).Find,
	)

	m.Action(route.Handle)

	log.Printf("Waiting for connections on %v...\n", AppConfig.ListenOn())

	go func() {
		if err := http.ListenAndServe(AppConfig.ListenOn(), m); err != nil {
			log.Fatal(err)
		}
	}()

	if err := http.ListenAndServeTLS(AppConfig.HttpsOn(), AppConfig.SSLCert(), AppConfig.SSLKey(), m); err != nil {
		log.Fatal(err)
	}
}

func construct(obj interface{}, args ...interface{}) martini.Handler {
	return func(ctx martini.Context, r *http.Request) {
		switch t := obj.(type) {
		case models.Model:
			ctx.Map(obj.(models.Model).Construct(args))

		case ctrl.Controller:
			ctx.Map(obj.(ctrl.Controller).Construct(args))

		default:
			panic(fmt.Sprintln("Unexpected type:", t))
		}
	}
}
