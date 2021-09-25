package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/repo"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/request"

	_video "github.com/rizkypujiraharja/Video-Course-API-Golang/service/video"
)

type VideoService interface {
	CreateVideo(videoRequest request.CreateVideoRequest) (*_video.VideoResponse, error)
	UpdateVideo(updateVideoRequest request.UpdateVideoRequest) (*_video.VideoResponse, error)
	FindOneVideoByID(videoID string) (*_video.VideoResponse, error)
	DeleteVideo(videoID string) error
}

type videoService struct {
	videoRepo repo.VideoRepository
}

func NewVideoService(videoRepo repo.VideoRepository) VideoService {
	return &videoService{
		videoRepo: videoRepo,
	}
}

func (c *videoService) CreateVideo(videoRequest request.CreateVideoRequest) (*_video.VideoResponse, error) {
	video := entity.Video{}
	err := smapping.FillStruct(&video, smapping.MapFields(&videoRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	p, err := c.videoRepo.InsertVideo(video)
	if err != nil {
		return nil, err
	}

	res := _video.NewVideoResponse(p)
	return &res, nil
}

func (c *videoService) FindOneVideoByID(videoID string) (*_video.VideoResponse, error) {
	video, err := c.videoRepo.FindOneVideoByID(videoID)

	if err != nil {
		return nil, err
	}

	res := _video.NewVideoResponse(video)
	return &res, nil
}

func (c *videoService) UpdateVideo(updateVideoRequest request.UpdateVideoRequest) (*_video.VideoResponse, error) {
	video, err := c.videoRepo.FindOneVideoByID(fmt.Sprintf("%d", updateVideoRequest.ID))
	if err != nil {
		return nil, err
	}

	video = entity.Video{}
	err = smapping.FillStruct(&video, smapping.MapFields(&updateVideoRequest))

	if err != nil {
		return nil, err
	}

	video, err = c.videoRepo.UpdateVideo(video)

	if err != nil {
		return nil, err
	}

	res := _video.NewVideoResponse(video)
	return &res, nil
}

func (c *videoService) DeleteVideo(videoID string) error {
	c.videoRepo.DeleteVideo(videoID)
	return nil

}
