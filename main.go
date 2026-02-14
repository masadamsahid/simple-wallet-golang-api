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

	swagger "github.com/gofiber/contrib/v3/swaggerui"
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
		APP_PORT = 8080
	}

	helpers.InitJWT()

	defer db.StopDBConnection()
	db.InitDBConnection()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//
	swaggerCfg := swagger.Config{
		BasePath: "/", // swagger ui base path
		FilePath: "./docs/openapi.yaml",
		Path:     "/swagger",
	}
	app.Use(swagger.New(swaggerCfg))

	api := app.Group("/api")

	// /api/auth
	userRepo := repository.NewUserRepository(db.DB)
	authService := service.NewAuthService(userRepo)
	authController := controller.NewAuthController(authService)
	routes.SetupAuthRoutes(api, authController)

	// /api/users
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	routes.SetupUsersRoutes(api, userController)

	// /api/transactions
	transactionRepo := repository.NewTransactionRepository(db.DB)
	transactionService := service.NewTransactionService(transactionRepo, userRepo, db.DB)
	transactionController := controller.NewTransactionController(transactionService)
	routes.SetupTransactionsRoutes(api, transactionController)

	fmt.Println("Available Routes:")
	for _, route := range app.GetRoutes() {
		fmt.Printf("%s\t%s\n", route.Method, route.Path)
	}

	log.Fatal(app.Listen(fmt.Sprintf(":%d", APP_PORT)))

}
