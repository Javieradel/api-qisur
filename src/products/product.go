package products

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       decimal.Decimal `gorm:"type:decimal(10,2)"`
	Stock       int
}

// TableName overrides the table name used by Product to `products`
func (Product) TableName() string {
	return "products"
}
