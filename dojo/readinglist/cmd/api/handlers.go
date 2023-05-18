package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) healthcheck(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(writer, "status: available")
	fmt.Fprintf(writer, "environment: %s\n", app.config.env)
	fmt.Fprintf(writer, "version: %s\n", version)
}

func (app *application) getCreateBooksHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		fmt.Fprintln(writer, "Dispaly a list of the book on the reading list")
	}

	if request.Method == http.MethodPost {
		fmt.Fprintln(writer, "Added a new book to the reading list")
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
	fmt.Fprintf(w, "Display the details of the book with ID: %d", idInt)
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
