package main

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/marstan/bookings/internals/config"
	"github.com/marstan/bookings/internals/models"
	"github.com/marstan/bookings/internals/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"
var sessionManager *scs.SessionManager

func main() {
	gob.Register(models.Reservation{})

	var app config.AppConfig

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	app.Session = sessionManager

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc

	render.SetConfig(&app)

	log.Println("Starting server at port: ", portNumber)

	serv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = serv.ListenAndServe()
	if err != nil {
		log.Fatal("Error while running server", err)
	}
}
