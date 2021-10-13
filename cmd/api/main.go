package main

import (
	"urlShortener/pkg/api"
)

func main() {
	app := api.NewApplication()
	err := app.Serve()
	if err != nil {
		app.ErrorLog.Fatal("encountered a fatal error ", err)
	}
}
