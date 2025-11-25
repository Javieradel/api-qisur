package products

import (
	"time"

	"github.com/Javieradel/api-qisur.git/src/categories"
	"github.com/Javieradel/api-qisur.git/src/shared"
)

type ProductService struct {
	repo     *ProductRepository
	eventBus *shared.EventBus
}

func NewProductService(repo *ProductRepository, eventBus *shared.EventBus) *ProductService {
	return &ProductService{repo: repo, eventBus: eventBus}
}

func (s *ProductService) FindAll(filters []shared.Criterion) ([]Product, error) {
	return s.repo.FindAll(filters)
}

func (s *ProductService) FindByID(id uint) (*Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) Create(product *Product) (*Product, error) {
	if err := s.repo.Create(product); err != nil {
		return nil, err
	}
	s.eventBus.Publish(ProductCreatedEvent{Product: *product})
	return product, nil
}

func (s *ProductService) Update(product *Product) (*Product, error) {
	oldProduct, err := s.repo.FindByID(product.ID)
	if err != nil {
		return nil, err
	}
	updatedProduct, err := s.repo.Update(product)
	if err != nil {
		return nil, err
	}
	s.eventBus.Publish(ProductUpdatedEvent{OldProduct: *oldProduct, NewProduct: *updatedProduct})
	return updatedProduct, nil
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

func (s *ProductService) Delete(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	s.eventBus.Publish(ProductDeletedEvent{ProductID: id})
	return nil
}

func (s *ProductService) FindHistoryByProductID(productID uint, start, end *time.Time) ([]ProductHistory, error) {
	return s.repo.FindHistoryByProductID(productID, start, end)
}
