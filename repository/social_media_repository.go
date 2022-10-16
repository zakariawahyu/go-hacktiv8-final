package repository

import (
	"github.com/zakariawahyu/go-hacktiv8-final/entity"
	"github.com/zakariawahyu/go-hacktiv8-final/exception"
	"gorm.io/gorm"
)

type SocialMediaRepositoryImpl struct {
	db *gorm.DB
}

func NewSocialMediaRepository(database *gorm.DB) SocialMediaRepository {
	return &SocialMediaRepositoryImpl{
		db: database,
	}
}

func (repository *SocialMediaRepositoryImpl) Create(socialMedia entity.SocialMedia) entity.SocialMedia {
	err := repository.db.Create(&socialMedia).Error
	exception.PanicIfNeeded(err)

	return socialMedia
}

func (repository *SocialMediaRepositoryImpl) GetAll(userID int64) []entity.SocialMedia {
	var socialMedia []entity.SocialMedia

	err := repository.db.Where("user_id = ?", userID).Preload("User").Find(&socialMedia).Error
	exception.PanicIfNeeded(err)

	return socialMedia
}

func (repository *SocialMediaRepositoryImpl) Update(socialMedia entity.SocialMedia) entity.SocialMedia {
	err := repository.db.Where("id = ? and user_id = ?", socialMedia.ID, socialMedia.UserID).Updates(&socialMedia).First(&socialMedia).Error
	exception.PanicIfNeeded(err)

	return socialMedia
}
