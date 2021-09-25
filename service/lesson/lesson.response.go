package _lesson

import (
	"fmt"

	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
	_category "github.com/rizkypujiraharja/Video-Course-API-Golang/service/category"
)

type LessonResponse struct {
	ID              int64                      `json:"id"`
	LessonTitle     string                     `json:"lesson_name"`
	Description     string                     `json:"description"`
	Price           uint64                     `json:"price"`
	ImageCoverUrl   string                     `json:"image_cover_url"`
	VideoPreviewUrl string                     `json:"video_preview_url"`
	Category        _category.CategoryResponse `json:"category"`
}

func NewLessonResponse(lesson entity.Lesson) LessonResponse {
	fmt.Println(lesson)
	return LessonResponse{
		ID:              lesson.ID,
		LessonTitle:     lesson.Title,
		Description:     lesson.Description,
		Price:           lesson.Price,
		ImageCoverUrl:   lesson.ImageCoverUrl,
		VideoPreviewUrl: lesson.VideoPreviewUrl,
		Category:        _category.NewCategoryResponse(lesson.Category),
	}
}

func NewLessonArrayResponse(lessons []entity.Lesson) []LessonResponse {
	lessonRes := []LessonResponse{}
	for _, v := range lessons {
		p := LessonResponse{
			ID:              v.ID,
			LessonTitle:     v.Title,
			Description:     v.Description,
			Price:           v.Price,
			ImageCoverUrl:   v.ImageCoverUrl,
			VideoPreviewUrl: v.VideoPreviewUrl,
			Category:        _category.NewCategoryResponse(v.Category),
		}
		lessonRes = append(lessonRes, p)
	}
	return lessonRes
}