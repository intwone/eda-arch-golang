package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/intwone/eda-arch-golang/internal/private_shared/err"
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
	var body inputs.AuthenticateInput
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	result, ucErr := ac.AuthenticateUseCase.Execute(useCase.AuthenticateInput{Email: body.Email, Password: body.Password})
	if ucErr != nil {
		errMsg, pErr := err.ParseErrorResponse(ucErr.Error())
		if pErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "internal server error"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": errMsg})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": result.Token})
}
