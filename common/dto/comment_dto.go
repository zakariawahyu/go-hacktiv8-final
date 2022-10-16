package dto

type UpdateCommentRequest struct {
	ID      int64  `json:"id"`
	Message string `json:"message" form:"message"`
	PhotoID int64  `json:"photo_id" form:"photo_id"`
	UserID  int64  `json:"user_id" form:"user_id"`
}

type CommentRequest struct {
	Message string `json:"message" form:"message"`
	PhotoID int64  `json:"photo_id" form:"photo_id"`
	UserID  int64  `json:"user_id" form:"user_id"`
}
