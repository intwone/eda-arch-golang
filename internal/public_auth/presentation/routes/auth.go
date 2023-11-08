package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/intwone/eda-arch-golang/internal/public_auth/presentation/controllers"
)

func SetupRoutes(a *fiber.App, c controllers.AuthControllers) {
	group := a.Group("/api/users")
	group.Post("/auth-create", c.AuthCreateController.Handle)
	group.Post("/authenticate", c.AuthenticateController.Handle)
}
