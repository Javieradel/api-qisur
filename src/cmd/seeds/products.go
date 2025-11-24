package main

import (
	"fmt"
	"math/rand"

	"github.com/Javieradel/api-qisur.git/src/products"
	"github.com/go-faker/faker/v4"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

var ProductsSeed = Seed{
	Name: "products",
	Run: func(db *gorm.DB) error {
		db.AutoMigrate(&products.Product{})

		fmt.Println("Seeding 100 products...")

		for i := 0; i < 100; i++ {
			price := float64(rand.Intn(100000)) / 100.0
			product := products.Product{
				Name:        faker.Word(),
				Description: faker.Sentence(),
				Price:       decimal.NewFromFloat(price),
				Stock:       rand.Intn(1000),
			}
			if err := db.Create(&product).Error; err != nil {
				return fmt.Errorf("failed to create product %d: %w", i, err)
			}
		}

		fmt.Println("100 products seeded successfully!")
		return nil
	},
}
