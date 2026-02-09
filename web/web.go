package web

import (
	"github.com/gofiber/fiber/v3"
)

func SetupWeb(app *fiber.App) {
	users := app.Group("/api/users")
	setupUserController(users)
}
