package web

import (
	"errors"
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"github.com/mtproject-ru/server/services"
	"github.com/mtproject-ru/server/web/dto"
)

func setupUserController(router fiber.Router) {
	router.Get("/:username", getUserByUsernameHandler)
}

func getUserByUsernameHandler(ctx fiber.Ctx) error {
	username := ctx.Params("username")
	if username == "" {
		err := errors.New("Не указано имя пользователя для поиска")
		return ctx.Status(fiber.StatusBadRequest).JSON(NewErrResponse(err))
	}

	userService := fiber.MustGetState[services.UserService](ctx.App().State(), services.USER_SERVICE_KEY)
	user, err := userService.GetUserByUsername(username)

	if err != nil {
		slog.Error(err.Error())
		return ctx.Status(fiber.StatusNotFound).JSON(NewErrResponse(err))
	}

	return ctx.Status(fiber.StatusOK).JSON(NewOkResponse(
		dto.ToUserResponse(user),
	))
}
