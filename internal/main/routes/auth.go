package routes

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/intwone/eda-arch-golang/internal/presentation"
)

func SetupRoutes(a *fiber.App, c controllers.AuthControllers) {
	group := a.Group("/api/users")
	group.Post("/auth-create", c.AuthCreateController.Handle)
}
