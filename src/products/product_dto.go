package products

import (
	"github.com/Javieradel/api-qisur.git/src/shared"
	"github.com/shopspring/decimal"
)

type ProductQueryDTO struct {
	Page        int              `query:"page" validate:"gte=0"`
	Limit       int              `query:"limit" validate:"gte=0,lte=100"`
	Name        string           `query:"name"`
	Description string           `query:"description"`
	PriceFrom   *decimal.Decimal `query:"price_from"`
	PriceTo     *decimal.Decimal `query:"price_to"`
	Stock       int              `query:"stock"`
}

type CreateProductDTO struct {
	Name        string          `json:"name" validate:"required"`
	Description string          `json:"description"`
	Price       decimal.Decimal `json:"price" validate:"gt=0"`
	Stock       int             `json:"stock" validate:"gte=0"`
}

func (dto *CreateProductDTO) ToProduct() *Product {
	return &Product{
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
		Stock:       dto.Stock,
	}
}

type UpdateProductDTO struct {
	Name        string          `json:"name" validate:"required"`
	Description string          `json:"description"`
	Price       decimal.Decimal `json:"price" validate:"gt=0"`
	Stock       int             `json:"stock" validate:"gte=0"`
}

type PatchProductDTO struct {
	Name        *string          `json:"name,omitempty"`
	Description *string          `json:"description,omitempty"`
	Price       *decimal.Decimal `json:"price,omitempty" validate:"omitempty,gt=0"`
	Stock       *int             `json:"stock,omitempty" validate:"omitempty,gte=0"`
}

// TODO abstract commons criteria & inherit it
func (dto *ProductQueryDTO) ToCriterions() []shared.Criterion {
	criterions := make([]shared.Criterion, 0)

	if dto.Name != "" {
		criterions = append(criterions, shared.Criterion{
			Field:    "name",
			Operator: shared.OpLike,
			Value:    "%" + dto.Name + "%",
		})
	}

	if dto.Description != "" {
		criterions = append(criterions, shared.Criterion{
			Field:    "description",
			Operator: shared.OpLike,
			Value:    "%" + dto.Description + "%",
		})
	}

	if dto.PriceFrom != nil {
		criterions = append(criterions, shared.Criterion{
			Field:    "price",
			Operator: shared.OpGte,
			Value:    *dto.PriceFrom,
		})
	}

	if dto.PriceTo != nil {
		criterions = append(criterions, shared.Criterion{
			Field:    "price",
			Operator: shared.OpLte,
			Value:    *dto.PriceTo,
		})
	}

	if dto.Stock > 0 {
		criterions = append(criterions, shared.Criterion{
			Field:    "stock",
			Operator: shared.OpGte,
			Value:    dto.Stock,
		})
	}

	limit := dto.Limit
	if limit <= 0 {
		limit = 10
	}
	criterions = append(criterions, shared.Criterion{
		Operator: shared.OpLimit,
		Value:    limit,
	})

	page := dto.Page
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit
	criterions = append(criterions, shared.Criterion{
		Operator: shared.OpOffset,
		Value:    offset,
	})

	return criterions
}
