package request

type CreateOrderRequest struct {
	LessonIds []int `json:"lesson_ids" form:"lesson_ids" binding:"required,min=1"`
}

type UpdateOrderRequest struct {
	ID       int64  `json:"id" form:"id"`
	Name     string `json:"name" form:"name" binding:"required,min=1"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}
