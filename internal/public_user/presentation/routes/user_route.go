package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/intwone/eda-arch-golang/internal/public_user/presentation/controllers"
)

func SetupUserRoutes(a *fiber.App, c controllers.UserControllers) {
	group := a.Group("/api/users")
	group.Post("/create", c.UserCreateController.Handle)
}
