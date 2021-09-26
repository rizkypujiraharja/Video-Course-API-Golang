package repo

import (
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
	"gorm.io/gorm"
)

type OrderedLessonRepository interface {
	InsertOrderedLesson(orderedLesson entity.OrderedLesson) (entity.OrderedLesson, error)
	FindOrderedLessonByUserID(userID string) ([]entity.OrderedLesson, error)
	DeleteOrderedLessonByUserID(userID string) error
}

type orderedLessonRepo struct {
	connection *gorm.DB
}

func NewOrderedLessonRepo(connection *gorm.DB) OrderedLessonRepository {
	return &orderedLessonRepo{
		connection: connection,
	}
}

func (c *orderedLessonRepo) InsertOrderedLesson(orderedLesson entity.OrderedLesson) (entity.OrderedLesson, error) {
	c.connection.Save(&orderedLesson)
	c.connection.Find(&orderedLesson)
	return orderedLesson, nil
}

func (c *orderedLessonRepo) FindOrderedLessonByUserID(userID string) ([]entity.OrderedLesson, error) {
	orderedLessons := []entity.OrderedLesson{}
	c.connection.Preload("Lesson.Category").Where("user_id = ?", userID).Find(&orderedLessons)
	return orderedLessons, nil
}

func (c *orderedLessonRepo) DeleteOrderedLessonByUserID(userID string) error {
	orderedLessons := entity.OrderedLesson{}
	c.connection.Where("user_id = ?", userID).Delete(&orderedLessons)
	return nil
}
