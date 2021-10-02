package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/marstan/bookings/pkg/config"
	"github.com/marstan/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"
var sessionManager *scs.SessionManager

func main() {
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
