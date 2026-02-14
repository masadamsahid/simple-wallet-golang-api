package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"wallet-app/controller"
	"wallet-app/db"
	"wallet-app/helpers"
	"wallet-app/repository"
	"wallet-app/routes"
	"wallet-app/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables from OS")
	} else {
		log.Println("Loading .env success")
	}
}

type structValidator struct {
	validate *validator.Validate
}

// Validator needs to implement the Validate method
func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

func main() {
	app := fiber.New(fiber.Config{
		StructValidator: &structValidator{
			validate: validator.New(),
		},
	})
	APP_PORT, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		APP_PORT = 3000
	}

	helpers.InitJWT()

	defer db.StopDBConnection()
	db.InitDBConnection()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")

	userRepo := repository.NewUserRepository(db.DB)
	authService := service.NewUserService(userRepo)
	authController := controller.NewAuthController(authService)
	routes.SetupAuthRoutes(api, authController)

	fmt.Println("Available Routes:")
	for _, route := range app.GetRoutes() {
		fmt.Printf("%s\t%s\n", route.Method, route.Path)
	}

	log.Fatal(app.Listen(fmt.Sprintf(":%d", APP_PORT)))

}
