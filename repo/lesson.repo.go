package repo

import (
	"fmt"

	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
	"gorm.io/gorm"
)

type LessonRepository interface {
	All() ([]entity.Lesson, error)
	InsertLesson(lesson entity.Lesson) (entity.Lesson, error)
	UpdateLesson(lesson entity.Lesson) (entity.Lesson, error)
	DeleteLesson(lessonID string) error
	FindOneLessonByID(ID string) (entity.Lesson, error)
	FindOneLessonByIDPublic(ID string) (entity.Lesson, error)
	FindLessonByIDS(LessonIDS []int) ([]entity.Lesson, error)
}

type lessonRepo struct {
	connection *gorm.DB
}

func NewLessonRepo(connection *gorm.DB) LessonRepository {
	return &lessonRepo{
		connection: connection,
	}
}

func (c *lessonRepo) All() ([]entity.Lesson, error) {
	lessons := []entity.Lesson{}
	c.connection.Preload("Category").Find(&lessons)
	return lessons, nil
}

func (c *lessonRepo) InsertLesson(lesson entity.Lesson) (entity.Lesson, error) {
	fmt.Println(lesson)
	c.connection.Save(&lesson)
	c.connection.Preload("Category").Find(&lesson)
	return lesson, nil
}

func (c *lessonRepo) UpdateLesson(lesson entity.Lesson) (entity.Lesson, error) {
	c.connection.Save(&lesson)
	c.connection.Preload("Category").Find(&lesson)
	return lesson, nil
}

func (c *lessonRepo) FindOneLessonByID(lessonID string) (entity.Lesson, error) {
	var lesson entity.Lesson
	res := c.connection.Preload("Category").Preload("SubLessons.Videos").Where("id = ?", lessonID).Take(&lesson)
	if res.Error != nil {
		return lesson, res.Error
	}
	return lesson, nil
}

func (c *lessonRepo) FindOneLessonByIDPublic(lessonID string) (entity.Lesson, error) {
	var lesson entity.Lesson
	res := c.connection.Preload("Category").Preload("SubLessons").Where("id = ?", lessonID).Take(&lesson)
	if res.Error != nil {
		return lesson, res.Error
	}
	return lesson, nil
}

func (c *lessonRepo) DeleteLesson(lessonID string) error {
	var lesson entity.Lesson
	res := c.connection.Preload("Category").Where("id = ?", lessonID).Take(&lesson)
	if res.Error != nil {
		return res.Error
	}
	c.connection.Delete(&lesson)
	return nil
}

func (c *lessonRepo) FindLessonByIDS(lessonIDS []int) ([]entity.Lesson, error) {
	var lessons []entity.Lesson
	res := c.connection.Find(&lessons, lessonIDS)
	if res.Error != nil {
		return lessons, res.Error
	}
	return lessons, nil
}
