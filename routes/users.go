package routes

import (
	"wallet-app/controller"
	"wallet-app/middlewares"

	"github.com/gofiber/fiber/v3"
)

func SetupUsersRoutes(app fiber.Router, userController *controller.UserController) {
	userRoute := app.Group("/users")

	userRoute.Get("/balance", middlewares.IsAuthenticated, userController.GetBalance)
}
