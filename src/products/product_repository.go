package products

import (
	"fmt"
	"log"

	"github.com/Javieradel/api-qisur.git/src/categories"
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

// ! Product updated are be inserted on last
func (r *ProductRepository) FindAll(criteria []shared.Criterion) ([]Product, error) {
	var products []Product
	query := r.DB.Model(&Product{}).Preload("Categories")

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
	if err := r.DB.Preload("Categories").First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Update(product *Product) (*Product, error) {
	if err := r.DB.Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) Delete(id uint) error {
	return r.DB.Delete(&Product{}, id).Error
}

func (r *ProductRepository) UpdateCategories(product *Product, categories []categories.Categories) error {
	return r.DB.Model(&product).Association("Categories").Replace(categories)
}
