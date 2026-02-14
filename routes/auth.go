package routes

import (
	"wallet-app/controller"

	"github.com/gofiber/fiber/v3"
)

func SetupAuthRoutes(app fiber.Router, authController *controller.AuthController) {
	authRoute := app.Group("/auth")
	authRoute.Post("/register", authController.Register)
	authRoute.Post("/login", authController.Login)
}
