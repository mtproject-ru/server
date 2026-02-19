package database

import (
	"log"
	"log/slog"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func SetupDatabase() {
	connectionString := os.Getenv("DATABASE_CONNECTION_STRING")
	if connectionString == "" {
		log.Fatal("DATABASE_CONNECTION_STRING не найдено в переменных окружения")
	}

	var err error
	DB, err = sqlx.Connect("pgx", connectionString)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	slog.Info("Успешно подключились к базе данных")
}
