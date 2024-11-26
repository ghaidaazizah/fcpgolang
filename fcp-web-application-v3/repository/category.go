package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	"fmt"
)

type CategoryRepository interface {
	Store(Category *model.Category) error
	Update(id int, category model.Category) error
	Delete(id int) error
	GetByID(id int) (*model.Category, error)
	GetList() ([]model.Category, error)
}

type categoryRepository struct {
	filebasedDb *filebased.Data
}

func NewCategoryRepo(filebasedDb *filebased.Data) *categoryRepository {
	return &categoryRepository{filebasedDb}
}

func (c *categoryRepository) Store(Category *model.Category) error {
	c.filebasedDb.StoreCategory(*Category) 
	return nil
}

func (c *categoryRepository) Update(id int, category model.Category) error {
	existingCategory, err := c.filebasedDb.GetCategoryByID(id)
	if err != nil {
		return fmt.Errorf("category not found: %v", err)
	}

	existingCategory.Name = category.Name
	existingCategory.Description = category.Description

	err = c.filebasedDb.StoreCategory(*existingCategory)
	if err != nil {
		return fmt.Errorf("failed to update category: %v", err)
	}
	return nil
}

func (c *categoryRepository) Delete(id int) error {
	category, err := c.filebasedDb.GetCategoryByID(id)
	if err != nil {
		return fmt.Errorf("category not found: %v", err)
	}

	err = c.filebasedDb.DeleteCategory(*category)
	if err != nil {
		return fmt.Errorf("failed to delete category: %v", err)
	}
	return nil
}

func (c *categoryRepository) GetByID(id int) (*model.Category, error) {
	category, err := c.filebasedDb.GetCategoryByID(id)
	if err != nil {
		return nil, fmt.Errorf("category not found: %v", err)
	}
	return category, nil
}

func (c *categoryRepository) GetList() ([]model.Category, error) {
	categories, err := c.filebasedDb.GetAllCategories()
	if err != nil {
		return nil, fmt.Errorf("failed to get category list: %v", err)
	}
	return categories, nil
}
