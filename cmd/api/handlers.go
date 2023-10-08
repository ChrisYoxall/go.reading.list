package main

import (
	"fmt"
	"log"
	"net/http"
	"reading.list/internal/data"
	"strconv"
	"time"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	status := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"status": status}); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("Error returning health check: %v\n", err)
		return
	}

}

func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		books := []data.Book{
			{
				ID:        1,
				CreatedAt: time.Now(),
				Title:     "The Darkening of Tristram",
				Published: 1998,
				Pages:     300,
				Genres:    []string{"Fiction", "Thriller"},
				Rating:    4.5,
				Version:   1,
			},
			{
				ID:        2,
				CreatedAt: time.Now(),
				Title:     "The Legacy of Cain",
				Published: 2007,
				Pages:     432,
				Genres:    []string{"Fiction", "Adventure"},
				Rating:    4.9,
				Version:   1,
			},
			{
				ID:        3,          // system generated
				CreatedAt: time.Now(), // system generated
				Title:     "The Black Soulstone",
				Version:   1, // system generated
			},
		}

		if err := app.writeJSON(w, http.StatusOK, envelope{"books": books}); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Printf("Error returning books: %v\n", err)
			return
		}
	}

	if r.Method == http.MethodPost {
		var input struct {
			Title     string   `json:"title"`
			Published int      `json:"published"`
			Pages     int      `json:"pages"`
			Genres    []string `json:"genres"`
			Rating    float32  `json:"rating"`
		}

		err := app.readJSON(w, r, &input)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Printf("Error reading json during create book: %v\n", err)
			return
		}

		fmt.Fprintf(w, "%+v\n", input)
	}

}

func (app *application) getUpdateDeleteBooksHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		app.getBook(w, r)
	case http.MethodPut:
		app.updateBook(w, r)
	case http.MethodDelete:
		app.deleteBook(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

}

func (app *application) getBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Printf("Error parsing ID from path: %s\n", r.URL.Path)
		return
	}

	book := data.Book{
		ID:        idInt,
		CreatedAt: time.Now(),
		Title:     "The Hitchhiker's Guide to the Galaxy",
		Published: 2000,
		Genres:    []string{"science fiction", "comedy"},
		Rating:    4.8,
		Version:   1,
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"book": book}); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("Error returning book ID: %d\n", idInt)
		return
	}
}

func (app *application) updateBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Printf("Error parsing ID from path: %s\n", r.URL.Path)
		return
	}

	var input struct {
		Title     *string  `json:"title"`
		Published *int     `json:"published"`
		Pages     *int     `json:"pages"`
		Genres    []string `json:"genres"`
		Rating    *float32 `json:"rating"`
	}

	// dummy data for existing book to be updated
	book := data.Book{
		ID:        idInt,
		CreatedAt: time.Now(),
		Title:     "Echoes in the Darkness",
		Published: 2019,
		Pages:     300,
		Genres:    []string{"Fiction", "Thriller"},
		Rating:    4.5,
		Version:   1,
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Printf("Error reading json during upate: %s\v", err)
		return
	}

	if input.Title != nil {
		book.Title = *input.Title
	}

	if input.Published != nil {
		book.Published = *input.Published
	}

	if input.Pages != nil {
		book.Pages = *input.Pages
	}

	if len(input.Genres) > 0 {
		book.Genres = input.Genres
	}

	if input.Rating != nil {
		book.Rating = *input.Rating
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"book": book}); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("Error writing json during update: %v\n", err)
		return
	}
}

func (app *application) deleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Delete a specific book with ID: %d\n", idInt)
}
