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

type VideoController interface {
	CreateVideo(ctx *gin.Context)
	UpdateVideo(ctx *gin.Context)
	DeleteVideo(ctx *gin.Context)
	FindOneVideoByID(ctx *gin.Context)
}

type videoController struct {
	videoService service.VideoService
	jwtService   service.JWTService
}

func NewVideoController(videoService service.VideoService, jwtService service.JWTService) VideoController {
	return &videoController{
		videoService: videoService,
		jwtService:   jwtService,
	}
}

func (c *videoController) CreateVideo(ctx *gin.Context) {
	var createVideoReq request.CreateVideoRequest
	err := ctx.ShouldBind(&createVideoReq)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	res, err := c.videoService.CreateVideo(createVideoReq)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusCreated, response)

}

func (c *videoController) FindOneVideoByID(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.videoService.FindOneVideoByID(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *videoController) DeleteVideo(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.videoService.DeleteVideo(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := response.BuildResponse(true, "OK!", obj.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}

func (c *videoController) UpdateVideo(ctx *gin.Context) {
	updateVideoRequest := request.UpdateVideoRequest{}
	err := ctx.ShouldBind(&updateVideoRequest)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id, _ := strconv.ParseInt(ctx.Param("id"), 0, 64)
	updateVideoRequest.ID = id
	video, err := c.videoService.UpdateVideo(updateVideoRequest)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", video)
	ctx.JSON(http.StatusOK, response)

}
