package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitConfig() (string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Panicf("Init config error: %v", err)
	}

	TOKEN := os.Getenv("TOKEN")
	WEATHER_KEY := os.Getenv("WEATHER_KEY")
	if TOKEN == "" || WEATHER_KEY == "" {
		log.Panic(".env variables are empty")
	}

	return TOKEN, WEATHER_KEY
}
