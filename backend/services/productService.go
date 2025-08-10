package services

import (
	"app-fastmenu-backend/config"
	"app-fastmenu-backend/models"

	"github.com/shopspring/decimal"
) 

type ProductInputDetails struct {
    Name        *string          `json:"name"`
    Description *string          `json:"description"`
    ImageURL    *string          `json:"image_url"`
    Price       *decimal.Decimal `json:"price"`
}

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductById(id string) error {
	var product models.Product 

	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return err
	}

	return nil
}

func CreateProduct(input ProductInputDetails) (*models.Product, error) {
    product := models.Product{}

    if input.Name != nil {
        product.Name = *input.Name
    }
    if input.Description != nil {
        product.Description = *input.Description
    }
    if input.ImageURL != nil {
        product.ImageURL = *input.ImageURL
    }
    if input.Price != nil {
        product.Price = *input.Price
    }

    if err := config.DB.Create(&product).Error; err != nil {
        return nil, err
    }
    return &product, nil
}

func UpdateProductById(id string, input ProductInputDetails) error {
	var product models.Product

	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return err
	}

	if input.Name != nil {
		product.Name = *input.Name
	}
	if input.Description != nil {
		product.Description = *input.Description
	}
	if input.ImageURL != nil {
		product.ImageURL = *input.ImageURL
	}
	if input.Price != nil {
		product.Price = *input.Price
	}

	if err := config.DB.Save(&product).Error; err != nil {
		return err
	}

	return nil
}

func DeleteProductById(id string) error {
	var product models.Product

	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return err
	}

	if err := config.DB.Delete(&product).Error; err != nil {
		return err
	}

	return nil
}
