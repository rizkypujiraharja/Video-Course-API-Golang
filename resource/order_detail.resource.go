package resource

import (
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
)

type OrderDetailResponse struct {
	ID     int64          `json:"id"`
	Lesson LessonResponse `json:"lesson"`
}

func NewOrderDetailResponse(orderDetail entity.OrderDetail) OrderDetailResponse {
	return OrderDetailResponse{
		ID:     orderDetail.ID,
		Lesson: NewLessonResponse(orderDetail.Lesson),
	}
}

func NewOrderDetailArrayResponse(orderDetails []entity.OrderDetail) []OrderDetailResponse {
	detailsRes := []OrderDetailResponse{}
	for _, v := range orderDetails {
		p := OrderDetailResponse{
			ID:     v.ID,
			Lesson: NewLessonResponse(v.Lesson),
		}
		detailsRes = append(detailsRes, p)
	}
	return detailsRes
}
