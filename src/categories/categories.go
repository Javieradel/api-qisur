package categories

import (
	"time"
)

type Categories struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Name        string
	Description string
}

func (Categories) TableName() string {
	return "categories"
}
