package repository

import (
	"github.com/zakariawahyu/go-hacktiv8-final/entity"
	"github.com/zakariawahyu/go-hacktiv8-final/exception"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: database,
	}
}

func (repository *UserRepositoryImpl) Create(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))

	err := repository.db.Create(&user).Error
	exception.PanicIfNeeded(err)

	return user
}

func (repository *UserRepositoryImpl) FindByEmail(email string) entity.User {
	var user entity.User

	err := repository.db.Where("email = ?", email).Take(&user).Error
	exception.PanicIfNeeded(err)

	return user
}

func (repository *UserRepositoryImpl) Update(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	}

	err := repository.db.Where("id = ?", user.ID).Updates(&user).First(&user, "id = ?", user.ID).Error
	exception.PanicIfNeeded(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(userID int64) entity.User {
	var user entity.User
	err := repository.db.Where("id = ?", userID).Delete(&user)
	if err.RowsAffected == 0 {
		exception.PanicIfNeeded("Record not found")
	}

	return user
}

func hashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	exception.PanicIfNeeded(err)
	return string(hash)
}
