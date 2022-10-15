package exception

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-hacktiv8-final/common/response"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(ValidationErr)
	if ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return ctx.Status(code).JSON(response.ErrorResponse{
		Status:  code,
		Message: err.Error(),
	})
}
