package api

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/teris-io/shortid"
	"net/http"
	"strings"
	"urlShortener/pkg/data"
)

func (app *Application) ShortenHandler(w http.ResponseWriter, r *http.Request) {

	// Input holds the URL for which we generate a short URL.
	type Input struct {
		Url string `json:"url"`
	}

	var input Input
	err := app.ReadJSON(w, r, &input)
	if err != nil {
		app.ErrorLog.Println("Failed to read JSON", err)
		err := app.writeJSON(w, http.StatusBadRequest, "Bad Request, please check the JSON body (include a 'url' field with a string argument", nil)
		if err != nil {
			app.ErrorLog.Println("encountered an error when attempting to write json to io.writer", err)
		}
		return
	}

	// Generate a short URL for the given Input
	shortURL := shortid.MustGenerate()
	for shortURL == "shorten" {
		shortURL = shortid.MustGenerate()
	}

	// Check if URL formatted correctly, without this the redirect function does not work properly as it thinks we have
	//a relative path rather than an absolute path.
	if strings.HasPrefix(input.Url, "www.") {
		input.Url = fmt.Sprintf("http://%s", input.Url)
	}

	app.InfoLog.Printf("generated short url %q for %s ", shortURL, input.Url)

	// Populate link data to pass to DB
	link := data.Link{
		ShortURL: shortURL,
		LongURL:  input.Url,
	}

	l, err := app.Models.Links.CreateLink(link)
	link.Id = l.Id

	if err != nil {
		app.ErrorLog.Println("failed to create a short url: ", err)
		err := app.writeJSON(w, http.StatusInternalServerError, "Sorry We Encountered An Internal Server Error", nil)
		if err != nil {
			app.ErrorLog.Println("encountered an error when attempting to write json to io.writer", err)
		}
		return
	}

	// Wrap the Link to provide more explanation of what's returned
	wrapped := wrapper{
		"Short URL": link,
	}

	err = app.writeJSON(w, http.StatusOK, wrapped, nil)
	if err != nil {
		app.ErrorLog.Println("encountered an error when attempting to write json to io.writer", err)
	}

}

func (app *Application) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the desired URL from the request Path.
	v := mux.Vars(r)
	urlVar := v["url"]

	// Create a Link within the database.
	l, err := app.Models.Links.GetLink(urlVar)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			fmt.Fprintf(w, "the requested URL was not found, please double check the short code")
			return
		default:
			app.ErrorLog.Println("encountered an error when querying db to lookup short code", err)
			fmt.Fprint(w, "sorry, we encountered an unexpected internal server error")
			return
		}
	}

	// check no urls slipped through the new where we had a relative path instead of an absolute path.
	url := l.LongURL

	if strings.HasPrefix(url, "www.") {
		url = "http://" + url
	}

	// Set the location header to the redirect URL
	w.Header().Set("Location", url)

	// Redirect to an absolute path
	http.Redirect(w, r, url, http.StatusSeeOther)

}
