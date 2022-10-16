package dto

type UpdatePhotoRequest struct {
	ID       int64  `json:"id"`
	Title    string `json:"title" form:"title"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photo_url" form:"photo_url"`
	UserID   int64  `json:"user_id"`
}

type PhotoRequest struct {
	Title    string `json:"title" form:"title"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photo_url" form:"photo_url"`
	UserID   int64  `json:"user_id"`
}
