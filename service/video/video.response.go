package _video

import (
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
)

type VideoResponse struct {
	ID          int64  `json:"id"`
	LessonTitle string `json:"lesson_name"`
	Description string `json:"description"`
	VideoUrl    string `json:"video_url"`
}

func NewVideoResponse(video entity.Video) VideoResponse {
	return VideoResponse{
		ID:          video.ID,
		LessonTitle: video.Title,
		Description: video.Description,
		VideoUrl:    video.VideoUrl,
	}
}

func NewVideoArrayResponse(videos []entity.Video) []VideoResponse {
	videosRes := []VideoResponse{}
	for _, v := range videos {
		p := VideoResponse{
			ID:          v.ID,
			LessonTitle: v.Title,
			Description: v.Description,
			VideoUrl:    v.VideoUrl,
		}
		videosRes = append(videosRes, p)
	}
	return videosRes
}
