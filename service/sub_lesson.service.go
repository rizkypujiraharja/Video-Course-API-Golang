package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/repo"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/request"

	_sub_lesson "github.com/rizkypujiraharja/Video-Course-API-Golang/service/sub_lesson"
)

type SubLessonService interface {
	CreateSubLesson(subLessonRequest request.CreateSubLessonRequest) (*_sub_lesson.SubLessonResponse, error)
	UpdateSubLesson(updateSubLessonRequest request.UpdateSubLessonRequest) (*_sub_lesson.SubLessonResponse, error)
	FindOneSubLessonByID(subLessonID string) (*_sub_lesson.SubLessonResponse, error)
	DeleteSubLesson(subLessonID string) error
}

type subLessonService struct {
	subLessonRepo repo.SubLessonRepository
}

func NewSubLessonService(subLessonRepo repo.SubLessonRepository) SubLessonService {
	return &subLessonService{
		subLessonRepo: subLessonRepo,
	}
}

func (c *subLessonService) CreateSubLesson(subLessonRequest request.CreateSubLessonRequest) (*_sub_lesson.SubLessonResponse, error) {
	subLesson := entity.SubLesson{}
	err := smapping.FillStruct(&subLesson, smapping.MapFields(&subLessonRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	p, err := c.subLessonRepo.InsertSubLesson(subLesson)
	if err != nil {
		return nil, err
	}

	res := _sub_lesson.NewSubLessonResponse(p)
	return &res, nil
}

func (c *subLessonService) FindOneSubLessonByID(subLessonID string) (*_sub_lesson.SubLessonResponse, error) {
	subLesson, err := c.subLessonRepo.FindOneSubLessonByID(subLessonID)

	if err != nil {
		return nil, err
	}

	res := _sub_lesson.NewSubLessonResponse(subLesson)
	return &res, nil
}

func (c *subLessonService) UpdateSubLesson(updateSubLessonRequest request.UpdateSubLessonRequest) (*_sub_lesson.SubLessonResponse, error) {
	subLesson, err := c.subLessonRepo.FindOneSubLessonByID(fmt.Sprintf("%d", updateSubLessonRequest.ID))
	if err != nil {
		return nil, err
	}

	subLesson = entity.SubLesson{}
	err = smapping.FillStruct(&subLesson, smapping.MapFields(&updateSubLessonRequest))

	if err != nil {
		return nil, err
	}

	subLesson, err = c.subLessonRepo.UpdateSubLesson(subLesson)

	if err != nil {
		return nil, err
	}

	res := _sub_lesson.NewSubLessonResponse(subLesson)
	return &res, nil
}

func (c *subLessonService) DeleteSubLesson(subLessonID string) error {
	c.subLessonRepo.DeleteSubLesson(subLessonID)
	return nil

}
