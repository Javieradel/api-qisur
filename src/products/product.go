package products

import (
	"time"

	"github.com/Javieradel/api-qisur.git/src/categories"
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
	Categories  []categories.Categories `gorm:"many2many:product_categories;joinForeignKey:ProductID;joinReferences:CategoryID"`
}

// TableName overrides the table name used by Product to `products`
func (Product) TableName() string {
	return "products"
}
