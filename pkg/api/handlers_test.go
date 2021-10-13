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
	handler := http.HandlerFunc(app.ShortenHandler)
	handler.ServeHTTP(rr, req)

	var link T
	err = json.Unmarshal(rr.Body.Bytes(), &link)
	if err != nil {
		log.Println("err: ", err)
	}

	// Assert that the LongURL is stored correctly, the ID > 0, and the ShortURL has a len of 9.

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

}

func TestApplication_RedirectHandler(t *testing.T) {
	// "YH-nYjDnR",

	want := "http://www.amazon.co.uk"
	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8000/YH-nYjDnR", nil)
	if err != nil {
		log.Println("Failed to generate a request for the google test handler")
	}

	app := NewApplication()

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/{url}", app.RedirectHandler)
	router.ServeHTTP(rr, req)
	router.ServeHTTP(rr, req)
	fmt.Println(rr.Body.String())
	loc := rr.Header().Get("Location")

	t.Run("check location correct", func(t *testing.T) {
		if loc != want {
			t.Errorf("got %s want %s", loc, want)
		}

	})
}
