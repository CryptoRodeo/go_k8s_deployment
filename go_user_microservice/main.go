package main

import (
	"log"
	"net/http"

	"go_user_microservice/router"
	"go_user_microservice/utils/settings"

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
