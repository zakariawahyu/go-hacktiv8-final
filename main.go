package main

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func main() {
	// Setup Fiber
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Hello world!",
		})
	})

	app.Listen(":8081")
}
