package _sub_lesson

import (
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
)

type SubLessonResponse struct {
	ID             int64  `json:"id"`
	SubLessonTitle string `json:"sub_lesson_title"`
}

func NewSubLessonResponse(lesson entity.SubLesson) SubLessonResponse {
	return SubLessonResponse{
		ID:             lesson.ID,
		SubLessonTitle: lesson.Title,
	}
}

func NewSubLessonArrayResponse(lessons []entity.SubLesson) []SubLessonResponse {
	lessonRes := []SubLessonResponse{}
	for _, v := range lessons {
		p := SubLessonResponse{
			ID:             v.ID,
			SubLessonTitle: v.Title,
		}
		lessonRes = append(lessonRes, p)
	}
	return lessonRes
}
