package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

//GetRoutes registers routes which are handled by Mux.
func (app *Application) GetRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/shorten", app.ShortenHandler).Methods(http.MethodPost)
	router.HandleFunc("/{url}", app.RedirectHandler).Methods(http.MethodGet)
	return router
}
