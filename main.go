package main

import (
	"citycodes/controllers"
	"citycodes/store"
	"citycodes/views"

	"flag"
	"net/http"

	"github.com/go-fuego/fuego"
)

type Ressources struct {
	Views views.Ressource
	Api   controllers.Ressource
}

func main() {
	serverOptions := []func(*fuego.Server){
		fuego.WithPort(":8083"),
	}

	dbPath := flag.String("db", "./database", "path to database file")
	flag.Parse()

	db := store.InitDB(*dbPath)

	store := store.New(db)

	// Create ressources that will be available in API controllers
	apiRessources := controllers.Ressource{
		SecretCacheRepository: store,
	}

	// Create ressources that will be available in HTML controllers
	viewsRessources := views.Ressource{
		SecretCacheRepository: store,
	}

	rs := Ressources{
		Api:   apiRessources,
		Views: viewsRessources,
	}

	app := fuego.NewServer(
		serverOptions...,
	)

	// fuego.Handle(app, "/", http.FileServer(http.Dir("images/")))
	fuego.Handle(app, "/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))

	rs.Views.Routes(fuego.Group(app, "/"))
	rs.Api.Routes(fuego.Group(app, "/api"))

	app.Run()
}
