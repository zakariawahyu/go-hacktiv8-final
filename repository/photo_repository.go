package repository

import (
	"github.com/zakariawahyu/go-hacktiv8-final/entity"
	"github.com/zakariawahyu/go-hacktiv8-final/exception"
	"gorm.io/gorm"
)

type PhotoRepositoryImpl struct {
	db *gorm.DB
}

func NewPhotoRepository(database *gorm.DB) PhotoRepository {
	return &PhotoRepositoryImpl{
		db: database,
	}
}

func (repository *PhotoRepositoryImpl) Create(photo entity.Photo) entity.Photo {
	err := repository.db.Create(&photo).Error
	exception.PanicIfNeeded(err)

	return photo
}

func (repository *PhotoRepositoryImpl) GetAll(userID int64) []entity.Photo {
	var photo []entity.Photo

	err := repository.db.Where("user_id = ?", userID).Preload("User").Find(&photo).Error
	exception.PanicIfNeeded(err)

	return photo
}

func (repository *PhotoRepositoryImpl) Update(photo entity.Photo) entity.Photo {
	err := repository.db.Where("id = ? and user_id = ?", photo.ID, photo.UserID).Updates(&photo).First(&photo).Error
	exception.PanicIfNeeded(err)

	return photo
}
