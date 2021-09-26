package repo

import (
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
	"gorm.io/gorm"
)

type VideoRepository interface {
	InsertVideo(video entity.Video) (entity.Video, error)
	UpdateVideo(video entity.Video) (entity.Video, error)
	DeleteVideo(videoID string) error
	FindOneVideoByID(ID string) (entity.Video, error)
}

type videoRepo struct {
	connection *gorm.DB
}

func NewVideoRepo(connection *gorm.DB) VideoRepository {
	return &videoRepo{
		connection: connection,
	}
}

func (c *videoRepo) InsertVideo(video entity.Video) (entity.Video, error) {
	c.connection.Save(&video)
	c.connection.Find(&video)
	return video, nil
}

func (c *videoRepo) UpdateVideo(video entity.Video) (entity.Video, error) {
	c.connection.Save(&video)
	c.connection.Find(&video)
	return video, nil
}

func (c *videoRepo) FindOneVideoByID(videoID string) (entity.Video, error) {
	var video entity.Video
	res := c.connection.Where("id = ?", videoID).Take(&video)
	if res.Error != nil {
		return video, res.Error
	}
	return video, nil
}

func (c *videoRepo) DeleteVideo(videoID string) error {
	var video entity.Video
	res := c.connection.Where("id = ?", videoID).Take(&video)
	if res.Error != nil {
		return res.Error
	}
	c.connection.Delete(&video)
	return nil
}
