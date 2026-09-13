package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitConfig() (string, string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Panicf("Init config error: %v", err)
	}

	TOKEN := os.Getenv("TOKEN")
	WEATHER_KEY := os.Getenv("WEATHER_KEY")
	GORM_KEY := os.Getenv("GORM_KEY")
	if TOKEN == "" || WEATHER_KEY == "" || GORM_KEY == "" {
		log.Panic(".env variables are empty")
	}

	return TOKEN, WEATHER_KEY, GORM_KEY
}

func InitDB(dsn string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("DataBase init")
	return db
}
