package controllers

import (
	"github.com/gofiber/fiber/v2"
	uc "github.com/intwone/eda-arch-golang/internal/domain/auth/use_cases"
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
	result := ac.AuthCreateUseCase.Execute(uc.AuthCreateInput{Email: "cassio@gmail.com"})

	return c.JSON(fiber.Map{"data": result})
}
