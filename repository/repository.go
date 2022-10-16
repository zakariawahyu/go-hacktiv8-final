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
