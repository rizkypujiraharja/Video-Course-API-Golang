package request

type CreateLessonRequest struct {
	Title           string `json:"title" form:"title" binding:"required,min=1"`
	Description     string `json:"description"`
	Price           uint64 `json:"price" binding:"required"`
	ImageCoverUrl   string `json:"image_cover_url" binding:"required,url"`
	VideoPreviewUrl string `json:"video_preview_url" binding:"required,url"`
	CategoryID      int64  `json:"category_id" binding:"required"`
}

type UpdateLessonRequest struct {
	ID              int64  `json:"id" form:"id"`
	Title           string `json:"title" form:"title" binding:"required,min=1"`
	Description     string `json:"description"`
	Price           uint64 `json:"price" binding:"required"`
	ImageCoverUrl   string `json:"image_cover_url" binding:"required,url"`
	VideoPreviewUrl string `json:"video_preview_url" binding:"required,url"`
	CategoryID      int64  `json:"category_id" binding:"required"`
}
