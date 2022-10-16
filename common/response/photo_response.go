package response

import (
	"github.com/zakariawahyu/go-hacktiv8-final/entity"
	"time"
)

type PhotoResponse struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type PhotoResponseAll struct {
	ID        int64        `json:"id"`
	Title     string       `json:"title"`
	Caption   string       `json:"caption"`
	PhotoUrl  string       `json:"photo_url"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	UserID    int64        `json:"user_id"`
	User      UserResponse `json:"user"`
}

func NewPhotoResponse(photo entity.Photo) PhotoResponse {
	return PhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
	}
}

func NewPhotoResponseArray(task []entity.Photo) []PhotoResponseAll {
	photoRes := []PhotoResponseAll{}
	for _, value := range task {
		photo := PhotoResponseAll{
			ID:        value.ID,
			Title:     value.Title,
			Caption:   value.Caption,
			PhotoUrl:  value.PhotoUrl,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
			UserID:    value.UserID,
			User:      NewUserResponse(value.User),
		}
		photoRes = append(photoRes, photo)
	}
	return photoRes
}
