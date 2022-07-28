package main

import (
	"go_ticket_microservice/router"
	"go_ticket_microservice/utils/settings"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	settings.Setup()
}

func main() {
	gin.SetMode(settings.ServerSettings.RunMode)
	routersInit := router.Init()

	port := settings.ServerSettings.Port
	appName := settings.ServerSettings.AppName
	readTimeout := settings.ServerSettings.ReadTimeout
	writeTimeout := settings.ServerSettings.WriteTimeout

	server := &http.Server{
		Addr:         port,
		Handler:      routersInit,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	log.Printf("[info] starting HTTP server for %s, listening on port %s", appName, port)
	server.ListenAndServe()
}
