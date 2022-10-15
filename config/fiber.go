package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-hacktiv8-final/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
