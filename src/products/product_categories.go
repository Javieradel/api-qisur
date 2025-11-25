package products

import (
	"time"

	"gorm.io/gorm"
)

type ProductCategories struct {
	ProductID  uint `gorm:"primaryKey"`
	CategoryID uint `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (ProductCategories) TableName() string {
	return "product_categories"
}