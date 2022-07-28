package settings

import (
	"fmt"
	"os"
	"time"
)

type AppSettings struct {
	Port         string
	AppName      string
	RunMode      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSettings = &AppSettings{}

func Setup() {
	ServerSettings.Port = fmt.Sprintf(":%s", getEnv("PORT", "8081"))
	ServerSettings.AppName = getEnv("APP_NAME", "go-ticket")
	ServerSettings.RunMode = getEnv("RUN_MODE", "debug") // or release
	ServerSettings.ReadTimeout = (60 * time.Second)
	ServerSettings.WriteTimeout = (60 * time.Second)
}

func getEnv(key, fallback string) string {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return fallback
}
