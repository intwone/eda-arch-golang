package controllers

import (
	"github.com/gofiber/fiber/v2"
	uc "github.com/intwone/eda-arch-golang/internal/domain/auth/use_cases"
	"github.com/intwone/eda-arch-golang/internal/presentation/dtos"
)

type AuthCreateController struct {
	AuthCreateUseCase uc.AuthCreateUseCaseInterface
}

func NewAuthCreateController(authCreateUsecase uc.AuthCreateUseCaseInterface) *AuthCreateController {
	c := AuthCreateController{
		AuthCreateUseCase: authCreateUsecase,
	}

	return &c
}

func (ac *AuthCreateController) Handle(c *fiber.Ctx) error {
	var req dtos.AuthCreateRequestDTO

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid email"})
	}

	result := ac.AuthCreateUseCase.Execute(uc.AuthCreateInput{Email: req.Value})

	return c.JSON(fiber.Map{"data": result})
}
