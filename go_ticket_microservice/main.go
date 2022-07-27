package main

import (
	"go_ticket_microservice/router"
	"go_ticket_microservice/utils/settings"
	"log"
	"net/http"
)

func init() {
	settings.Setup()
}

func main() {
	routersInit := router.InitRouter()

	port := settings.ServerSettings.Port
	appName := settings.ServerSettings.AppName

	server := &http.Server{
		Addr:    port,
		Handler: routersInit,
	}

	log.Printf("[info] starting HTTP server for %s, listening on port %s", appName, port)
	server.ListenAndServe()
}
