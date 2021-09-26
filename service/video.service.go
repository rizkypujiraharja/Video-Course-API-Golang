package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/repo"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/request"
)

type VideoService interface {
	CreateVideo(videoRequest request.CreateVideoRequest) (*entity.Video, error)
	UpdateVideo(updateVideoRequest request.UpdateVideoRequest) (*entity.Video, error)
	FindOneVideoByID(videoID string) (*entity.Video, error)
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

func (c *videoService) CreateVideo(videoRequest request.CreateVideoRequest) (*entity.Video, error) {
	video := entity.Video{}
	err := smapping.FillStruct(&video, smapping.MapFields(&videoRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	vid, err := c.videoRepo.InsertVideo(video)
	if err != nil {
		return nil, err
	}

	return &vid, nil
}

func (c *videoService) FindOneVideoByID(videoID string) (*entity.Video, error) {
	video, err := c.videoRepo.FindOneVideoByID(videoID)

	if err != nil {
		return nil, err
	}

	return &video, nil
}

func (c *videoService) UpdateVideo(updateVideoRequest request.UpdateVideoRequest) (*entity.Video, error) {
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

	return &video, nil
}

func (c *videoService) DeleteVideo(videoID string) error {
	c.videoRepo.DeleteVideo(videoID)
	return nil

}
