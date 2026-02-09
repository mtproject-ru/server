package services

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mtproject-ru/server/database/repositories"
)

func SetupServices(app *fiber.App) {
	userRepository := &repositories.UserRepositoryImpl{}
	userService := &UserServiceImpl{
		repo: userRepository,
	}

	app.State().Set(USER_SERVICE_KEY, userService)
}
