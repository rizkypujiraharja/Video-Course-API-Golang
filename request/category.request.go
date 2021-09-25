package request

type CreateCategoryRequest struct {
	Name string `json:"name" form:"name" binding:"required,min=1"`
}

type UpdateCategoryRequest struct {
	ID   int64  `json:"id" form:"id"`
	Name string `json:"name" form:"name" binding:"required,min=1"`
}
