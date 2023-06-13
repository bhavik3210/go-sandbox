package main

import (
	"fmt"
	"net/http"
)

func main() {
	// learn what is locally scoped serve mux?
	// able to prevent someone else from modifying global variable
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", healthCheck)
	err := http.ListenAndServe(":4023", mux)
	if err != nil {
		fmt.Println(err)
	}
}

func healthCheck(writer http.ResponseWriter, requestPointer *http.Request) {
	fmt.Fprintln(writer, "status: avaialable")
	fmt.Fprintf(writer, "environment: %s \n", "dev")
	fmt.Fprintf(writer, "version: %s\n", "1.0.0")
}