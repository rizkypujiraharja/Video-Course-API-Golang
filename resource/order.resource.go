package resource

import (
	"time"

	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
	// _order_detail "github.com/rizkypujiraharja/Video-Course-API-Golang/service/order_detail"
)

type OrderResponse struct {
	ID          int64        `json:"id"`
	InvoiceCode string       `json:"invoice_code"`
	Status      string       `json:"status"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	User        UserResponse `json:"user"`
	// OrderDetails _order_detail `json:"order_details,omitempty"`
}

func NewOrderResponse(order entity.Order) OrderResponse {
	return OrderResponse{
		ID:          order.ID,
		InvoiceCode: order.InvoiceCode,
		Status:      order.Status,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
		User:        NewUserResponse(order.User),
		// OrderDetails: order.OrderDetails
	}
}

func NewOrderArrayResponse(orders []entity.Order) []OrderResponse {
	detailsRes := []OrderResponse{}
	for _, v := range orders {
		p := OrderResponse{
			ID:          v.ID,
			InvoiceCode: v.InvoiceCode,
			Status:      v.Status,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			User:        NewUserResponse(v.User),
		}
		detailsRes = append(detailsRes, p)
	}
	return detailsRes
}
