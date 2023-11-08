package controllers

import (
	"github.com/gofiber/fiber/v2"
	useCase "github.com/intwone/eda-arch-golang/internal/public_auth/application/use_cases"
	"github.com/intwone/eda-arch-golang/internal/public_auth/presentation/inputs"
)

type AuthenticateController struct {
	AuthenticateUseCase useCase.AuthenticateUseCaseInterface
}

func NewAuthenticateController(AuthenticateUsecase useCase.AuthenticateUseCaseInterface) *AuthenticateController {
	c := AuthenticateController{
		AuthenticateUseCase: AuthenticateUsecase,
	}

	return &c
}

func (ac *AuthenticateController) Handle(c *fiber.Ctx) error {
	var req inputs.AuthenticateInput
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	result, err := ac.AuthenticateUseCase.Execute(useCase.AuthenticateInput{Email: req.Email, Password: req.Password})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": result})
}
