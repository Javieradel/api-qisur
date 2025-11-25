package products

import (
	"github.com/Javieradel/api-qisur.git/src/categories"
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

func (s *ProductService) Update(product *Product) error {
	//TODO add bussines validations
	return s.repo.Update(product)
}

func (s *ProductService) UpdateCategories(product *Product, categoriesID []uint) error {
	var cats []categories.Categories
	if len(categoriesID) > 0 {
		if err := s.repo.DB.Find(&cats, categoriesID).Error; err != nil {
			return err
		}
	}
	return s.repo.UpdateCategories(product, cats)
}
