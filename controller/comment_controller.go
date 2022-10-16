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
	app.Post("/comments", middleware.AuthorizeJWT(controller.jwtServices), controller.Create)
	app.Get("/comments", middleware.AuthorizeJWT(controller.jwtServices), controller.GetAllComment)
	app.Put("/comments/:id", middleware.AuthorizeJWT(controller.jwtServices), controller.Update)
}

func (controller *CommentController) Create(ctx *fiber.Ctx) error {
	var request dto.CommentRequest

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	claims := controller.jwtServices.GetClaimsJWT(ctx)
	email := fmt.Sprintf("%v", claims["email"])
	user := controller.userServices.FindUserByEmail(email)
	request.UserID = user.ID

	comment := controller.commentServices.CreateComment(request)
	res := response.BuildSuccessResponse(fiber.StatusCreated, "Success", comment)
	return ctx.Status(fiber.StatusCreated).JSON(res)
}

func (controller *CommentController) GetAllComment(ctx *fiber.Ctx) error {
	claims := controller.jwtServices.GetClaimsJWT(ctx)
	email := fmt.Sprintf("%v", claims["email"])
	user := controller.userServices.FindUserByEmail(email)

	comment := controller.commentServices.AllComment(user.ID)
	res := response.BuildSuccessResponse(fiber.StatusOK, "Success", comment)
	return ctx.JSON(res)
}

func (controller *CommentController) Update(ctx *fiber.Ctx) error {
	var request dto.UpdateCommentRequest
	id, _ := strconv.Atoi(ctx.Params("id"))

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	claims := controller.jwtServices.GetClaimsJWT(ctx)
	email := fmt.Sprintf("%v", claims["email"])
	findUser := controller.userServices.FindUserByEmail(email)
	request.ID = int64(id)
	request.UserID = findUser.ID

	comment := controller.commentServices.UpdateComment(request)
	res := response.BuildSuccessResponse(fiber.StatusOK, "Success", comment)
	return ctx.JSON(res)
}
