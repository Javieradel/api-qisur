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

func (s *ProductService) FindAll() ([]Product, error) {
	criteria := []shared.Criterion{}
	return s.repo.FindAll(criteria)
}
