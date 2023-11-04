package routes

import (
	"github.com/gofiber/fiber/v2"
	c "github.com/intwone/eda-arch-golang/internal/presentation"
)

func SetupRoutes(a *fiber.App, c c.AuthControllers) {
	group := a.Group("/api/users")
	group.Post("/auth-create", c.AuthCreateController.Handle)
}
