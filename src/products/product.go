package products

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `gorm:"index"`
	Name        string
	Description string
	Price       decimal.Decimal `gorm:"type:decimal(10,2)"`
	Stock       int
}

// TableName overrides the table name used by Product to `products`
func (Product) TableName() string {
	return "products"
}
