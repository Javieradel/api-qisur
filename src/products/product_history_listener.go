package products

import (
	"fmt"
	"reflect"
	"time"

	"github.com/Javieradel/api-qisur.git/src/shared"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductHistoryListener struct {
	DB *gorm.DB
}

func NewProductHistoryListener(db *gorm.DB) *ProductHistoryListener {
	return &ProductHistoryListener{DB: db}
}

func (l *ProductHistoryListener) Handle(event shared.Event) {
	switch e := event.(type) {
	case ProductCreatedEvent:
		l.handleProductCreated(e)
	case ProductUpdatedEvent:
		l.handleProductUpdated(e)
	case ProductDeletedEvent:
	}
}

func (l *ProductHistoryListener) handleProductCreated(event ProductCreatedEvent) {
	history := ProductHistory{
		UUID:      uuid.New(),
		ProductID: event.Product.ID,
		ChangedAt: time.Now(),
	}
	if err := l.DB.Create(&history).Error; err != nil {
		fmt.Println("Error creating product history:", err)
		return
	}

	val := reflect.ValueOf(event.Product)
	typeOfS := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typeOfS.Field(i).Name
		if field == "ID" || field == "CreatedAt" || field == "UpdatedAt" || field == "DeletedAt" {
			continue
		}
		detail := ProductHistoryDetail{
			ProductHistoryID: history.ID,
			Field:            field,
			NewValue:         fmt.Sprintf("%v", val.Field(i).Interface()),
		}
		if err := l.DB.Create(&detail).Error; err != nil {
			fmt.Println("Error creating product history detail:", err)
		}
	}
}

func (l *ProductHistoryListener) handleProductUpdated(event ProductUpdatedEvent) {
	history := ProductHistory{
		UUID:      uuid.New(),
		ProductID: event.NewProduct.ID,
		ChangedAt: time.Now(),
	}
	if err := l.DB.Create(&history).Error; err != nil {
		fmt.Println("Error creating product history:", err)
		return
	}

	oldVal := reflect.ValueOf(event.OldProduct)
	newVal := reflect.ValueOf(event.NewProduct)
	typeOfS := oldVal.Type()

	for i := 0; i < oldVal.NumField(); i++ {
		field := typeOfS.Field(i).Name
		if field == "ID" || field == "CreatedAt" || field == "UpdatedAt" || field == "DeletedAt" {
			continue
		}

		oldValue := fmt.Sprintf("%v", oldVal.Field(i).Interface())
		newValue := fmt.Sprintf("%v", newVal.Field(i).Interface())

		if oldValue != newValue {
			oldValuePtr := &oldValue
			detail := ProductHistoryDetail{
				ProductHistoryID: history.ID,
				Field:            field,
				OldValue:         oldValuePtr,
				NewValue:         newValue,
			}
			if err := l.DB.Create(&detail).Error; err != nil {
				fmt.Println("Error creating product history detail:", err)
			}
		}
	}
}