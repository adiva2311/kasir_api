package repositories

import (
	"kasir_api/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(category models.Category) error
	UpdateCategory(id int, category models.Category) error
	DeleteCategory(id int) error
	GetCategoryByID(id int) (*models.Category, error)
	GetAllCategories() ([]models.Category, error)
	FindByName(name string) (*models.Category, error)
}

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

// FindByName implements CategoryRepository.
func (c *CategoryRepositoryImpl) FindByName(name string) (*models.Category, error) {
	var category models.Category
	result := c.DB.Where("categories_name = ?", name).Where("deleted_at is NULL").First(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

// CreateCategory implements CategoryRepository.
func (c *CategoryRepositoryImpl) CreateCategory(category models.Category) error {
	return c.DB.Create(&category).Error
}

// DeleteCategory implements CategoryRepository.
func (c *CategoryRepositoryImpl) DeleteCategory(id int) error {
	result := c.DB.Where("id = ?", id).Delete(&models.Category{})
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

// GetAllCategories implements CategoryRepository.
func (c *CategoryRepositoryImpl) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	result := c.DB.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

// GetCategoryByID implements CategoryRepository.
func (c *CategoryRepositoryImpl) GetCategoryByID(id int) (*models.Category, error) {
	var category *models.Category
	result := c.DB.Where("id = ?", id).First(&category)
	if result.Error != nil {
		return &models.Category{}, result.Error
	}
	return category, nil
}

// UpdateCategory implements CategoryRepository.
func (c *CategoryRepositoryImpl) UpdateCategory(id int, category models.Category) error {
	result := c.DB.Where("id = ?", id).Updates(category)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		DB: db,
	}
}
