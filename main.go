package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/zakariawahyu/go-hacktiv8-final/config"
)

func main() {
	// Setup database
	db := config.DatabaseConnection()
	config.CloseDatabaseConnection(db)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello world!",
		})
	})

	app.Listen(":8081")
}
