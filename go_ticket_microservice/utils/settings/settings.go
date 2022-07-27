package settings

import (
	"fmt"
	"os"
)

type AppSettings struct {
	Port    string
	AppName string
}

var ServerSettings = &AppSettings{}

func Setup() {
	ServerSettings.Port = fmt.Sprintf(":%s", getEnv("PORT", "8081"))
	ServerSettings.AppName = getEnv("APP_NAME", "go-ticket")
}

func getEnv(key, fallback string) string {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return fallback
}
