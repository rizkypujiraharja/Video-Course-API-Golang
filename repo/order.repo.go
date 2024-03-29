package repo

import (
	"fmt"

	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
	"gorm.io/gorm"
)

type OrderRepository interface {
	All() ([]entity.Order, error)
	InsertOrder(order entity.Order) (entity.Order, error)
	UpdateOrder(order entity.Order) (entity.Order, error)
	FindOneOrderByID(ID string) (entity.Order, error)
	FindOrderByUserID(ID string) ([]entity.Order, error)
}

type orderRepo struct {
	connection *gorm.DB
}

func NewOrderRepo(connection *gorm.DB) OrderRepository {
	return &orderRepo{
		connection: connection,
	}
}

func (c *orderRepo) All() ([]entity.Order, error) {
	orders := []entity.Order{}
	c.connection.Preload("User").Find(&orders)
	return orders, nil
}

func (c *orderRepo) FindOrderByUserID(userID string) ([]entity.Order, error) {
	orders := []entity.Order{}
	c.connection.Preload("User").Where("user_id = ?", userID).Find(&orders)
	return orders, nil
}

func (c *orderRepo) InsertOrder(order entity.Order) (entity.Order, error) {
	c.connection.Save(&order)
	c.connection.Preload("User").Find(&order)
	return order, nil
}

func (c *orderRepo) UpdateOrder(order entity.Order) (entity.Order, error) {
	c.connection.Save(&order)
	c.connection.Preload("User").Find(&order)
	return order, nil
}

func (c *orderRepo) FindOneOrderByID(orderID string) (entity.Order, error) {
	var order entity.Order
	res := c.connection.Preload("User").Preload("OrderDetails.Lesson.Category").Where("id = ?", orderID).Take(&order)
	fmt.Println(res, order)
	if res.Error != nil {
		return order, res.Error
	}
	return order, nil
}
