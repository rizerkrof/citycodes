package main

import (
	"github.com/go-op/op"
	"citycodes/store"
	"citycodes/controller"
)

func main() {
	db := store.InitDB("./database")

	queries := store.New(db)

	api := controller.NewRessource(*queries)

	app := op.NewServer(
		op.WithPort(":8083"),
	)

	api.Routes(app)

	app.Run()
}
