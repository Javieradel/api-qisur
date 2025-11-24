package products

import (
	"github.com/Javieradel/api-qisur.git/src/shared"
)

type ProductService struct {
	repo *ProductRepository
}

func NewProductService(repo *ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) FindAll(filters []shared.Criterion) ([]Product, error) {
	return s.repo.FindAll(filters)
}
