package request

type LoginRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}

type RegisterRequest struct {
	Name     string `json:"name" form:"name" binding:"required,min=1"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}

type UpdateUserLoginRequest struct {
	Name  string `json:"name" form:"name" binding:"required,min=1"`
	Email string `json:"email" form:"email" binding:"required,email"`
}
