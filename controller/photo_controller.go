package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-hacktiv8-final/common/dto"
	"github.com/zakariawahyu/go-hacktiv8-final/common/response"
	"github.com/zakariawahyu/go-hacktiv8-final/exception"
	"github.com/zakariawahyu/go-hacktiv8-final/middleware"
	"github.com/zakariawahyu/go-hacktiv8-final/services"
	"strconv"
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
	app.Post("/photos", middleware.AuthorizeJWT(controller.jwtServices), controller.Create)
	app.Get("/photos", middleware.AuthorizeJWT(controller.jwtServices), controller.GetAllPhoto)
	app.Put("/photos/:id", middleware.AuthorizeJWT(controller.jwtServices), controller.Update)
	app.Delete("/photos/:id", middleware.AuthorizeJWT(controller.jwtServices), controller.Delete)
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

func (controller *PhotoController) GetAllPhoto(ctx *fiber.Ctx) error {
	claims := controller.jwtServices.GetClaimsJWT(ctx)
	email := fmt.Sprintf("%v", claims["email"])
	user := controller.userServices.FindUserByEmail(email)

	photo := controller.photoServices.AllPhoto(user.ID)
	res := response.BuildSuccessResponse(fiber.StatusOK, "Success", photo)
	return ctx.JSON(res)
}

func (controller *PhotoController) Update(ctx *fiber.Ctx) error {
	var request dto.UpdatePhotoRequest
	id, _ := strconv.Atoi(ctx.Params("id"))

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	claims := controller.jwtServices.GetClaimsJWT(ctx)
	email := fmt.Sprintf("%v", claims["email"])
	findUser := controller.userServices.FindUserByEmail(email)
	request.ID = int64(id)
	request.UserID = findUser.ID

	photo := controller.photoServices.UpdatePhoto(request)
	res := response.BuildSuccessResponse(fiber.StatusOK, "Success", photo)
	return ctx.JSON(res)
}

func (controller *PhotoController) Delete(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	claims := controller.jwtServices.GetClaimsJWT(ctx)
	email := fmt.Sprintf("%v", claims["email"])
	findUser := controller.userServices.FindUserByEmail(email)

	_ = controller.photoServices.DeleteById(int64(id), findUser.ID)
	res := response.BuildSuccessResponse(fiber.StatusOK, "Success", fiber.Map{
		"message": "Your photo has been successfully deleted",
	})
	return ctx.JSON(res)
}
