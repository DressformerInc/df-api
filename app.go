package main

import (
	ctrl "df/api/controllers"
	"df/api/models"
	. "df/api/utils"
	"fmt"
	"github.com/3d0c/binding"
	"github.com/3d0c/martini-contrib/config"
	"github.com/go-martini/martini"
	"github.com/gorilla/securecookie"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	config.Init("./config.json")
	config.LoadInto(AppConfig)
}

func main() {
	m := martini.New()
	route := martini.NewRouter()

	m.Map(securecookie.New(AppConfig.HashKey(), AppConfig.BlockKey()))

	m.Use(render.Renderer(render.Options{
		IndentJSON: true,
	}))

	m.Use(LogHandler)
	m.Use(CorsHandler)
	m.Use(TokenHandler)

	// Creates user model for each request
	m.Use(func(c martini.Context, token Token) {
		c.Map((&models.User{}).Construct(token))
	})

	route.Options("/**")

	// User

	route.Get("/user",
		construct(&ctrl.User{}),
		(*ctrl.User).Find,
	)

	// Garments

	route.Get("/garments",
		binding.Form(models.URLOptionsScheme{}),
		ErrorHandler,
		construct(&ctrl.Garment{}),
		(*ctrl.Garment).FindAll,
	)

	route.Get("/garments/:id",
		construct(&ctrl.Garment{}),
		(*ctrl.Garment).Find,
	)

	route.Post("/garments",
		binding.Json(models.GarmentScheme{}),
		ErrorHandler,
		construct(&ctrl.Garment{}),
		(*ctrl.Garment).Create,
	)

	route.Put("/garments/:id",
		binding.Json(models.GarmentScheme{}),
		ErrorHandler,
		construct(&ctrl.Garment{}),
		(*ctrl.Garment).Put,
	)

	route.Delete(
		"/garments/:id",
		construct(&ctrl.Garment{}),
		(*ctrl.Garment).Remove,
	)

	// Dummies

	route.Get("/dummies",
		binding.Form(models.URLOptionsScheme{}),
		ErrorHandler,
		construct(&ctrl.Dummy{}),
		(*ctrl.Dummy).FindAll,
	)

	route.Get("/dummies/:id",
		construct(&ctrl.Dummy{}),
		(*ctrl.Dummy).Find,
	)

	route.Post("/dummies",
		binding.Json(models.DummyScheme{}),
		ErrorHandler,
		construct(&ctrl.Dummy{}),
		(*ctrl.Dummy).Create,
	)

	route.Put("/dummies/:id",
		binding.Json(models.DummyScheme{}),
		ErrorHandler,
		construct(&ctrl.Dummy{}),
		(*ctrl.Dummy).Put,
	)

	route.Delete("/dummies/:id",
		construct(&ctrl.Dummy{}),
		(*ctrl.Dummy).Remove,
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

// @weird args... accumulates values, on append
func construct(obj interface{}, args ...interface{}) martini.Handler {
	return func(ctx martini.Context, r *http.Request) {
		switch t := obj.(type) {
		case models.Model:
			ctx.Map(obj.(models.Model).Construct(args...))

		case ctrl.Controller:
			ctx.Map(obj.(ctrl.Controller).Construct(args...))

		default:
			panic(fmt.Sprintln("Unexpected type:", t))
		}
	}
}
