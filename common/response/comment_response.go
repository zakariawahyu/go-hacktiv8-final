package response

import (
	"github.com/zakariawahyu/go-hacktiv8-final/entity"
	"time"
)

type CommentResponse struct {
	ID        int64     `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int64     `json:"photo_id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentResponseAll struct {
	ID        int64         `json:"id"`
	Message   string        `json:"message"`
	PhotoID   int64         `json:"photo_id"`
	Photo     PhotoResponse `json:"photo"`
	UserID    int64         `json:"user_id"`
	User      UserResponse  `json:"user"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

func NewCommentResponse(comment entity.Comment) CommentResponse {
	return CommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt,
	}
}

func NewCommentResponseArray(comment []entity.Comment) []CommentResponseAll {
	commentRes := []CommentResponseAll{}
	for _, value := range comment {
		comment := CommentResponseAll{
			ID:        value.ID,
			Message:   value.Message,
			PhotoID:   value.PhotoID,
			Photo:     NewPhotoResponse(value.Photo),
			UserID:    value.UserID,
			User:      NewUserResponse(value.User),
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		commentRes = append(commentRes, comment)
	}
	return commentRes
}
