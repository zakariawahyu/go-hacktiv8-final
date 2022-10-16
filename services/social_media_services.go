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

type SocialMediaServicesImpl struct {
	socialMediaRepo repository.SocialMediaRepository
}

func NewSocialMediaServices(mediaRepository repository.SocialMediaRepository) SocialMediaServices {
	return &SocialMediaServicesImpl{
		socialMediaRepo: mediaRepository,
	}
}

func (services *SocialMediaServicesImpl) CreateSocialMedia(request dto.SocialMediaRequest) response.SocialMediaResponse {
	var socialMedia entity.SocialMedia

	errValidate := validations.ValidateSocialMedia(request)
	exception.PanicIfNeeded(errValidate)

	errMap := smapping.FillStruct(&socialMedia, smapping.MapFields(&request))
	exception.PanicIfNeeded(errMap)

	result := services.socialMediaRepo.Create(socialMedia)
	res := response.NewSocialMediaResponse(result)

	return res
}

func (services *SocialMediaServicesImpl) AllSocialMedia(userID int64) []response.SocialMediaResponseAll {
	result := services.socialMediaRepo.GetAll(userID)
	res := response.NewSocialMediaResponseArray(result)

	return res
}
