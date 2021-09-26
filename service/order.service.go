package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/repo"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/request"
)

type OrderService interface {
	All() (*[]entity.Order, error)
	CreateOrder(orderRequest request.CreateOrderRequest) (*entity.Order, error)
	UpdateOrder(updateOrderRequest request.UpdateOrderRequest) (*entity.Order, error)
	FindOneOrderByID(orderID string) (*entity.Order, error)
}

type orderService struct {
	orderRepo repo.OrderRepository
}

func NewOrderService(orderRepo repo.OrderRepository) OrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (c *orderService) All() (*[]entity.Order, error) {
	orders, err := c.orderRepo.All()
	if err != nil {
		return nil, err
	}

	return &orders, nil
}

func (c *orderService) CreateOrder(orderRequest request.CreateOrderRequest) (*entity.Order, error) {
	order := entity.Order{}
	err := smapping.FillStruct(&order, smapping.MapFields(&orderRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	ord, err := c.orderRepo.InsertOrder(order)
	if err != nil {
		return nil, err
	}

	return &ord, nil
}

func (c *orderService) FindOneOrderByID(orderID string) (*entity.Order, error) {
	order, err := c.orderRepo.FindOneOrderByID(orderID)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *orderService) UpdateOrder(updateOrderRequest request.UpdateOrderRequest) (*entity.Order, error) {
	order, err := c.orderRepo.FindOneOrderByID(fmt.Sprintf("%d", updateOrderRequest.ID))
	if err != nil {
		return nil, err
	}

	order = entity.Order{}
	err = smapping.FillStruct(&order, smapping.MapFields(&updateOrderRequest))

	if err != nil {
		return nil, err
	}

	order, err = c.orderRepo.UpdateOrder(order)

	if err != nil {
		return nil, err
	}

	return &order, nil
}
