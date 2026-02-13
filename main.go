package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"wallet-app/db"

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

func main() {
	app := fiber.New()
	APP_PORT, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		APP_PORT = 3000
	}

	func() {

		defer db.StopDBConnection()
		db.InitDBConnection()

		app.Get("/", func(c fiber.Ctx) error {
			return c.SendString("Hello, World!")
		})

		log.Fatal(app.Listen(fmt.Sprintf(":%d", APP_PORT)))

	}()

}
