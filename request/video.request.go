package request

type CreateVideoRequest struct {
	Title       string `json:"title" form:"title" binding:"required,min=1"`
	Description string `json:"description"`
	VideoUrl    string `json:"video_url" binding:"required"`
	SubLessonID int64  `json:"sub_lesson_id" binding:"required"`
}

type UpdateVideoRequest struct {
	ID          int64  `json:"id" form:"id"`
	Title       string `json:"title" form:"title" binding:"required,min=1"`
	Description string `json:"description"`
	VideoUrl    string `json:"video_url" binding:"required"`
	SubLessonID int64  `json:"sub_lesson_id" binding:"required"`
}
