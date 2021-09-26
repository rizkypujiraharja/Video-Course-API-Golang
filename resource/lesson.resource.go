package resource

import (
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
)

type LessonResponse struct {
	ID              int64               `json:"id"`
	LessonTitle     string              `json:"lesson_name"`
	Description     string              `json:"description"`
	Price           uint64              `json:"price"`
	ImageCoverUrl   string              `json:"image_cover_url"`
	VideoPreviewUrl string              `json:"video_preview_url"`
	Category        CategoryResponse    `json:"category"`
	SubLessons      []SubLessonResponse `json:"sub_lessons,omitempty"`
}

func NewLessonResponse(lesson entity.Lesson) LessonResponse {
	return LessonResponse{
		ID:              lesson.ID,
		LessonTitle:     lesson.Title,
		Description:     lesson.Description,
		Price:           lesson.Price,
		ImageCoverUrl:   lesson.ImageCoverUrl,
		VideoPreviewUrl: lesson.VideoPreviewUrl,
		Category:        NewCategoryResponse(lesson.Category),
		SubLessons:      NewSubLessonArrayResponse(lesson.SubLessons),
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
			Category:        NewCategoryResponse(v.Category),
		}
		lessonRes = append(lessonRes, p)
	}
	return lessonRes
}

func NewLessonFromOrderedLessonArrayResponse(orderedLessons []entity.OrderedLesson) []LessonResponse {
	lessonRes := []LessonResponse{}
	for _, v := range orderedLessons {
		p := LessonResponse{
			ID:              v.Lesson.ID,
			LessonTitle:     v.Lesson.Title,
			Description:     v.Lesson.Description,
			Price:           v.Lesson.Price,
			ImageCoverUrl:   v.Lesson.ImageCoverUrl,
			VideoPreviewUrl: v.Lesson.VideoPreviewUrl,
			Category:        NewCategoryResponse(v.Lesson.Category),
		}
		lessonRes = append(lessonRes, p)
	}
	return lessonRes
}
