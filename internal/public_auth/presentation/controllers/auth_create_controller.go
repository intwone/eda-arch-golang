package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/intwone/eda-arch-golang/internal/private_shared/err"
	useCase "github.com/intwone/eda-arch-golang/internal/public_auth/application/use_cases"
	"github.com/intwone/eda-arch-golang/internal/public_auth/presentation/inputs"
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
	var body inputs.AuthCreateInput
	if pErr := c.BodyParser(&body); pErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "incorrect fields"})
	}
	ucErr := ac.AuthCreateUseCase.Execute(useCase.AuthCreateInput{Email: body.Value})
	if ucErr != nil {
		errMsg, pErr := err.ParseErrorResponse(ucErr.Error())
		if pErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "internal server error"})
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": errMsg})
	}
	return c.JSON(fiber.Map{"data": true})
}
