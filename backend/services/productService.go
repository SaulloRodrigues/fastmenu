package services

import (
	"fmt"
	"menu-delivery/config"
	"menu-delivery/models"
	"github.com/shopspring/decimal"
)

type ProductInputDetails struct {
	Name        *string          `json:"name"`
	Description *string          `json:"description"`
	ImageURL    *string          `json:"image_url"`
	Price       *decimal.Decimal `json:"price"`
	CategoryIDs []int            `json:"category_ids"`
}

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := config.DB.Preload("Categories").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductById(id string) (models.Product, error) {
	var product models.Product

	if err := config.DB.Preload("Categories").Where("id = ?", id).First(&product).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func GetAllProductsByCategory(categoryID string) ([]models.Product, error) {
	var products []models.Product

	if err := config.DB.Preload("Categories").
		Joins("JOIN product_categories ON products.id = product_categories.product_id").
		Where("product_categories.category_id = ?", categoryID).
		Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
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

	// Associar categorias se IDs foram fornecidos
	if len(input.CategoryIDs) > 0 {
		var categories []models.Category
		if err := config.DB.Where("id IN ?", input.CategoryIDs).Find(&categories).Error; err != nil {
			return nil, err
		}

		// Verificar se todos os IDs fornecidos existem
		if len(categories) != len(input.CategoryIDs) {
			return nil, fmt.Errorf("algumas categorias não foram encontradas")
		}

		if err := config.DB.Model(&product).Association("Categories").Append(categories); err != nil {
			return nil, err
		}
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

	// Atualizar categorias se IDs foram fornecidos
	if input.CategoryIDs != nil {
		// Limpar associações existentes
		if err := config.DB.Model(&product).Association("Categories").Clear(); err != nil {
			return err
		}

		// Adicionar novas categorias
		if len(input.CategoryIDs) > 0 {
			var categories []models.Category
			if err := config.DB.Where("id IN ?", input.CategoryIDs).Find(&categories).Error; err != nil {
				return err
			}

			// Verificar se todos os IDs fornecidos existem
			if len(categories) != len(input.CategoryIDs) {
				return fmt.Errorf("algumas categorias não foram encontradas")
			}

			if err := config.DB.Model(&product).Association("Categories").Append(categories); err != nil {
				return err
			}
		}
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
