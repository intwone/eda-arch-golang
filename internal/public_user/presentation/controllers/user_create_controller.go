package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	useCase "github.com/intwone/eda-arch-golang/internal/public_user/application/use_cases"
	"github.com/intwone/eda-arch-golang/internal/public_user/presentation/inputs"
	"github.com/intwone/eda-arch-golang/internal/public_user/presentation/outputs"
)

type UserCreateController struct {
	UserCreateUseCase useCase.UserCreateUseCaseInterface
}

func NewUserCreateController(userCreateUsecase useCase.UserCreateUseCaseInterface) *UserCreateController {
	c := UserCreateController{
		UserCreateUseCase: userCreateUsecase,
	}
	return &c
}

func (cc *UserCreateController) Handle(c *fiber.Ctx) error {
	var body inputs.CreateUserInput
	if pErr := c.BodyParser(&body); pErr != nil {
		fmt.Println(pErr)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": pErr.Error()})
	}
	result, uErr := cc.UserCreateUseCase.Execute(useCase.UserCreateInput{Name: body.Name, Email: body.Email, Cpf: body.Cpf, Birthdate: body.Birthdate})
	if uErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "internal server error"})
	}
	output := outputs.CreateUserOutputMapper(result.User)
	return c.JSON(fiber.Map{"data": output})
}
