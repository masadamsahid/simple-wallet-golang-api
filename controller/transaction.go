package controller

import (
	"wallet-app/dto"
	"wallet-app/helpers"
	"wallet-app/model"
	"wallet-app/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type TransactionController struct {
	transactionService service.ITransactionService
}

func NewTransactionController(transactionService service.ITransactionService) *TransactionController {
	return &TransactionController{transactionService: transactionService}
}

func (t *TransactionController) Deposit(c fiber.Ctx) error {
	user, ok := c.Locals("user").(*helpers.AuthTokenClaims)
	if !ok {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(&dto.ResponseGeneric[any]{
			Message: "Unauthorized",
		})
	}

	body := new(dto.TransactionRequest)

	if err := c.Bind().Body(body); err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(&dto.ResponseGeneric[any]{
				Message: "Invalid request body",
				Errors:  err.Error(),
			})
		}

		return c.JSON(&dto.ResponseGeneric[any]{
			Message: "Failed",
			Errors:  helpers.HandleValidationErrors(validationErrors),
		})
	}

	transaction, err := t.transactionService.Deposit(user.ID, body.Amount)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(&dto.ResponseGeneric[any]{
			Message: err.Error(),
		})
	}

	return c.JSON(&dto.ResponseGeneric[*model.Transaction]{
		Message: "Deposit successful",
		Data:    &transaction,
	})
}

func (t *TransactionController) Withdraw(c fiber.Ctx) error {
	user, ok := c.Locals("user").(*helpers.AuthTokenClaims)
	if !ok {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(&dto.ResponseGeneric[any]{
			Message: "Unauthorized",
		})
	}

	body := new(dto.TransactionRequest)

	if err := c.Bind().Body(body); err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(&dto.ResponseGeneric[any]{
				Message: "Invalid request body",
				Errors:  err.Error(),
			})
		}

		return c.JSON(&dto.ResponseGeneric[any]{
			Message: "Failed",
			Errors:  helpers.HandleValidationErrors(validationErrors),
		})
	}

	transaction, err := t.transactionService.Withdraw(user.ID, body.Amount)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(&dto.ResponseGeneric[any]{
			Message: err.Error(),
		})
	}

	return c.JSON(&dto.ResponseGeneric[*model.Transaction]{
		Message: "Withdrawal successful",
		Data:    &transaction,
	})
}

func (t *TransactionController) GetHistory(c fiber.Ctx) error {
	user, ok := c.Locals("user").(*helpers.AuthTokenClaims)
	if !ok {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(&dto.ResponseGeneric[any]{
			Message: "Unauthorized",
		})
	}

	history, err := t.transactionService.GetHistory(user.ID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(&dto.ResponseGeneric[any]{
			Message: "Failed to fetch transaction history",
		})
	}

	return c.JSON(&dto.ResponseGeneric[[]model.Transaction]{
		Message: "Success",
		Data:    &history,
	})
}
