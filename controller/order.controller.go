package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/common/obj"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/common/response"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/request"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/service"
)

type OrderController interface {
	All(ctx *gin.Context)
	CreateOrder(ctx *gin.Context)
	UpdateOrder(ctx *gin.Context)
	FindOneOrderByID(ctx *gin.Context)
}

type orderController struct {
	orderService service.OrderService
	jwtService   service.JWTService
}

func NewOrderController(orderService service.OrderService, jwtService service.JWTService) OrderController {
	return &orderController{
		orderService: orderService,
		jwtService:   jwtService,
	}
}

func (c *orderController) All(ctx *gin.Context) {
	orders, err := c.orderService.All()
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", orders)
	ctx.JSON(http.StatusOK, response)
}

func (c *orderController) CreateOrder(ctx *gin.Context) {
	var createOrderReq request.CreateOrderRequest
	err := ctx.ShouldBind(&createOrderReq)
	fmt.Println(createOrderReq)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	res, err := c.orderService.CreateOrder(createOrderReq)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusCreated, response)

}

func (c *orderController) FindOneOrderByID(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.orderService.FindOneOrderByID(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *orderController) UpdateOrder(ctx *gin.Context) {
	updateOrderRequest := request.UpdateOrderRequest{}
	err := ctx.ShouldBind(&updateOrderRequest)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id, _ := strconv.ParseInt(ctx.Param("id"), 0, 64)
	updateOrderRequest.ID = id
	order, err := c.orderService.UpdateOrder(updateOrderRequest)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", order)
	ctx.JSON(http.StatusOK, response)

}
