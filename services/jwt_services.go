package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-hacktiv8-final/exception"
	"os"
	"strings"
	"time"
)

type JWTServicesImpl struct {
	issuer    string
	secretKey string
}

type jwtCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func NewJWTServices() JWTServices {
	return &JWTServicesImpl{
		issuer:    "admin",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	key := os.Getenv("JWT_SECRET")

	return key
}

func (jwtServices *JWTServicesImpl) GenerateToken(email string) string {
	claims := &jwtCustomClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, 1).Unix(),
			Issuer:    jwtServices.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokens.SignedString([]byte(jwtServices.secretKey))
	exception.PanicIfNeeded(err)

	return token
}

func (jwtServices *JWTServicesImpl) ValidateToken(tokens string, ctx *fiber.Ctx) *jwt.Token {
	t, err := jwt.Parse(tokens, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return []byte(jwtServices.secretKey), nil
	})

	if err != nil {
		return nil
	}

	return t
}

func (jwtServices *JWTServicesImpl) GetClaimsJWT(ctx *fiber.Ctx) jwt.MapClaims {
	header := ctx.Get("Authorization")
	tokenArr := strings.Split(header, "Bearer ")

	if len(tokenArr) != 2 {
		exception.PanicIfNeeded("Unauthorized")
	}

	tokenStr := tokenArr[1]
	token := jwtServices.ValidateToken(tokenStr, ctx)
	claims := token.Claims.(jwt.MapClaims)

	return claims
}
