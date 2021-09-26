package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/repo"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/request"
)

type LessonService interface {
	All() (*[]entity.Lesson, error)
	CreateLesson(lessonRequest request.CreateLessonRequest) (*entity.Lesson, error)
	UpdateLesson(updateLessonRequest request.UpdateLessonRequest) (*entity.Lesson, error)
	FindOneLessonByID(lessonID string) (*entity.Lesson, error)
	DeleteLesson(lessonID string) error
}

type lessonService struct {
	lessonRepo repo.LessonRepository
}

func NewLessonService(lessonRepo repo.LessonRepository) LessonService {
	return &lessonService{
		lessonRepo: lessonRepo,
	}
}

func (c *lessonService) All() (*[]entity.Lesson, error) {
	lessons, err := c.lessonRepo.All()
	if err != nil {
		return nil, err
	}

	return &lessons, nil
}

func (c *lessonService) CreateLesson(lessonRequest request.CreateLessonRequest) (*entity.Lesson, error) {
	lesson := entity.Lesson{}
	err := smapping.FillStruct(&lesson, smapping.MapFields(&lessonRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	les, err := c.lessonRepo.InsertLesson(lesson)
	if err != nil {
		return nil, err
	}

	return &les, nil
}

func (c *lessonService) FindOneLessonByID(lessonID string) (*entity.Lesson, error) {
	lesson, err := c.lessonRepo.FindOneLessonByID(lessonID)

	if err != nil {
		return nil, err
	}

	return &lesson, nil
}

func (c *lessonService) UpdateLesson(updateLessonRequest request.UpdateLessonRequest) (*entity.Lesson, error) {
	lesson, err := c.lessonRepo.FindOneLessonByID(fmt.Sprintf("%d", updateLessonRequest.ID))
	if err != nil {
		return nil, err
	}

	lesson = entity.Lesson{}
	err = smapping.FillStruct(&lesson, smapping.MapFields(&updateLessonRequest))

	if err != nil {
		return nil, err
	}

	lesson, err = c.lessonRepo.UpdateLesson(lesson)

	if err != nil {
		return nil, err
	}

	return &lesson, nil
}

func (c *lessonService) DeleteLesson(lessonID string) error {
	c.lessonRepo.DeleteLesson(lessonID)
	return nil

}
