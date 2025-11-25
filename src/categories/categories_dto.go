package categories

import (
	"github.com/Javieradel/api-qisur.git/src/shared"
)

type CategoryQueryDTO struct {
	Page        int    `query:"page" validate:"gte=0"`
	Limit       int    `query:"limit" validate:"gte=0,lte=100"`
	Name        string `query:"name"`
	Description string `query:"description"`
}

type CreateCategoryDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

func (dto *CreateCategoryDTO) ToCategory() *Categories {
	return &Categories{
		Name:        dto.Name,
		Description: dto.Description,
	}
}

type UpdateCategoryDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type PatchCategoryDTO struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (dto *CategoryQueryDTO) ToCriterions() []shared.Criterion {
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
