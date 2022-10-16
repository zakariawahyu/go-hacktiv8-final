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

func (services *CommentServicesImpl) CreateComment(request dto.CommentRequest) response.CommentResponse {
	var comment entity.Comment

	errValidate := validations.ValidateComment(request)
	exception.PanicIfNeeded(errValidate)

	errMap := smapping.FillStruct(&comment, smapping.MapFields(&request))
	exception.PanicIfNeeded(errMap)

	result := services.commentRepo.Create(comment)
	res := response.NewCommentResponse(result)

	return res
}

func (services *CommentServicesImpl) AllComment(userID int64) []response.CommentResponseAll {
	result := services.commentRepo.GetAll(userID)
	res := response.NewCommentResponseArray(result)

	return res
}

func (services *CommentServicesImpl) UpdateComment(request dto.UpdateCommentRequest) response.CommentResponse {
	var comment entity.Comment

	errValidate := validations.ValidateUpdateComment(request)
	exception.PanicIfNeeded(errValidate)

	errMap := smapping.FillStruct(&comment, smapping.MapFields(&request))
	exception.PanicIfNeeded(errMap)

	result := services.commentRepo.Update(comment)
	res := response.NewCommentResponse(result)

	return res
}

func (services *CommentServicesImpl) DeleteById(id int64, userId int64) response.CommentResponse {
	result := services.commentRepo.Delete(id, userId)

	res := response.NewCommentResponse(result)
	return res
}
