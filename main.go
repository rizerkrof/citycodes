package main

import (
	"github.com/go-op/op"
)

func main() {
	app := op.NewServer(
		op.WithPort(":8083"),
	)

	app.Run()
}
