package services

import (
	"menu-delivery/config"
	"menu-delivery/models"
)

type CategoryInputDetails struct {
	Name *string `json:"name"`
}

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := config.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func GetCategoryById(id string) (models.Category, error) {
	var category models.Category

	if err := config.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func CreateCategory(input CategoryInputDetails) (*models.Category, error) {
	category := models.Category{}

	if input.Name != nil {
		category.Name = *input.Name
	}

	if err := config.DB.Create(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func UpdateCategoryById(id string, input CategoryInputDetails) error {
	var category models.Category

	if err := config.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return err
	}

	if input.Name != nil {
		category.Name = *input.Name
	}

	if err := config.DB.Save(&category).Error; err != nil {
		return err
	}

	return nil
}

func DeleteCategoryById(id string) error {
	var category models.Category

	if err := config.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return err
	}

	if err := config.DB.Delete(&category).Error; err != nil {
		return err
	}

	return nil
}
