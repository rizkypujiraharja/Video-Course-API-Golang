package service

import (
	"fmt"
	"log"
	"strings"

	"github.com/mashingan/smapping"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/repo"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/request"
)

type CategoryService interface {
	All() (*[]entity.Category, error)
	CreateCategory(categoryRequest request.CreateCategoryRequest) (*entity.Category, error)
	UpdateCategory(updateCategoryRequest request.UpdateCategoryRequest) (*entity.Category, error)
	FindOneCategoryByID(categoryID string) (*entity.Category, error)
	DeleteCategory(categoryID string) error
}

type categoryService struct {
	categoryRepo repo.CategoryRepository
}

func createSlug(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}

func (c *categoryService) All() (*[]entity.Category, error) {
	categories, err := c.categoryRepo.All()
	if err != nil {
		return nil, err
	}

	return &categories, nil
}

func NewCategoryService(categoryRepo repo.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (c *categoryService) CreateCategory(categoryRequest request.CreateCategoryRequest) (*entity.Category, error) {
	category := entity.Category{}
	err := smapping.FillStruct(&category, smapping.MapFields(&categoryRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}
	category.Slug = createSlug(category.Name)
	cat, err := c.categoryRepo.InsertCategory(category)
	if err != nil {
		return nil, err
	}

	return &cat, nil
}

func (c *categoryService) FindOneCategoryByID(categoryID string) (*entity.Category, error) {
	category, err := c.categoryRepo.FindOneCategoryByID(categoryID)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (c *categoryService) UpdateCategory(updateCategoryRequest request.UpdateCategoryRequest) (*entity.Category, error) {
	category, err := c.categoryRepo.FindOneCategoryByID(fmt.Sprintf("%d", updateCategoryRequest.ID))
	if err != nil {
		return nil, err
	}

	category = entity.Category{}
	err = smapping.FillStruct(&category, smapping.MapFields(&updateCategoryRequest))

	if err != nil {
		return nil, err
	}

	category.Slug = createSlug(category.Name)
	fmt.Println(category)
	category, err = c.categoryRepo.UpdateCategory(category)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (c *categoryService) DeleteCategory(categoryID string) error {
	_, err := c.categoryRepo.FindOneCategoryByID(categoryID)
	if err != nil {
		return err
	}

	c.categoryRepo.DeleteCategory(categoryID)
	return nil

}
