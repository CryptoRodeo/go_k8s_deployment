package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go_ticket_microservice/router"
)

var PORT string = fmt.Sprintf(":%s", getEnv("PORT", "8081"))
var APP_NAME string = getEnv("APP_NAME", "go-ticket")

func main() {

	routersInit := router.InitRouter()

	server := &http.Server{
		Addr:    PORT,
		Handler: routersInit,
	}

	log.Printf("[info] starting HTTP server, listening on port %s", PORT)
	server.ListenAndServe()
}

func getEnv(key, fallback string) string {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return fallback
}
