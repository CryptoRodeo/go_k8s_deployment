package settings

import (
	"fmt"
	"os"
	"time"
)

var ServicePorts = map[string]string{
	"tickets": getEnv("TICKET_SERVICE", "localhost:8081"),
}

var ServiceEndpoints = map[string]string{
	"ticket-service": fmt.Sprintf("http://%s/api/v1/tickets", ServicePorts["tickets"]),
}

type AppSettings struct {
	Port         string
	AppName      string
	RunMode      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	// CORS
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
}

type CorsConfig struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
}

var ServerSettings = &AppSettings{}

var Cors = &CorsConfig{}

func Setup() {
	// SERVER
	ServerSettings.Port = fmt.Sprintf(":%s", getEnv("PORT", "8080"))
	ServerSettings.AppName = getEnv("APP_NAME", "go-users")
	ServerSettings.RunMode = getEnv("RUN_MODE", "debug") // or release
	ServerSettings.ReadTimeout = (60 * time.Second)
	ServerSettings.WriteTimeout = (60 * time.Second)
	// CORS
	frontend := getEnv("FRONT_END_ENDPOINT", "http://localhost:3000")
	Cors.AllowOrigins = []string{frontend}
	Cors.AllowMethods = []string{"GET"}
	Cors.AllowHeaders = []string{"Origin", "Content-Type"}
	Cors.AllowCredentials = true
}

func getEnv(key, fallback string) string {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return fallback
}
