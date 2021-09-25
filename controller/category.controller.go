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

type CategoryController interface {
	All(ctx *gin.Context)
	CreateCategory(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
	FindOneCategoryByID(ctx *gin.Context)
}

type categoryController struct {
	categoryService service.CategoryService
	jwtService      service.JWTService
}

func NewCategoryController(categoryService service.CategoryService, jwtService service.JWTService) CategoryController {
	return &categoryController{
		categoryService: categoryService,
		jwtService:      jwtService,
	}
}

func (c *categoryController) All(ctx *gin.Context) {
	categories, err := c.categoryService.All()
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", categories)
	ctx.JSON(http.StatusOK, response)
}

func (c *categoryController) CreateCategory(ctx *gin.Context) {
	var createCategoryReq request.CreateCategoryRequest
	err := ctx.ShouldBind(&createCategoryReq)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res, err := c.categoryService.CreateCategory(createCategoryReq)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusCreated, response)

}

func (c *categoryController) FindOneCategoryByID(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.categoryService.FindOneCategoryByID(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *categoryController) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.categoryService.DeleteCategory(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := response.BuildResponse(true, "OK!", obj.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}

func (c *categoryController) UpdateCategory(ctx *gin.Context) {
	updateCategoryRequest := request.UpdateCategoryRequest{}
	err := ctx.ShouldBind(&updateCategoryRequest)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id, _ := strconv.ParseInt(ctx.Param("id"), 0, 64)
	updateCategoryRequest.ID = id
	category, err := c.categoryService.UpdateCategory(updateCategoryRequest)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", category)
	ctx.JSON(http.StatusOK, response)

}
