package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-hacktiv8-final/config"
	"net/http"
)

func main() {
	// Setup database
	db := config.DatabaseConnection()
	config.CloseDatabaseConnection(db)

	// Setup Fiber
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Hello world!",
		})
	})

	app.Listen(":8081")
}
