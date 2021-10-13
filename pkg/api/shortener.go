package api

import (
	"log"
	"os"
	"urlShortener/pkg/data"
)

type Application struct {
	Models   data.Models
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func NewApplication() *Application {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := data.OpenDB(infoLog)
	if err != nil {
		log.Fatal("encountered an error when attempting to open the database: ", err)
	}

	// Enclosed our Link Model within a Separate struct for future expandability (adding Auth token etc).
	models := data.NewModels(db)

	app := &Application{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
		Models:   models,
	}
	return app

}
