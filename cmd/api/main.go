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

	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := application{
		config: cfg,
		logger: logger,
	}

	addr := fmt.Sprintf(":%d", cfg.port)

	srv := &http.Server{
		Addr:         addr,
		Handler:      app.route(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Using the default mux (i.e. specify nil as the second argument to ListenAndServe) is a security risk as
	// it's stored in a global variable where any third party package can access it and register new handlers.
	// err := http.ListenAndServe(":4000", nil)

	// Instead of this next approach, will define a server struct and use that to configure the server.
	// err := http.ListenAndServe(addr, app.route())

	logger.Printf("starting %s server on %s", cfg.env, addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
