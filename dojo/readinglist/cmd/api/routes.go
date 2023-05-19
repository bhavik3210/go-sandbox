package main

import "net/http"

// this method is like extension functions in Kotlin, btw there is no special keyword to refrence reciever like app here
func (app *application) routeHandler() *http.ServeMux {

	// create a new server
	mux := http.NewServeMux()

	// Attach a Handler that can handle incoming routes
	mux.HandleFunc("/v1/healthcheck", app.healthcheck)
	mux.HandleFunc("/v1/books", app.getCreateBooksHandler)
	mux.HandleFunc("/v1/books/", app.getUpdateDeleteBooksHandler)
	return mux
}
