package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type envelope map[string]any

// Credit: Alex Edwards, Let's Go Further
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {

	// Limit the size of the request body to 1MB.
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	// Note use of decoding rather than unmarshalling here. Both can usually be used interchangeably, but
	// decoding is a more general approach and uses streams rather reading the whole document in one go.
	// Initialize the json.Decoder, and call the DisallowUnknownFields() method on it before decoding.
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		// Good tips on custom error handling in Alex Edwards, Let's Go Further Chapter 4
		return err
	}

	err := dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON object")
	}

	return nil
}
