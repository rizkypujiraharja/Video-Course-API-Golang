package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/repo"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/request"
)

type SubLessonService interface {
	CreateSubLesson(subLessonRequest request.CreateSubLessonRequest) (*entity.SubLesson, error)
	UpdateSubLesson(updateSubLessonRequest request.UpdateSubLessonRequest) (*entity.SubLesson, error)
	FindOneSubLessonByID(subLessonID string) (*entity.SubLesson, error)
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

func (c *subLessonService) CreateSubLesson(subLessonRequest request.CreateSubLessonRequest) (*entity.SubLesson, error) {
	subLesson := entity.SubLesson{}
	err := smapping.FillStruct(&subLesson, smapping.MapFields(&subLessonRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	sub, err := c.subLessonRepo.InsertSubLesson(subLesson)
	if err != nil {
		return nil, err
	}

	return &sub, nil
}

func (c *subLessonService) FindOneSubLessonByID(subLessonID string) (*entity.SubLesson, error) {
	subLesson, err := c.subLessonRepo.FindOneSubLessonByID(subLessonID)

	if err != nil {
		return nil, err
	}

	return &subLesson, nil
}

func (c *subLessonService) UpdateSubLesson(updateSubLessonRequest request.UpdateSubLessonRequest) (*entity.SubLesson, error) {
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

	return &subLesson, nil
}

func (c *subLessonService) DeleteSubLesson(subLessonID string) error {
	c.subLessonRepo.DeleteSubLesson(subLessonID)
	return nil

}
