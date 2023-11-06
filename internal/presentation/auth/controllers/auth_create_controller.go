package controllers

import (
	"github.com/gofiber/fiber/v2"
	useCase "github.com/intwone/eda-arch-golang/internal/domain/modules/auth/use_cases"
	models "github.com/intwone/eda-arch-golang/internal/presentation/auth/models"
)

type AuthCreateController struct {
	AuthCreateUseCase useCase.AuthCreateUseCaseInterface
}

func NewAuthCreateController(authCreateUsecase useCase.AuthCreateUseCaseInterface) *AuthCreateController {
	c := AuthCreateController{
		AuthCreateUseCase: authCreateUsecase,
	}

	return &c
}

func (ac *AuthCreateController) Handle(c *fiber.Ctx) error {
	var req models.AuthCreateRequestModel

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid email"})
	}

	result := ac.AuthCreateUseCase.Execute(useCase.AuthCreateInput{Email: req.Value})

	return c.JSON(fiber.Map{"data": result})
}
