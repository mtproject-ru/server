package services

import "github.com/gofiber/fiber/v3"

func SetupServices(app *fiber.App) {
	userService := &UserServiceImpl{}

	app.State().Set(USER_SERVICE_KEY, userService)
}
