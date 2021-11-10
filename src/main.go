package main

import (
	"log"
	"net/http"
	"flag"
	"os"
	"fmt"
	"time"
)

const apiVersion = "1.0.0"

type config struct {
	port int
	env string
	db struct {
		uri string
	}
	stripe struct {
		secret string
		key string
	}
}

type application struct {
	config config
	infoLogger *log.Logger
	errorLogger *log.Logger
	version string
}

func (app *application) serve() error {
	srv := &http.Server {
		Addr: fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
		IdleTimeout: 30*time.Second,
		ReadTimeout: 10*time.Second,
		ReadHeaderTimeout: 5*time.Second,
		WriteTimeout: 5*time.Second,
	}

	app.infoLogger.Println(fmt.Sprintf("Starting server on port %d in %s mode...", app.config.port, app.config.env))

	return srv.ListenAndServe()
}

func main() {
	var conf config
	
	flag.IntVar(&conf.port, "port", 3000, "Server Port!")
	flag.StringVar(&conf.env, "env", "development", "Server Environment - {development|production}")

	flag.Parse()

	conf.stripe.secret = os.Getenv("STRIPE_SECRET")
	conf.stripe.key = os.Getenv("STRIPE_KEY")

	infoLogger := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application {
		config: conf,
		infoLogger: infoLogger,
		errorLogger: errorLogger,
		version: apiVersion,
	}

	err := app.serve()
	if err != nil {
		log.Fatal(err)
	}

}