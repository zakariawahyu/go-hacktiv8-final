package dto

type UpdateSocialMediaRequest struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         int64  `json:"user_id"`
}

type SocialMediaRequest struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         int64  `json:"user_id"`
}
