package settings

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
)

type AppSettings struct {
	Port         string
	AppName      string
	RunMode      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSettings = &AppSettings{}

var Cors = cors.Config{}

func Setup() {
	ServerSettings.Port = fmt.Sprintf(":%s", getEnv("PORT", "8081"))
	ServerSettings.AppName = getEnv("APP_NAME", "go-ticket")
	ServerSettings.RunMode = getEnv("RUN_MODE", "debug") // or release
	ServerSettings.ReadTimeout = (60 * time.Second)
	ServerSettings.WriteTimeout = (60 * time.Second)
	// CORS
	frontend := getEnv("FRONT_END_ENDPINT", "http://localhost:3000")
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
