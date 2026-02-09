package config

import (
	"log"

	"github.com/joho/godotenv"
)

func SetupEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Не удалось загрузить .env файл: %v", err)
	}
}
