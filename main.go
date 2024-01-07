package main

import (
	"github.com/go-fuego/fuego"
	"citycodes/store"
	"citycodes/controller"
)

func main() {
	db := store.InitDB("./database")

	queries := store.New(db)

	api := controller.NewRessource(*queries)

	app := fuego.NewServer(
		fuego.WithPort(":8083"),
	)

	api.Routes(app)

	app.Run()
}
