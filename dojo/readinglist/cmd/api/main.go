package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	// config, get flags from Environment or use assigned defaults
	var config config
	flag.IntVar(&config.port, "port", 4000, "API Server Port")
	flag.StringVar(&config.env, "env", "dev", "Environment (dev|stage|prod)")
	flag.Parse()

	// logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// define app
	app := &application{
		config: config,
		logger: logger,
	}

	// start listening on the port
	address := fmt.Sprintf(":%d", config.port)

	server := &http.Server{
		Addr:         address,
		Handler:      app.route(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", config.env, address)
	err := server.ListenAndServe()
	logger.Fatal(err)
}
