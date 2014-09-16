package main

import (
	ctrl "df/api/controllers"
	"df/api/models"
	. "df/api/utils"
	// "encoding/json"
	"fmt"
	"github.com/3d0c/binding"
	"github.com/go-martini/martini"
	"github.com/gorilla/securecookie"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
	// "os"
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
		Directory:  "templates",                // Specify what path to load the templates from.
		Layout:     "layout",                   // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		Charset:    "UTF-8",                    // Sets encoding for json and html content-types. Default is "UTF-8".
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

	// Boot

	// route.Group("/boot", func(router martini.Router) {
	// 	router.Get("/:id?", func(render render.Render, params martini.Params) {
	// 		render.HTML(200, "index", nil)
	// 	})
	// })

	route.Get("/boot/ext", func(user *models.User, r render.Render) {
		// userJson, _ := json.Marshal(user.Object)
		// log.Println("user:", userJson)
		// os.Stdout.Write(userJson)
		r.HTML(200, "ext", user.Object)
	})

	route.Get("/boot", func(render render.Render) {
		render.HTML(200, "index", nil)
	})

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
