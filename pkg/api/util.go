package api

import (
	"encoding/json"
	"net/http"
)

type wrapper map[string]interface{}

func (app *Application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {

	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// pretty terminal output
	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	// We cannot set the header after called to WriteHeader or Write. We cannot Write the Header after Writing to
	// ResponseWriter. Calls to w.Write without w.WriteHeader will return http.StatusOK.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func (app *Application) ReadJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	return dec.Decode(dst)

}
