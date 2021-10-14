package api

import (
	"log"
	"os"
	"urlShortener/pkg/data"
)

// Application holds the main App Logic with reference to the data store and logging.
type Application struct {
	Models   data.Models
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

// NewApplication is a singleton which returns a single instance of the main Application struct. It is responsible for
// opening any required db connection.
func NewApplication() *Application {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := data.OpenDB(infoLog)
	if err != nil {
		log.Fatal("encountered an error when attempting to open the database: ", err)
	}

	models := data.NewModels(db)

	app := &Application{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
		Models:   models,
	}
	return app

}
