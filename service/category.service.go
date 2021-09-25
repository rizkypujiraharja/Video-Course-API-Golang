package service

import (
	"fmt"
	"log"
	"strings"

	"github.com/mashingan/smapping"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/repo"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/request"

	_category "github.com/rizkypujiraharja/Video-Course-API-Golang/service/category"
)

type CategoryService interface {
	All() (*[]_category.CategoryResponse, error)
	CreateCategory(categoryRequest request.CreateCategoryRequest) (*_category.CategoryResponse, error)
	UpdateCategory(updateCategoryRequest request.UpdateCategoryRequest) (*_category.CategoryResponse, error)
	FindOneCategoryByID(categoryID string) (*_category.CategoryResponse, error)
	DeleteCategory(categoryID string) error
}

type categoryService struct {
	categoryRepo repo.CategoryRepository
}

func createSlug(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}

func (c *categoryService) All() (*[]_category.CategoryResponse, error) {
	categories, err := c.categoryRepo.All()
	if err != nil {
		return nil, err
	}

	resCategories := _category.NewCategoryArrayResponse(categories)
	return &resCategories, nil
}

func NewCategoryService(categoryRepo repo.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (c *categoryService) CreateCategory(categoryRequest request.CreateCategoryRequest) (*_category.CategoryResponse, error) {
	category := entity.Category{}
	err := smapping.FillStruct(&category, smapping.MapFields(&categoryRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}
	category.Slug = createSlug(category.Name)
	p, err := c.categoryRepo.InsertCategory(category)
	if err != nil {
		return nil, err
	}

	res := _category.NewCategoryResponse(p)
	return &res, nil
}

func (c *categoryService) FindOneCategoryByID(categoryID string) (*_category.CategoryResponse, error) {
	category, err := c.categoryRepo.FindOneCategoryByID(categoryID)

	if err != nil {
		return nil, err
	}

	res := _category.NewCategoryResponse(category)
	return &res, nil
}

func (c *categoryService) UpdateCategory(updateCategoryRequest request.UpdateCategoryRequest) (*_category.CategoryResponse, error) {
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

	res := _category.NewCategoryResponse(category)
	return &res, nil
}

func (c *categoryService) DeleteCategory(categoryID string) error {
	_, err := c.categoryRepo.FindOneCategoryByID(categoryID)
	if err != nil {
		return err
	}

	c.categoryRepo.DeleteCategory(categoryID)
	return nil

}
