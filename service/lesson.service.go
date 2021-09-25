package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/repo"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/request"

	_lesson "github.com/rizkypujiraharja/Video-Course-API-Golang/service/lesson"
)

type LessonService interface {
	All() (*[]_lesson.LessonResponse, error)
	CreateLesson(lessonRequest request.CreateLessonRequest) (*_lesson.LessonResponse, error)
	UpdateLesson(updateLessonRequest request.UpdateLessonRequest) (*_lesson.LessonResponse, error)
	FindOneLessonByID(lessonID string) (*_lesson.LessonResponse, error)
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

func (c *lessonService) All() (*[]_lesson.LessonResponse, error) {
	lessons, err := c.lessonRepo.All()
	if err != nil {
		return nil, err
	}

	lessonsRes := _lesson.NewLessonArrayResponse(lessons)
	return &lessonsRes, nil
}

func (c *lessonService) CreateLesson(lessonRequest request.CreateLessonRequest) (*_lesson.LessonResponse, error) {
	lesson := entity.Lesson{}
	err := smapping.FillStruct(&lesson, smapping.MapFields(&lessonRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	p, err := c.lessonRepo.InsertLesson(lesson)
	if err != nil {
		return nil, err
	}

	res := _lesson.NewLessonResponse(p)
	return &res, nil
}

func (c *lessonService) FindOneLessonByID(lessonID string) (*_lesson.LessonResponse, error) {
	lesson, err := c.lessonRepo.FindOneLessonByID(lessonID)

	if err != nil {
		return nil, err
	}

	res := _lesson.NewLessonResponse(lesson)
	return &res, nil
}

func (c *lessonService) UpdateLesson(updateLessonRequest request.UpdateLessonRequest) (*_lesson.LessonResponse, error) {
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

	res := _lesson.NewLessonResponse(lesson)
	return &res, nil
}

func (c *lessonService) DeleteLesson(lessonID string) error {
	c.lessonRepo.DeleteLesson(lessonID)
	return nil

}
