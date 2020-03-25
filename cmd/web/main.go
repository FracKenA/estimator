package main

import (
	"flag"
	"html/template"
	"net/http"
	"os"
	"time"

	// TODO replace this with gorilla?
	// "github.com/golangcollege/sessions"
	"github.com/sirupsen/logrus"
	"log"
)

type application struct {
	// session       *sessions.Session
	log           *logrus.Logger
	templateCache map[string]*template.Template
}

func main() {

	// TODO replace with viper
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// NOTE Needed for http.Serve since it wont accep logrus yet.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Logging
	logger := logrus.New()
	logger.Formatter = new(logrus.JSONFormatter)
	logger.Formatter = new(logrus.TextFormatter)
	logger.Level = logrus.InfoLevel
	logger.Out = os.Stdout

	// HTML cache
	templateCache, err := newTemplateCache("./ui/html")
	if err != nil {
		logger.WithFields(logrus.Fields{
			"event": "templateCache",
		}).Fatal(err)
	}

	// Create application instance
	app := &application{
		log:           logger,
		templateCache: templateCache,
	}

	// Define the server
	srv := &http.Server{
		Addr: *addr,
		// NOTE log needed bc http.Serve
		// will not accept logrus interface
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.WithFields(logrus.Fields{
		"port": *addr,
	}).Info("Starting server")

	err = srv.ListenAndServe()
	logger.WithFields(logrus.Fields{
		"event": "ListenAndServe()",
	}).Fatal(err)
}
