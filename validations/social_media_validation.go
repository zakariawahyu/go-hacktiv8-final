package validations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/zakariawahyu/go-hacktiv8-final/common/dto"
)

func ValidateSocialMedia(socialMedia dto.SocialMediaRequest) error {
	return validation.ValidateStruct(&socialMedia,
		validation.Field(&socialMedia.Name, validation.Required),
		validation.Field(&socialMedia.SocialMediaUrl, validation.Required))
}
