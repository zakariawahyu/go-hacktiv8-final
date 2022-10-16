package services

import (
	"github.com/mashingan/smapping"
	"github.com/zakariawahyu/go-hacktiv8-final/common/dto"
	"github.com/zakariawahyu/go-hacktiv8-final/common/response"
	"github.com/zakariawahyu/go-hacktiv8-final/entity"
	"github.com/zakariawahyu/go-hacktiv8-final/exception"
	"github.com/zakariawahyu/go-hacktiv8-final/repository"
	"github.com/zakariawahyu/go-hacktiv8-final/validations"
)

type CommentServicesImpl struct {
	commentRepo repository.CommentRepository
}

func NewCommentServices(commentRepository repository.CommentRepository) CommentServices {
	return &CommentServicesImpl{
		commentRepo: commentRepository,
	}
}

func (services *CommentServicesImpl) CreatePhoto(request dto.CommentRequest) response.CommentResponse {
	var comment entity.Comment

	errValidate := validations.ValidateComment(request)
	exception.PanicIfNeeded(errValidate)

	errMap := smapping.FillStruct(&comment, smapping.MapFields(&request))
	exception.PanicIfNeeded(errMap)

	result := services.commentRepo.Create(comment)
	res := response.NewCommentResponse(result)

	return res
}

func (services *CommentServicesImpl) AllPhoto(userID int64) []response.CommentResponseAll {
	result := services.commentRepo.GetAll(userID)
	res := response.NewCommentResponseArray(result)

	return res
}
