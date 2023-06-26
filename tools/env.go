package tools

import (
	"log"
	"github.com/joho/godotenv"
	"os"
)

func Init_Env() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Get_env(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}