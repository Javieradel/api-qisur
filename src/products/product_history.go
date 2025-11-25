package products

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ProductHistory represents the history of changes for a product
type ProductHistory struct {
	ID        uint      `gorm:"primaryKey"`
	UUID      uuid.UUID `gorm:"type:uuid;"`
	ProductID uint      `gorm:"index"`
	ChangedAt time.Time
	Details   []ProductHistoryDetail `gorm:"foreignKey:ProductHistoryID"`
}

func (p *ProductHistory) BeforeCreate(tx *gorm.DB) (err error) {
	if p.UUID == uuid.Nil {
		p.UUID = uuid.New()
	}
	return
}

func (ProductHistory) TableName() string {
	return "product_histories"
}

// ProductHistoryDetail represents the details of a specific field change within a product history entry
type ProductHistoryDetail struct {
	ID               uint    `gorm:"primaryKey"`
	ProductHistoryID uint    `gorm:"index"`
	Field            string  `gorm:"type:varchar(255)"`
	OldValue         *string `gorm:"type:text"`
	NewValue         string  `gorm:"type:text"`
}

// TableName overrides the table name used by ProductHistoryDetail to `product_history_details`
func (ProductHistoryDetail) TableName() string {
	return "product_history_details"
}
