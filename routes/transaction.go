package routes

import (
	"wallet-app/controller"
	"wallet-app/middlewares"

	"github.com/gofiber/fiber/v3"
)

func SetupTransactionsRoutes(app fiber.Router, transactionController *controller.TransactionController) {
	transactionRoute := app.Group("/transactions", middlewares.IsAuthenticated)

	transactionRoute.Post("/deposit", transactionController.Deposit)
	transactionRoute.Post("/withdraw", transactionController.Withdraw)
	transactionRoute.Get("/", transactionController.GetHistory)
}
