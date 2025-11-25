package main

import (
	"fmt"
	"os"

	"github.com/Javieradel/api-qisur.git/src/db"
	"gorm.io/gorm"
)

type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}

var seeds []Seed

func registerSeeds() {
	seeds = append(seeds, ProductsSeed)
	seeds = append(seeds, CategoriesSeed)
}

func main() {
	db.InitDB()
	registerSeeds()
	if len(os.Args) > 1 {
		seedName := os.Args[1]
		found := false
		for _, seed := range seeds {
			if seed.Name == seedName {
				fmt.Printf("Running seed: %s\n", seed.Name)
				if err := seed.Run(db.DB); err != nil {
					fmt.Printf("Error running seed %s: %v\n", seed.Name, err)
				}
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Seed not found: %s\n", seedName)
		}
		return
	}

	fmt.Println("Running all seeds", len(seeds))

	for _, seed := range seeds {
		fmt.Printf("Running seed: %s\n", seed.Name)
		if err := seed.Run(db.DB); err != nil {
			fmt.Printf("Error running seed %s: %v\n", seed.Name, err)
		}
	}
}
