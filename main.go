package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/mtproject-ru/server/config"
	"github.com/mtproject-ru/server/database"
	"github.com/mtproject-ru/server/services"
	"github.com/mtproject-ru/server/web"
)

func main() {
	config.SetupEnv()
	config.SetupLogging()

	database.SetupDatabase()
	SetupMigrations()

	app := config.SetupFiber()

	services.SetupServices(app)
	web.SetupWeb(app)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		slog.Info("SERVER_PORT не найдена в переменных окружения. Используем 3000 порт")
		port = ":3000"
	}

	log.Fatal(app.Listen(port))
}
