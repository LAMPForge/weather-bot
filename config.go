package main

import (
	"os"
)

type Config struct {
	API_KEY            string
	LAT                string
	LONG               string
	TELEGRAM_BOT_TOKEN string
	CHAT_ID            string
}

func getEnv() *Config {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }
	conf := Config{
		API_KEY:            os.Getenv("API_KEY"),
		LAT:                os.Getenv("LAT"),
		LONG:               os.Getenv("LONG"),
		TELEGRAM_BOT_TOKEN: os.Getenv("TELEGRAM_BOT_TOKEN"),
		CHAT_ID:            os.Getenv("CHAT_ID"),
	}
	return &conf
}
