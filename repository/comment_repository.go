package repository

import (
	"github.com/zakariawahyu/go-hacktiv8-final/entity"
	"github.com/zakariawahyu/go-hacktiv8-final/exception"
	"gorm.io/gorm"
)

type CommentRepositoryImpl struct {
	db *gorm.DB
}

func NewCommentRepository(database *gorm.DB) CommentRepository {
	return &CommentRepositoryImpl{
		db: database,
	}
}

func (repository *CommentRepositoryImpl) Create(comment entity.Comment) entity.Comment {
	err := repository.db.Create(&comment).Error
	exception.PanicIfNeeded(err)

	return comment
}

func (repository *CommentRepositoryImpl) GetAll(userID int64) []entity.Comment {
	var comment []entity.Comment

	err := repository.db.Where("user_id = ?", userID).Find(&comment).Error
	exception.PanicIfNeeded(err)

	repository.db.Preload("User").Preload("Photo").Find(&comment)
	return comment
}
