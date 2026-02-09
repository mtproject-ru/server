package config

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"
)

func SetupLogging() {
	year, month, day := time.Now().Date()

	logFile, err := os.OpenFile(
		fmt.Sprintf("./logs/%d-%d-%d.log", year, month, day),
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0666,
	)
	if err != nil {
		fmt.Printf("Не удалось создать файл логов: %v\n", err)
	}

	logger := slog.New(
		slog.NewTextHandler(
			io.MultiWriter(logFile, os.Stdout),
			nil,
		),
	)
	slog.SetDefault(logger)
}
