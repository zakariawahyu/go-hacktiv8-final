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

	err := repository.db.Where("user_id = ?", userID).Preload("User").Preload("Photo").Find(&comment).Error
	exception.PanicIfNeeded(err)

	return comment
}

func (repository *CommentRepositoryImpl) Update(comment entity.Comment) entity.Comment {
	err := repository.db.Where("id = ? and user_id = ?", comment.ID, comment.UserID).Updates(&comment).First(&comment).Error
	exception.PanicIfNeeded(err)

	return comment
}

func (repository *CommentRepositoryImpl) Delete(id int64, userID int64) entity.Comment {
	var comment entity.Comment
	err := repository.db.Where("id = ? AND user_id = ?", id, userID).Delete(&comment)
	if err.RowsAffected == 0 {
		exception.PanicIfNeeded("Record not found")
	}

	return comment
}
