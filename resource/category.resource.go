package resource

import (
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
)

type CategoryResponse struct {
	ID           int64  `json:"id"`
	CategoryName string `json:"category_name"`
	Slug         string `json:"slug"`
}

func NewCategoryResponse(category entity.Category) CategoryResponse {
	return CategoryResponse{
		ID:           category.ID,
		CategoryName: category.Name,
		Slug:         category.Slug,
	}
}

func NewCategoryArrayResponse(categories []entity.Category) []CategoryResponse {
	lessonRes := []CategoryResponse{}
	for _, v := range categories {
		p := CategoryResponse{
			ID:           v.ID,
			CategoryName: v.Name,
			Slug:         v.Slug,
		}
		lessonRes = append(lessonRes, p)
	}
	return lessonRes
}
