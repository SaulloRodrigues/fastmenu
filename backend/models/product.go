package models

import "github.com/shopspring/decimal"

type Product struct {
	ID          int             `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	ImageURL    string          `json:"image_url"`
	Price       decimal.Decimal `gorm:"type:decimal(10,2)" json:"price"`
	Categories  []Category      `gorm:"many2many:product_categories;" json:"categories"`
}
