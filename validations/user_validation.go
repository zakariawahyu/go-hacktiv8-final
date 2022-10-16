package validations

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/zakariawahyu/go-hacktiv8-final/common/dto"
)

func ValidateUpdateUser(user dto.UpdateUserRequest) error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Password, validation.Length(6, 0)),
		validation.Field(&user.Age, validation.By(checkAge(user.Age))))
}

func ValidateRegisterUser(user dto.RegisterRequest) error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Username, validation.Required),
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(6, 0)),
		validation.Field(&user.Age, validation.Required, validation.By(checkAge(user.Age))))
}

func ValidateLoginUser(user dto.LoginRequest) error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(6, 0)))
}

func checkAge(age int) validation.RuleFunc {
	return func(value interface{}) error {
		if age < 8 {
			return errors.New("age must be above 8")
		}
		return nil
	}
}
