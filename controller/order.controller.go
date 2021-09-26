package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/common/obj"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/common/response"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/request"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/resource"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/service"
)

type OrderController interface {
	All(ctx *gin.Context)
	MyOrder(ctx *gin.Context)
	CreateOrder(ctx *gin.Context)
	FindOneOrderByID(ctx *gin.Context)
	UpdatePaidOrder(ctx *gin.Context)
	UpdateUnpaidOrder(ctx *gin.Context)
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

	res := resource.NewOrderArrayResponse(*orders)
	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *orderController) MyOrder(ctx *gin.Context) {
	userId := c.jwtService.GetUserId(ctx)

	orders, err := c.orderService.FindOrderByUserID(userId)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res := resource.NewOrderArrayResponse(*orders)
	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *orderController) CreateOrder(ctx *gin.Context) {
	var createOrderReq request.CreateOrderRequest
	err := ctx.ShouldBind(&createOrderReq)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userId := c.jwtService.GetUserId(ctx)

	order, err := c.orderService.CreateOrder(createOrderReq, userId)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}
	res := resource.NewOrderResponse(*order)
	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusCreated, response)
}

func (c *orderController) FindOneOrderByID(ctx *gin.Context) {
	id := ctx.Param("id")

	order, err := c.orderService.FindOneOrderByID(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res := resource.NewOrderResponse(*order)

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *orderController) UpdatePaidOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	order, err := c.orderService.UpdatePaidOrder(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	res := resource.NewOrderResponse(*order)
	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *orderController) UpdateUnpaidOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	order, err := c.orderService.UpdateUnpaidOrder(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	res := resource.NewOrderResponse(*order)
	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}
