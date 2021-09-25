package repo

import (
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	InsertCategory(category entity.Category) (entity.Category, error)
	UpdateCategory(category entity.Category) (entity.Category, error)
	DeleteCategory(categoryID string) error
	FindOneCategoryByID(categoryID string) (entity.Category, error)
	All() ([]entity.Category, error)
}

type categoryRepo struct {
	connection *gorm.DB
}

func NewCategoryRepo(connection *gorm.DB) CategoryRepository {
	return &categoryRepo{
		connection: connection,
	}
}

func (c *categoryRepo) All() ([]entity.Category, error) {
	categories := []entity.Category{}
	c.connection.Find(&categories)
	return categories, nil
}

func (c *categoryRepo) InsertCategory(category entity.Category) (entity.Category, error) {
	c.connection.Save(&category)
	c.connection.Find(&category)
	return category, nil
}

func (c *categoryRepo) UpdateCategory(category entity.Category) (entity.Category, error) {
	c.connection.Save(&category)
	c.connection.Find(&category)
	return category, nil
}

func (c *categoryRepo) FindOneCategoryByID(categoryID string) (entity.Category, error) {
	var category entity.Category
	res := c.connection.Where("id = ?", categoryID).Take(&category)
	if res.Error != nil {
		return category, res.Error
	}
	return category, nil
}

func (c *categoryRepo) DeleteCategory(categoryID string) error {
	var category entity.Category
	res := c.connection.Where("id = ?", categoryID).Take(&category)
	if res.Error != nil {
		return res.Error
	}
	c.connection.Delete(&category)
	return nil
}
