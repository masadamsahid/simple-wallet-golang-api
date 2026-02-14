package controller

import (
	"wallet-app/dto"
	"wallet-app/helpers"
	"wallet-app/service"

	"github.com/gofiber/fiber/v3"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{userService: userService}
}

func (u *UserController) GetBalance(c fiber.Ctx) error {
	user, ok := c.Locals("user").(*helpers.AuthTokenClaims)
	if !ok {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(&dto.ResponseGeneric[any]{
			Message: "Unauthorized",
		})
	}

	userData, err := u.userService.GetUserByID(user.ID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(&dto.ResponseGeneric[any]{
			Message: "Failed to fetch user balance",
		})
	}

	return c.JSON(&dto.ResponseGeneric[fiber.Map]{
		Message: "Success",
		Data: &fiber.Map{
			"id":      userData.ID,
			"email":   userData.Email,
			"balance": userData.Balance,
		},
	})
}
