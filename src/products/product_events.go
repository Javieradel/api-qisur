package products

import "github.com/Javieradel/api-qisur.git/src/shared"

// ProductCreatedEvent is published when a product is created
type ProductCreatedEvent struct {
	shared.Event
	Product Product
}

func (e ProductCreatedEvent) Topic() string {
	return "product.created"
}

// ProductUpdatedEvent is published when a product is updated
type ProductUpdatedEvent struct {
	shared.Event
	OldProduct Product
	NewProduct Product
}

func (e ProductUpdatedEvent) Topic() string {
	return "product.updated"
}

// ProductDeletedEvent is published when a product is deleted
type ProductDeletedEvent struct {
	shared.Event
	ProductID uint
}

func (e ProductDeletedEvent) Topic() string {
	return "product.deleted"
}
