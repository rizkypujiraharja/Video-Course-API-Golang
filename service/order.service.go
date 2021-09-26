package service

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/mashingan/smapping"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/repo"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/request"
)

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func createInvoice() string {
	year, month, day := time.Now().Date()
	return "INV-" + strconv.Itoa(year) + strconv.Itoa(int(month)) + strconv.Itoa(day) + randomString(6)
}

type OrderService interface {
	All() (*[]entity.Order, error)
	CreateOrder(orderRequest request.CreateOrderRequest, userId string) (*entity.Order, error)
	UpdateOrder(updateOrderRequest request.UpdateOrderRequest) (*entity.Order, error)
	FindOneOrderByID(orderID string) (*entity.Order, error)
}

type orderService struct {
	orderRepo       repo.OrderRepository
	orderDetailRepo repo.OrderDetailRepository
	lessonRepo      repo.LessonRepository
}

func NewOrderService(orderRepo repo.OrderRepository, orderDetailRepo repo.OrderDetailRepository, lessonRepo repo.LessonRepository) OrderService {
	return &orderService{
		orderRepo:       orderRepo,
		orderDetailRepo: orderDetailRepo,
		lessonRepo:      lessonRepo,
	}
}

func (c *orderService) All() (*[]entity.Order, error) {
	orders, err := c.orderRepo.All()
	if err != nil {
		return nil, err
	}

	return &orders, nil
}

func (c *orderService) CreateOrder(orderRequest request.CreateOrderRequest, userId string) (*entity.Order, error) {
	order := entity.Order{}
	// Find Lessons
	lessons, err := c.lessonRepo.FindLessonByIDS(orderRequest.LessonIds)
	if err != nil {
		return nil, err
	}
	if len(lessons) == 0 {
		return nil, errors.New("lessons not found")
	}

	// Create Invoice Code
	order.InvoiceCode = createInvoice()

	// Assign Value To Order
	intUserId, _ := strconv.Atoi(userId)
	order.UserID = int64(intUserId)
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	order.Status = "unpaid"

	// Count total order
	for _, lesson := range lessons {
		order.TotalOrder += lesson.Price
	}

	ord, err := c.orderRepo.InsertOrder(order)
	if err != nil {
		return nil, err
	}

	fmt.Println("asd", ord)

	for _, lesson := range lessons {
		var ordDetail entity.OrderDetail
		ordDetail.LessonID = lesson.ID
		ordDetail.OrderID = ord.ID
		ordDetail.Price = int64(lesson.Price)

		c.orderDetailRepo.InsertOrderDetail(ordDetail)
		if err != nil {
			return nil, err
		}
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
