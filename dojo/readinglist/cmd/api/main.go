// this code is littered with comments to give an understanding of how this
// works. This is not a best practice to allow this in professional app.
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

// application struct will have member methods that are being declared and attached (this is go's way of oop instead of having dedicated classes file and throwing it all there)
// in routes and handlers go files
type application struct {
	config config
	logger *log.Logger
}

/*
Basic Server

- APP (created as a pointer and passed around for usage)
  - Configuration (PORT, ENV)
  - Logger: to print status to standard output cli for printing log statements of server

- Address: need for port info
- Server
  - Address that has port
  - Handler: need multiplexer that can match routes that can be served and operated on
  - idle timeouts unit
  - read timeouts
  - write timeouts

- START SERVER
*/
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
	//	app is a pointer to an application object in mem and it can be passed around, so others can use the pointer and see the object stored at that mem address and see the actual object by using * operator
	//	fmt.Println(app)

	//	app is an object dereference by * with the mem address provided to it via &app
	//	fmt.Println(*app)

	// Assigning designated port from config object :PORT_NUM format
	address := fmt.Sprintf(":%d", config.port)

	// Creating a pointer to http server that is being created
	server := &http.Server{
		Addr:         address,
		Handler:      app.routeHandler(), // routeHandler() is a member function(method) of "application" struct above, that's why . notation can be used here
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// just Used for logging
	logger.Printf("starting %s server on %s", config.env, address)

	// Start server (serving) by listening on port provided by server object's address
	err := server.ListenAndServe()

	// output any error while serving from server
	logger.Fatal(err)

}
