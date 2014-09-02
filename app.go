package main

import (
	ctrl "df/api/controllers"
	"df/api/models"
	. "df/api/utils"
	"fmt"
	"github.com/3d0c/binding"
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

func main() {
	m := martini.New()
	route := martini.NewRouter()

	m.Use(func(c martini.Context, w http.ResponseWriter) {
		c.MapTo(encoder.JsonEncoder{PrettyPrint: true}, (*encoder.Encoder)(nil))
		w.Header().Set("Content-Type", "application/json")
	})

	m.Use(func(w http.ResponseWriter, req *http.Request) {
		log.Println(req.RequestURI)

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
	})

	route.Options("/**")

	// User

	route.Get("/user",
		construct(&models.User{}),
		construct(&ctrl.User{}),
		(*ctrl.User).Find,
	)

	// Garments

	route.Get("/garments",
		binding.Bind(models.URLOptionsScheme{}),
		construct(&models.User{}),
		construct(&ctrl.Garment{}),
		(*ctrl.Garment).FindAll,
	)

	route.Get("/garments/:id",
		construct(&models.User{}),
		construct(&ctrl.Garment{}),
		(*ctrl.Garment).Find,
	)

	route.Post("/garments",
		binding.Bind(models.GarmentScheme{}),
		construct(&models.User{}),
		construct(&ctrl.Garment{}),
		(*ctrl.Garment).Create,
	)

	route.Put("/garments/:id",
		binding.Bind(models.GarmentScheme{}),
		construct(&models.User{}),
		construct(&ctrl.Garment{}),
		(*ctrl.Garment).Put,
	)

	route.Delete(
		"/garments/:id",
		construct(&models.User{}),
		construct(&ctrl.Garment{}),
		(*ctrl.Garment).Remove,
	)

	// Dummies

	route.Get("/dummies",
		binding.Bind(models.URLOptionsScheme{}),
		construct(&models.User{}),
		construct(&ctrl.Dummy{}),
		(*ctrl.Dummy).FindAll,
	)

	route.Get("/dummies/:id",
		construct(&models.User{}),
		construct(&ctrl.Dummy{}),
		(*ctrl.Dummy).Find,
	)

	route.Post("/dummies",
		binding.Bind(models.DummyScheme{}),
		construct(&models.User{}),
		construct(&ctrl.Dummy{}),
		(*ctrl.Dummy).Create,
	)

	route.Put("/dummies/:id",
		construct(&models.User{}),
		construct(&ctrl.Dummy{}),
		(*ctrl.Dummy).Put,
	)

	route.Delete("/dummies/:id",
		construct(&models.User{}),
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
