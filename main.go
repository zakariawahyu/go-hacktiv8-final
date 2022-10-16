package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/zakariawahyu/go-hacktiv8-final/config"
	"github.com/zakariawahyu/go-hacktiv8-final/controller"
	"github.com/zakariawahyu/go-hacktiv8-final/repository"
	"github.com/zakariawahyu/go-hacktiv8-final/services"
)

func main() {
	db := config.DatabaseConnection()
	userRepository := repository.NewUserRepository(db)
	photoRepository := repository.NewPhotoRepository(db)
	jwtServices := services.NewJWTServices()
	userServices := services.NewUserServices(userRepository)
	photoServices := services.NewPhotoServices(photoRepository)
	userController := controller.NewUserController(userServices, jwtServices)
	photoController := controller.NewPhotoController(photoServices, jwtServices, userServices)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello world!",
		})
	})
	userController.Routes(app)
	photoController.Routes(app)
	app.Listen(":8081")
}
