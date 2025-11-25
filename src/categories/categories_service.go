package categories

import (
	"github.com/Javieradel/api-qisur.git/src/shared"
)

type CategoryService struct {
	repo *CategoryRepository
}

func NewCategoryService(repo *CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) FindAll(filters []shared.Criterion) ([]Categories, error) {
	return s.repo.FindAll(filters)
}

func (s *CategoryService) FindByID(id uint) (*Categories, error) {
	return s.repo.FindByID(id)
}

func (s *CategoryService) Create(category *Categories) error {
	//TODO add bussines validations
	return s.repo.Create(category)
}

func (s *CategoryService) Update(category *Categories) error {
	//TODO add bussines validations
	return s.repo.Update(category)
}

func (s *CategoryService) Delete(id uint) error {
	return s.repo.Delete(id)
}