package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-hacktiv8-final/common/dto"
	"github.com/zakariawahyu/go-hacktiv8-final/common/response"
	"github.com/zakariawahyu/go-hacktiv8-final/exception"
	"github.com/zakariawahyu/go-hacktiv8-final/services"
)

type CommentController struct {
	commentServices services.CommentServices
	jwtServices     services.JWTServices
	userServices    services.UserServices
}

func NewCommentController(commentServices services.CommentServices, jwtServices services.JWTServices, userServices services.UserServices) CommentController {
	return CommentController{
		commentServices: commentServices,
		jwtServices:     jwtServices,
		userServices:    userServices,
	}
}

func (controller *CommentController) Routes(app *fiber.App) {
	app.Post("/comments", controller.Create)
	app.Get("/comments", controller.GetAllTask)
}

func (controller *CommentController) Create(ctx *fiber.Ctx) error {
	var request dto.CommentRequest

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	claims := controller.jwtServices.GetClaimsJWT(ctx)
	email := fmt.Sprintf("%v", claims["email"])
	user := controller.userServices.FindUserByEmail(email)
	request.UserID = user.ID

	comment := controller.commentServices.CreatePhoto(request)
	res := response.BuildSuccessResponse(fiber.StatusCreated, "Success", comment)
	return ctx.Status(fiber.StatusCreated).JSON(res)
}

func (controller *CommentController) GetAllTask(ctx *fiber.Ctx) error {
	claims := controller.jwtServices.GetClaimsJWT(ctx)
	email := fmt.Sprintf("%v", claims["email"])
	user := controller.userServices.FindUserByEmail(email)

	comment := controller.commentServices.AllPhoto(user.ID)
	res := response.BuildSuccessResponse(fiber.StatusOK, "Success", comment)
	return ctx.JSON(res)
}
