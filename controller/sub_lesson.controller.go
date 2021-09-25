package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/common/obj"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/common/response"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/request"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/service"
)

type SubLessonController interface {
	CreateSubLesson(ctx *gin.Context)
	UpdateSubLesson(ctx *gin.Context)
	DeleteSubLesson(ctx *gin.Context)
	FindOneSubLessonByID(ctx *gin.Context)
}

type subLessonController struct {
	subLessonService service.SubLessonService
	jwtService       service.JWTService
}

func NewSubLessonController(subLessonService service.SubLessonService, jwtService service.JWTService) SubLessonController {
	return &subLessonController{
		subLessonService: subLessonService,
		jwtService:       jwtService,
	}
}

func (c *subLessonController) CreateSubLesson(ctx *gin.Context) {
	var createSubLessonReq request.CreateSubLessonRequest
	err := ctx.ShouldBind(&createSubLessonReq)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	res, err := c.subLessonService.CreateSubLesson(createSubLessonReq)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusCreated, response)

}

func (c *subLessonController) FindOneSubLessonByID(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.subLessonService.FindOneSubLessonByID(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *subLessonController) DeleteSubLesson(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.subLessonService.DeleteSubLesson(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := response.BuildResponse(true, "OK!", obj.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}

func (c *subLessonController) UpdateSubLesson(ctx *gin.Context) {
	updateSubLessonRequest := request.UpdateSubLessonRequest{}
	err := ctx.ShouldBind(&updateSubLessonRequest)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id, _ := strconv.ParseInt(ctx.Param("id"), 0, 64)
	updateSubLessonRequest.ID = id
	subLesson, err := c.subLessonService.UpdateSubLesson(updateSubLessonRequest)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", subLesson)
	ctx.JSON(http.StatusOK, response)

}
