package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-hacktiv8-final/common/dto"
	"github.com/zakariawahyu/go-hacktiv8-final/common/response"
	"github.com/zakariawahyu/go-hacktiv8-final/exception"
	"github.com/zakariawahyu/go-hacktiv8-final/services"
)

type SocialMediaController struct {
	socialMediaServices services.SocialMediaServices
	jwtServices         services.JWTServices
	userServices        services.UserServices
}

func NewSocialMediaController(socialMediaServices services.SocialMediaServices, jwtServices services.JWTServices, userServices services.UserServices) SocialMediaController {
	return SocialMediaController{
		socialMediaServices: socialMediaServices,
		jwtServices:         jwtServices,
		userServices:        userServices,
	}
}

func (controller *SocialMediaController) Routes(app *fiber.App) {
	app.Post("/socialmedias", controller.Create)
	app.Get("/socialmedias", controller.GetAllSocialMedia)
}

func (controller *SocialMediaController) Create(ctx *fiber.Ctx) error {
	var request dto.SocialMediaRequest

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	claims := controller.jwtServices.GetClaimsJWT(ctx)
	email := fmt.Sprintf("%v", claims["email"])
	user := controller.userServices.FindUserByEmail(email)
	request.UserID = user.ID

	socialMedia := controller.socialMediaServices.CreateSocialMedia(request)
	res := response.BuildSuccessResponse(fiber.StatusCreated, "Success", socialMedia)
	return ctx.Status(fiber.StatusCreated).JSON(res)
}

func (controller *SocialMediaController) GetAllSocialMedia(ctx *fiber.Ctx) error {
	claims := controller.jwtServices.GetClaimsJWT(ctx)
	email := fmt.Sprintf("%v", claims["email"])
	user := controller.userServices.FindUserByEmail(email)

	socialMedia := controller.socialMediaServices.AllSocialMedia(user.ID)
	res := response.BuildSuccessResponse(fiber.StatusOK, "Success", socialMedia)
	return ctx.JSON(res)
}
