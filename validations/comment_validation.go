package validations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/zakariawahyu/go-hacktiv8-final/common/dto"
)

func ValidateComment(comment dto.CommentRequest) error {
	return validation.ValidateStruct(&comment,
		validation.Field(&comment.Message, validation.Required),
		validation.Field(&comment.PhotoID, validation.Required))
}

func ValidateUpdateComment(comment dto.UpdateCommentRequest) error {
	return validation.ValidateStruct(&comment,
		validation.Field(&comment.Message, validation.Required),
		validation.Field(&comment.PhotoID, validation.Required))
}
