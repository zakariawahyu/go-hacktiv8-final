package services

import (
	"github.com/mashingan/smapping"
	"github.com/zakariawahyu/go-hacktiv8-final/common/dto"
	"github.com/zakariawahyu/go-hacktiv8-final/common/response"
	"github.com/zakariawahyu/go-hacktiv8-final/entity"
	"github.com/zakariawahyu/go-hacktiv8-final/exception"
	"github.com/zakariawahyu/go-hacktiv8-final/repository"
	"github.com/zakariawahyu/go-hacktiv8-final/validations"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserServicesImpl struct {
	userRepo repository.UserRepository
}

func NewUserServices(userRepo repository.UserRepository) UserServices {
	return &UserServicesImpl{
		userRepo: userRepo,
	}
}

func (services *UserServicesImpl) RegisterUser(request dto.RegisterRequest) response.UserResponse {
	var user entity.User

	errValidate := validations.ValidateRegisterUser(request)
	exception.PanicIfNeeded(errValidate)

	errMap := smapping.FillStruct(&user, smapping.MapFields(&request))
	exception.PanicIfNeeded(errMap)

	result := services.userRepo.Create(user)
	res := response.NewUserResponse(result)

	return res
}

func (services *UserServicesImpl) LoginUser(request dto.LoginRequest) response.UserResponse {
	errValidate := validations.ValidateLoginUser(request)
	exception.PanicIfNeeded(errValidate)

	result := services.userRepo.FindByEmail(request.Email)
	isValidPass := comparePassword(result.Password, []byte(request.Password))
	if !isValidPass {
		exception.PanicIfNeeded("wrong username and password")
	}
	res := response.NewUserResponse(result)

	return res
}

func comparePassword(hashPass string, plainPass []byte) bool {
	byteHash := []byte(hashPass)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPass)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

func (services *UserServicesImpl) FindUserByEmail(email string) response.UserResponse {
	result := services.userRepo.FindByEmail(email)
	res := response.NewUserResponse(result)

	return res
}

func (services *UserServicesImpl) UpdateUser(request dto.UpdateUserRequest) response.UserResponse {
	var user entity.User

	errValidate := validations.ValidateUpdateUser(request)
	exception.PanicIfNeeded(errValidate)

	errMap := smapping.FillStruct(&user, smapping.MapFields(&request))
	exception.PanicIfNeeded(errMap)

	result := services.userRepo.Update(user)
	res := response.NewUserResponse(result)

	return res
}
