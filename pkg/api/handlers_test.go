package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Create a separate structure to parse the JSON Response Body. This is necessary because we wrapped the original JSON
// within a map to increase the readability of the structure however now we cannot re-use that structure to decode the response.
type T struct {
	ShortURL struct {
		Id       int    `json:"id"`
		LongUrl  string `json:"long_url"`
		ShortUrl string `json:"short_url"`
	} `json:"Short URL"`
}

func TestApplication_ShortenHandler(t *testing.T) {

	app := NewApplication()

	type in struct {
		Url string `json:"url"`
	}

	input := in{
		Url: "https://www.amazon.co.uk",
	}

	js, err := json.Marshal(input)
	i := bytes.NewReader(js)
	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8000/shorten", i)
	if err != nil {
		log.Println(err)
	}
	rr := httptest.NewRecorder()
	// No need to use Gorilla Mux, std http handler is fine
	handler := http.HandlerFunc(app.ShortenHandler)
	handler.ServeHTTP(rr, req)

	var link T
	err = json.Unmarshal(rr.Body.Bytes(), &link)
	if err != nil {
		log.Println("err: ", err)
	}

	t.Run("Long URL is correct", func(t *testing.T) {
		if input.Url != link.ShortURL.LongUrl {
			t.Errorf("got %s want %s", link.ShortURL.LongUrl, input.Url)
		}
	})

	t.Run("ID is non-negative int", func(t *testing.T) {
		if link.ShortURL.Id <= 0 {
			t.Errorf("expected a non-negative integer for ID, got %d", link.ShortURL.Id)
		}
	})

	t.Run("The ShortURL has a length of 9", func(t *testing.T) {
		if len(link.ShortURL.ShortUrl) != 9 {
			t.Errorf("got %s want %d", link.ShortURL.ShortUrl, 9)
		}
	})

	t.Run("Correct status code", func(t *testing.T) {
		if rr.Code != http.StatusOK {
			t.Errorf("want %v got %v", http.StatusOK, rr.Code)
		}
	})

}

func TestApplication_RedirectHandler(t *testing.T) {
	// Short:"YH-nYjDnR"
	// Long:"http://www.amazon.co.uk"

	want := "http://www.amazon.co.uk"
	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8000/YH-nYjDnR", nil)
	if err != nil {
		log.Println("Failed to generate a request for the google test handler")
	}

	app := NewApplication()

	rr := httptest.NewRecorder()

	// Use Gorilla mux to get url as route variable.
	router := mux.NewRouter()
	router.HandleFunc("/{url}", app.RedirectHandler)
	router.ServeHTTP(rr, req)
	router.ServeHTTP(rr, req)
	fmt.Println(rr.Body.String())
	loc := rr.Header().Get("Location")

	t.Run("redirect function produces correct Location in header", func(t *testing.T) {
		if loc != want {
			t.Errorf("got %s want %s", loc, want)
		}

	})

	t.Run("redirect function produces correct status code", func(t *testing.T) {
		if rr.Code != http.StatusSeeOther {
			t.Errorf("got %v want %v", rr.Code, http.StatusSeeOther)
		}
	})

}
