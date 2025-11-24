package main

import (
	"log"

	"github.com/Javieradel/api-qisur.git/src/db"
	"github.com/Javieradel/api-qisur.git/src/products"
	"github.com/gofiber/fiber/v3"
)

func main() {
	db.InitDB()

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("HELLO")
	})

	productController := products.NewProductController()
	productController.RegisterRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
