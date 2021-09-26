package repo

import (
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
	"gorm.io/gorm"
)

type OrderDetailRepository interface {
	InsertOrderDetail(orderDetail entity.OrderDetail) (entity.OrderDetail, error)
}

type orderDetailRepo struct {
	connection *gorm.DB
}

func NewOrderDetailRepo(connection *gorm.DB) OrderDetailRepository {
	return &orderDetailRepo{
		connection: connection,
	}
}

func (c *orderDetailRepo) InsertOrderDetail(orderDetail entity.OrderDetail) (entity.OrderDetail, error) {
	c.connection.Save(&orderDetail)
	c.connection.Find(&orderDetail)
	return orderDetail, nil
}
