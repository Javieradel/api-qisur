package products

import (
	"fmt"
	"log"

	"github.com/Javieradel/api-qisur.git/src/shared"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) Create(product *Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductRepository) FindAll(criteria []shared.Criterion) ([]Product, error) {
	var products []Product
	query := r.DB.Model(&Product{}) // Partimos de la query base

	if len(criteria) > 0 {
		query = shared.ApplyCriteria(query, criteria)
	}

	err := query.Find(&products).Error

	if err != nil {
		log.Printf("Error fetching products %+v: %v", criteria, err)
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}

	return products, nil
}

func (r *ProductRepository) FindByID(id uint) (*Product, error) {
	var product Product
	if err := r.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Update(product *Product) error {
	return r.DB.Save(product).Error
}

func (r *ProductRepository) Delete(id uint) error {
	return r.DB.Delete(&Product{}, id).Error
}
