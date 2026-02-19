package main

import (
	"embed"
	"log"
	"log/slog"

	"github.com/mtproject-ru/server/database"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func SetupMigrations() {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Не удалось установить диалект БД: %v", err)
	}

	if err := goose.Up(database.DB.DB, "migrations"); err != nil {
		log.Fatalf("Не удалось запустить миграцию БД: %v", err)
	}

	slog.Info("Успешно создали все необходимое в базе данных")
}
