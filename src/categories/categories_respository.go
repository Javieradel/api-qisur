package categories

import (
	"fmt"
	"log"

	"github.com/Javieradel/api-qisur.git/src/shared"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) Create(category *Categories) error {
	return r.DB.Create(category).Error
}

func (r *CategoryRepository) FindAll(criteria []shared.Criterion) ([]Categories, error) {
	var categories []Categories
	query := r.DB.Model(&Categories{})

	if len(criteria) > 0 {
		query = shared.ApplyCriteria(query, criteria)
	}

	err := query.Find(&categories).Error

	if err != nil {
		log.Printf("Error fetching categories %+v: %v", criteria, err)
		return nil, fmt.Errorf("failed to fetch categories: %w", err)
	}

	return categories, nil
}

func (r *CategoryRepository) FindByID(id uint) (*Categories, error) {
	var category Categories
	if err := r.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) Update(category *Categories) error {
	return r.DB.Save(category).Error
}

func (r *CategoryRepository) Delete(id uint) error {
	return r.DB.Delete(&Categories{}, id).Error
}