package response

import (
	"github.com/zakariawahyu/go-hacktiv8-final/entity"
	"time"
)

type SocialMediaResponse struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         int64     `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type SocialMediaResponseAll struct {
	ID             int64        `json:"id"`
	Name           string       `json:"name"`
	SocialMediaUrl string       `json:"social_media_url"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	UserID         int64        `json:"user_id"`
	User           UserResponse `json:"user"`
}

func NewSocialMediaResponse(socialMedia entity.SocialMedia) SocialMediaResponse {
	return SocialMediaResponse{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserID:         socialMedia.UserID,
		CreatedAt:      socialMedia.CreatedAt,
	}
}

func NewSocialMediaResponseArray(socialMedia []entity.SocialMedia) []SocialMediaResponseAll {
	socialMediaRes := []SocialMediaResponseAll{}
	for _, value := range socialMedia {
		socialMedia := SocialMediaResponseAll{
			ID:             value.ID,
			Name:           value.Name,
			SocialMediaUrl: value.SocialMediaUrl,
			UserID:         value.UserID,
			User:           NewUserResponse(value.User),
			CreatedAt:      value.CreatedAt,
			UpdatedAt:      value.UpdatedAt,
		}
		socialMediaRes = append(socialMediaRes, socialMedia)
	}
	return socialMediaRes
}
