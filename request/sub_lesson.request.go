package request

type CreateSubLessonRequest struct {
	Title    string `json:"title" form:"title" binding:"required,min=1"`
	LessonID int64  `json:"lesson_id" binding:"required"`
}

type UpdateSubLessonRequest struct {
	ID       int64  `json:"id" form:"id"`
	Title    string `json:"title" form:"title" binding:"required,min=1"`
	LessonID int64  `json:"lesson_id" binding:"required"`
}
