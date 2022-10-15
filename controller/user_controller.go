package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-hacktiv8-final/common/dto"
	"github.com/zakariawahyu/go-hacktiv8-final/common/response"
	"github.com/zakariawahyu/go-hacktiv8-final/exception"
	"github.com/zakariawahyu/go-hacktiv8-final/services"
)

type UserController struct {
	userServices services.UserServices
	jwtServices  services.JWTServices
}

func NewUserController(userServices services.UserServices, jwtServices services.JWTServices) UserController {
	return UserController{
		userServices: userServices,
		jwtServices:  jwtServices,
	}
}

func (controller *UserController) Routes(app *fiber.App) {
	app.Post("/users/register", controller.Register)
	app.Post("/users/login", controller.Login)
}
func (controller *UserController) Register(ctx *fiber.Ctx) error {
	var request dto.RegisterRequest

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	user := controller.userServices.RegisterUser(request)

	res := response.BuildSuccessResponse(fiber.StatusCreated, "Registered", user)
	return ctx.Status(fiber.StatusCreated).JSON(res)
}

func (controller *UserController) Login(ctx *fiber.Ctx) error {
	var request dto.LoginRequest

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	user := controller.userServices.LoginUser(request)
	token := controller.jwtServices.GenerateToken(user.Email)

	res := response.BuildSuccessResponse(fiber.StatusOK, "Success", fiber.Map{
		"token": token,
	})
	return ctx.JSON(res)
}
