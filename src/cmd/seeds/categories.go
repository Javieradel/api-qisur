package main

import (
	"fmt"

	"github.com/Javieradel/api-qisur.git/src/categories"
	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

var CategoriesSeed = Seed{
	Name: "categories",
	Run: func(db *gorm.DB) error {
		db.AutoMigrate(&categories.Categories{})

		fmt.Println("Seeding 100 categories...")

		for i := 0; i < 100; i++ {
			category := categories.Categories{
				Name:        faker.Word(),
				Description: faker.Sentence(),
			}
			if err := db.Create(&category).Error; err != nil {
				return fmt.Errorf("failed to create category %d: %w", i, err)
			}
		}

		fmt.Println("100 categories seeded successfully!")
		return nil
	},
}
