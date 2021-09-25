package repo

import (
	"fmt"

	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
	"gorm.io/gorm"
)

type SubLessonRepository interface {
	InsertSubLesson(subLesson entity.SubLesson) (entity.SubLesson, error)
	UpdateSubLesson(subLesson entity.SubLesson) (entity.SubLesson, error)
	DeleteSubLesson(subLessonID string) error
	FindOneSubLessonByID(ID string) (entity.SubLesson, error)
}

type subLessonRepo struct {
	connection *gorm.DB
}

func NewSubLessonRepo(connection *gorm.DB) SubLessonRepository {
	return &subLessonRepo{
		connection: connection,
	}
}

func (c *subLessonRepo) InsertSubLesson(subLesson entity.SubLesson) (entity.SubLesson, error) {
	fmt.Println(subLesson)
	c.connection.Save(&subLesson)
	c.connection.Find(&subLesson)
	return subLesson, nil
}

func (c *subLessonRepo) UpdateSubLesson(subLesson entity.SubLesson) (entity.SubLesson, error) {
	c.connection.Save(&subLesson)
	c.connection.Find(&subLesson)
	return subLesson, nil
}

func (c *subLessonRepo) FindOneSubLessonByID(subLessonID string) (entity.SubLesson, error) {
	var subLesson entity.SubLesson
	res := c.connection.Where("id = ?", subLessonID).Take(&subLesson)
	if res.Error != nil {
		return subLesson, res.Error
	}
	return subLesson, nil
}

func (c *subLessonRepo) DeleteSubLesson(subLessonID string) error {
	var subLesson entity.SubLesson
	res := c.connection.Where("id = ?", subLessonID).Take(&subLesson)
	if res.Error != nil {
		return res.Error
	}
	c.connection.Delete(&subLesson)
	return nil
}
