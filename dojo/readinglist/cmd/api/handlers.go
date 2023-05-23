package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"readinglist.web.dojo/internal/data"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	if err := app.WriteJSON(w, http.StatusOK, envelope{"app": data}); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusMethodNotAllowed)
		return
	}
}

func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	data := []data.Book{
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
			Title:     "The Legacy of Deckard Cain",
			Published: 2007,
			Pages:     300,
			Genres:    []string{"Fiction", "Adventure"},
			Rating:    4.9,
			Version:   1,
		},
	}

	if err := app.WriteJSON(w, http.StatusOK, envelope{"books": data}); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (app *application) getUpdateDeleteBooksHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		app.getBook(writer, request)
	case http.MethodPut:
		app.updateBook(writer, request)
	case http.MethodDelete:
		app.deleteBook(writer, request)
	default:
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (app *application) getBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	data := data.Book{
		ID:        idInt,
		CreatedAt: time.Now(),
		Title:     "Echoes in the Darkness",
		Published: 2019,
		Pages:     300,
		Genres:    []string{"Fiction", "Thriller"},
		Rating:    4.5,
		Version:   1,
	}

	if err := app.WriteJSON(w, http.StatusOK, envelope{"book": data}); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (app *application) updateBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Update the details of the book with ID: %d", idInt)
}

func (app *application) deleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Deleted the book with ID: %d", idInt)
}
