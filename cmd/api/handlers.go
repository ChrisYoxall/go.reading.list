package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	_, err := fmt.Fprintf(w, "status: available\nenvironment: %s\nversion: %s\n", app.config.env, version)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error performing health check: ", err)
		return
	}

}

func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Display a list of books on the reading list.")
	}

	if r.Method == http.MethodPost {
		fmt.Fprintln(w, "Added a new book to the reading list.")
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
		return
	}
	fmt.Fprintf(w, "Display the details of book with ID: %d\n", idInt)
}

func (app *application) updateBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Update the details of a specific book with ID: %d\n", idInt)
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
