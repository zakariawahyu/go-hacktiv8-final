package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-hacktiv8-final/exception"
	"github.com/zakariawahyu/go-hacktiv8-final/services"
	"log"
	"strings"
)

func AuthorizeJWT(jwtServices services.JWTServices) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		tokenArr := strings.Split(authHeader, "Bearer ")

		if len(tokenArr) != 2 {
			exception.PanicIfNeeded("No token provided")
		}

		tokenStr := tokenArr[1]
		token := jwtServices.ValidateToken(tokenStr, ctx)
		if token != nil {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[email]: ", claims["email"])
			log.Println("Claim[issuer] :", claims["issuer"])
			return ctx.Next()
		} else {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Your token is not valid",
			})
		}
	}
}
