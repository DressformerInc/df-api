package main

import (
	ctrl "df/api/controllers"
	"df/api/models"
	. "df/api/utils"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/gorilla/securecookie"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	InitConfigFrom("./config.json")
}

func main() {
	m := martini.New()
	route := martini.NewRouter()

	m.Map(securecookie.New(AppConfig.HashKey(), AppConfig.BlockKey()))

	m.Use(render.Renderer(render.Options{
		Directory: "templates",
		// Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
		IndentJSON: true,
	}))

	m.Use(martini.Static("public"))

	m.Use(LogHandler)
	m.Use(CorsHandler)
	m.Use(TokenHandler)

	// Creates user model for each request
	m.Use(func(c martini.Context, token Token) {
		c.Map((&models.User{}).Construct(token))
	})

	route.Options("/**")

	// Index

	route.Get("/",
		construct(&ctrl.Widget{}),
		construct(&models.Garment{}),
		(*ctrl.Widget).Index,
	)

	route.Get("/tryon", func(r render.Render) {
		r.HTML(200, "tryon", nil)
	})

	// Widget

	route.Get("/widget/ext",
		binding.Form(models.URLOptionsScheme{}),
		construct(&ctrl.Widget{}, "widget-ext"),
		construct(&models.Garment{}),
		(*ctrl.Widget).FindAll,
	)

	route.Get("/widget/ext/:id",
		construct(&ctrl.Widget{}, "widget-ext"),
		construct(&models.Garment{}),
		(*ctrl.Widget).Get,
	)

	route.Get("/widget",
		binding.Form(models.URLOptionsScheme{}),
		construct(&ctrl.Widget{}, "widget"),
		construct(&models.Garment{}),
		(*ctrl.Widget).FindAll,
	)

	route.Get("/widget/:id",
		construct(&ctrl.Widget{}, "widget"),
		construct(&models.Garment{}),
		(*ctrl.Widget).Get,
	)

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

	// Materials

	route.Post("/materials",
		binding.Json([]models.MaterialScheme{}),
		ErrorHandler,
		construct(&ctrl.Material{}),
		(*ctrl.Material).Create,
	)

	route.Get("/materials",
		binding.Form(models.URLOptionsScheme{}),
		ErrorHandler,
		construct(&ctrl.Material{}),
		(*ctrl.Material).FindAll,
	)

	route.Get("/materials/:id",
		construct(&ctrl.Material{}),
		(*ctrl.Material).Find,
	)

	route.Put("/materials/:id",
		binding.Json(models.MaterialScheme{}),
		ErrorHandler,
		construct(&ctrl.Material{}),
		(*ctrl.Material).Put,
	)

	//

	m.Action(route.Handle)

	log.Printf("Waiting for connections on %v...\n", AppConfig.ListenOn())

	go func() {
		if err := http.ListenAndServeTLS(AppConfig.HttpsOn(), AppConfig.SSLCert(), AppConfig.SSLKey(), m); err != nil {
			log.Println(err)
		}

	}()

	if err := http.ListenAndServe(AppConfig.ListenOn(), m); err != nil {
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
