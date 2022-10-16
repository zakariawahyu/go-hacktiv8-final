package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-hacktiv8-final/common/dto"
	"github.com/zakariawahyu/go-hacktiv8-final/common/response"
)

type JWTServices interface {
	GenerateToken(email string) string
	ValidateToken(tokens string, ctx *fiber.Ctx) *jwt.Token
	GetClaimsJWT(ctx *fiber.Ctx) jwt.MapClaims
}

type UserServices interface {
	RegisterUser(request dto.RegisterRequest) response.UserResponse
	LoginUser(request dto.LoginRequest) response.UserResponse
	FindUserByEmail(email string) response.UserResponse
}

type PhotoServices interface {
	CreatePhoto(request dto.PhotoRequest) response.PhotoResponse
	AllPhoto(userID int64) []response.PhotoResponseAll
}
