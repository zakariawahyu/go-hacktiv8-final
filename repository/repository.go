package repository

import "github.com/zakariawahyu/go-hacktiv8-final/entity"

type UserRepository interface {
	Create(user entity.User) entity.User
	FindByEmail(email string) entity.User
}

type PhotoRepository interface {
	Create(photo entity.Photo) entity.Photo
	GetAll(userID int64) []entity.Photo
}

type CommentRepository interface {
	Create(comment entity.Comment) entity.Comment
	GetAll(userID int64) []entity.Comment
}

type SocialMediaRepository interface {
	Create(socialMedia entity.SocialMedia) entity.SocialMedia
	GetAll(userID int64) []entity.SocialMedia
}
