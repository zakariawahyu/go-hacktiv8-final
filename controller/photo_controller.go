package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-hacktiv8-final/common/dto"
	"github.com/zakariawahyu/go-hacktiv8-final/common/response"
	"github.com/zakariawahyu/go-hacktiv8-final/exception"
	"github.com/zakariawahyu/go-hacktiv8-final/services"
)

type PhotoController struct {
	photoServices services.PhotoServices
	jwtServices   services.JWTServices
	userServices  services.UserServices
}

func NewPhotoController(photoServices services.PhotoServices, jwtServices services.JWTServices, userServices services.UserServices) PhotoController {
	return PhotoController{
		photoServices: photoServices,
		jwtServices:   jwtServices,
		userServices:  userServices,
	}
}

func (controller *PhotoController) Routes(app *fiber.App) {
	app.Post("/photos", controller.Create)
	app.Get("/photos", controller.GetAllTask)
}

func (controller *PhotoController) Create(ctx *fiber.Ctx) error {
	var request dto.PhotoRequest

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	claims := controller.jwtServices.GetClaimsJWT(ctx)
	email := fmt.Sprintf("%v", claims["email"])
	user := controller.userServices.FindUserByEmail(email)
	request.UserID = user.ID

	photo := controller.photoServices.CreatePhoto(request)
	res := response.BuildSuccessResponse(fiber.StatusCreated, "Success", photo)
	return ctx.Status(fiber.StatusCreated).JSON(res)
}

func (controller *PhotoController) GetAllTask(ctx *fiber.Ctx) error {
	claims := controller.jwtServices.GetClaimsJWT(ctx)
	email := fmt.Sprintf("%v", claims["email"])
	user := controller.userServices.FindUserByEmail(email)

	photo := controller.photoServices.AllPhoto(user.ID)
	res := response.BuildSuccessResponse(fiber.StatusOK, "Success", photo)
	return ctx.JSON(res)
}
