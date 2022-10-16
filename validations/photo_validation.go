package validations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/zakariawahyu/go-hacktiv8-final/common/dto"
)

func ValidatePhoto(photo dto.PhotoRequest) error {
	return validation.ValidateStruct(&photo,
		validation.Field(&photo.Title, validation.Required),
		validation.Field(&photo.PhotoUrl, validation.Required))
}

func ValidateUpdatePhoto(photo dto.UpdatePhotoRequest) error {
	return validation.ValidateStruct(&photo,
		validation.Field(&photo.Title, validation.Required),
		validation.Field(&photo.PhotoUrl, validation.Required))
}
