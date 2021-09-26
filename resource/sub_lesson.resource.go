package resource

import (
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
)

type SubLessonResponse struct {
	ID             int64           `json:"id"`
	SubLessonTitle string          `json:"sub_lesson_title"`
	Videos         []VideoResponse `json:"videos"`
}

func NewSubLessonResponse(lesson entity.SubLesson) SubLessonResponse {
	return SubLessonResponse{
		ID:             lesson.ID,
		SubLessonTitle: lesson.Title,
		Videos:         NewVideoArrayResponse(lesson.Videos),
	}
}

func NewSubLessonArrayResponse(lessons []entity.SubLesson) []SubLessonResponse {
	lessonRes := []SubLessonResponse{}
	for _, v := range lessons {
		p := SubLessonResponse{
			ID:             v.ID,
			SubLessonTitle: v.Title,
			Videos:         NewVideoArrayResponse(v.Videos),
		}
		lessonRes = append(lessonRes, p)
	}
	return lessonRes
}
