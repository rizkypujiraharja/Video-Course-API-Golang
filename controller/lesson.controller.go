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

type LessonController interface {
	All(ctx *gin.Context)
	MyLesson(ctx *gin.Context)
	CreateLesson(ctx *gin.Context)
	UpdateLesson(ctx *gin.Context)
	DeleteLesson(ctx *gin.Context)
	FindOneLessonByID(ctx *gin.Context)
	FindOneLessonByIDPublic(ctx *gin.Context)
}

type lessonController struct {
	lessonService service.LessonService
	jwtService    service.JWTService
}

func NewLessonController(lessonService service.LessonService, jwtService service.JWTService) LessonController {
	return &lessonController{
		lessonService: lessonService,
		jwtService:    jwtService,
	}
}

func (c *lessonController) All(ctx *gin.Context) {
	lessons, err := c.lessonService.All()
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	res := resource.NewLessonArrayResponse(*lessons)
	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *lessonController) MyLesson(ctx *gin.Context) {
	userId := c.jwtService.GetUserId(ctx)

	lessons, err := c.lessonService.MyLesson(userId)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	res := resource.NewLessonFromOrderedLessonArrayResponse(*lessons)
	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *lessonController) CreateLesson(ctx *gin.Context) {
	var createLessonReq request.CreateLessonRequest
	err := ctx.ShouldBind(&createLessonReq)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	lesson, err := c.lessonService.CreateLesson(createLessonReq)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	res := resource.NewLessonResponse(*lesson)
	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusCreated, response)

}

func (c *lessonController) FindOneLessonByID(ctx *gin.Context) {
	id := ctx.Param("id")

	lesson, err := c.lessonService.FindOneLessonByID(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res := resource.NewLessonResponse(*lesson)
	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *lessonController) FindOneLessonByIDPublic(ctx *gin.Context) {
	id := ctx.Param("id")

	lesson, err := c.lessonService.FindOneLessonByIDPublic(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res := resource.NewLessonResponse(*lesson)
	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *lessonController) DeleteLesson(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.lessonService.DeleteLesson(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := response.BuildResponse(true, "OK!", obj.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}

func (c *lessonController) UpdateLesson(ctx *gin.Context) {
	updateLessonRequest := request.UpdateLessonRequest{}
	err := ctx.ShouldBind(&updateLessonRequest)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id, _ := strconv.ParseInt(ctx.Param("id"), 0, 64)
	updateLessonRequest.ID = id
	lesson, err := c.lessonService.UpdateLesson(updateLessonRequest)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	res := resource.NewLessonResponse(*lesson)
	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)

}
