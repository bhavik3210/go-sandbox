package main

import "net/http"

func (app *application) route() *http.ServeMux {

	// create a new server
	mux := http.NewServeMux()

	// Attach a Handler that can handle incoming routes
	mux.HandleFunc("/v1/healthcheck", app.healthcheck)
	mux.HandleFunc("/v1/books", app.getCreateBooksHandler)
	mux.HandleFunc("/v1/books/", app.getUpdateDeleteBooksHandler)
	return mux
}
