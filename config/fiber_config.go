package config

import (
	"log/slog"

	"github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
)

func SetupFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		StructValidator: NewStructValidator(),
		ErrorHandler: func(ctx fiber.Ctx, err error) error {
			slog.Warn("SERVER_ERROR", slog.String("error", err.Error()))

			ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
			return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		},
	})

	// Swagger documentation
	app.Get("/swagger/*", swaggo.HandlerDefault)

	return app
}
