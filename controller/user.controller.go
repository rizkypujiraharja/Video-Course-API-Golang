package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/common/obj"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/common/response"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/request"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/resource"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/service"
)

type UserController interface {
	Profile(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(
	userService service.UserService,
	jwtService service.JWTService,
) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) Update(ctx *gin.Context) {
	var updateUserRequest request.UpdateUserRequest

	err := ctx.ShouldBind(&updateUserRequest)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id := c.jwtService.GetUserId(ctx)

	if id == "" {
		response := response.BuildErrorResponse("Error", "Failed to validate token", obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	_id, _ := strconv.ParseInt(id, 0, 64)
	updateUserRequest.ID = _id
	user, err := c.userService.UpdateUser(updateUserRequest)

	if err != nil {
		response := response.BuildErrorResponse("Error", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	res := resource.NewUserResponse(*user)
	response := response.BuildResponse(true, "OK", res)
	ctx.JSON(http.StatusOK, response)

}

func (c *userController) Profile(ctx *gin.Context) {
	userId := c.jwtService.GetUserId(ctx)
	user, _ := c.userService.FindUserByID(userId)

	res := resource.NewUserResponse(*user)
	response := response.BuildResponse(true, "OK", res)
	ctx.JSON(http.StatusOK, response)
}
