package _sub_lesson

import (
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
	_video "github.com/rizkypujiraharja/Video-Course-API-Golang/service/video"
)

type SubLessonResponse struct {
	ID             int64                  `json:"id"`
	SubLessonTitle string                 `json:"sub_lesson_title"`
	Videos         []_video.VideoResponse `json:"videos"`
}

func NewSubLessonResponse(lesson entity.SubLesson) SubLessonResponse {
	return SubLessonResponse{
		ID:             lesson.ID,
		SubLessonTitle: lesson.Title,
		Videos:         _video.NewVideoArrayResponse(lesson.Videos),
	}
}

func NewSubLessonArrayResponse(lessons []entity.SubLesson) []SubLessonResponse {
	lessonRes := []SubLessonResponse{}
	for _, v := range lessons {
		p := SubLessonResponse{
			ID:             v.ID,
			SubLessonTitle: v.Title,
			Videos:         _video.NewVideoArrayResponse(v.Videos),
		}
		lessonRes = append(lessonRes, p)
	}
	return lessonRes
}
