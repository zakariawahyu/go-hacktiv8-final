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

type PhotoServicesImpl struct {
	photoRepo repository.PhotoRepository
}

func NewPhotoServices(photoRepository repository.PhotoRepository) PhotoServices {
	return &PhotoServicesImpl{
		photoRepo: photoRepository,
	}
}

func (services *PhotoServicesImpl) CreatePhoto(request dto.PhotoRequest) response.PhotoResponse {
	var photo entity.Photo

	errValidate := validations.ValidatePhoto(request)
	exception.PanicIfNeeded(errValidate)

	errMap := smapping.FillStruct(&photo, smapping.MapFields(&request))
	exception.PanicIfNeeded(errMap)

	result := services.photoRepo.Create(photo)
	res := response.NewPhotoResponse(result)

	return res
}

func (services *PhotoServicesImpl) AllPhoto(userID int64) []response.PhotoResponseAll {
	result := services.photoRepo.GetAll(userID)
	res := response.NewPhotoResponseArray(result)

	return res
}

func (services *PhotoServicesImpl) UpdatePhoto(request dto.UpdatePhotoRequest) response.PhotoResponse {
	var photo entity.Photo

	errValidate := validations.ValidateUpdatePhoto(request)
	exception.PanicIfNeeded(errValidate)

	errMap := smapping.FillStruct(&photo, smapping.MapFields(&request))
	exception.PanicIfNeeded(errMap)

	result := services.photoRepo.Update(photo)
	res := response.NewPhotoResponse(result)

	return res
}
