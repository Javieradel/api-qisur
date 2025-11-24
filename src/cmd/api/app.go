package main

import (
	"log"

	_ "github.com/Javieradel/api-qisur.git/docs"
	"github.com/Javieradel/api-qisur.git/src/db"
	"github.com/Javieradel/api-qisur.git/src/products"
	swaggo "github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
)

// @title           API Example
// @version         1.0
// @description     Example API with Fiber and Swagger.
// @host            localhost:3000
// @BasePath        /api/
func main() {
	db.InitDB()

	//TODO add a container to DI
	productRepo := products.NewProductRepository(db.DB)
	productService := products.NewProductService(productRepo)

	app := fiber.New()

	app.Get("/api/docs/*", swaggo.HandlerDefault)

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("HELLO")
	})

	productController := products.NewProductController(productService)
	productController.RegisterRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
