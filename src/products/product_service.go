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

func (s *ProductService) FindByID(id uint) (*Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) Create(product *Product) error {
	//TODO add bussines validations
	return s.repo.Create(product)
}
