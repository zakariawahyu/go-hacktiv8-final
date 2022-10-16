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
	UpdateUser(request dto.UpdateUserRequest) response.UserResponse
}

type PhotoServices interface {
	CreatePhoto(request dto.PhotoRequest) response.PhotoResponse
	AllPhoto(userID int64) []response.PhotoResponseAll
	UpdatePhoto(request dto.UpdatePhotoRequest) response.PhotoResponse
}

type CommentServices interface {
	CreateComment(request dto.CommentRequest) response.CommentResponse
	AllComment(userID int64) []response.CommentResponseAll
	UpdateComment(request dto.UpdateCommentRequest) response.CommentResponse
}

type SocialMediaServices interface {
	CreateSocialMedia(request dto.SocialMediaRequest) response.SocialMediaResponse
	AllSocialMedia(userID int64) []response.SocialMediaResponseAll
	UpdateSocialMedia(request dto.UpdateSocialMediaRequest) response.SocialMediaResponse
}
