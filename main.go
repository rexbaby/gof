package main

import (
	"gof/api"
	"gof/migrations"

	"gofr.dev/pkg/gofr"
)

func main() {
	// initialise gofr object
	app := gofr.New()

	app.Migrate(migrations.All())

	// register route greet
	app.GET("/", func(ctx *gofr.Context) (interface{}, error) {

		return "Hello World!", nil
	})

	api.RegisterCRUD(app)

	// Runs the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Run()
}
