package main

import (
	"log"

	_ "github.com/Javieradel/api-qisur.git/docs"
	"github.com/Javieradel/api-qisur.git/src/categories"
	"github.com/Javieradel/api-qisur.git/src/db"
	"github.com/Javieradel/api-qisur.git/src/products"
	swaggo "github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

// @title           API Example
// @version         1.0
// @description     Example API with Fiber and Swagger.
// @host            localhost:3000
// @BasePath        /api/v1
// ? Swagger retrieve 302 code status??
func main() {
	db.InitDB()

	//TODO add a container to DI
	productRepo := products.NewProductRepository(db.DB)
	productService := products.NewProductService(productRepo)
	categoryRepo := categories.NewCategoryRepository(db.DB)
	categoryService := categories.NewCategoryService(categoryRepo)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	app.Get("/api/docs/*", swaggo.HandlerDefault)

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("HELLO")
	})

	productController := products.NewProductController(productService)
	productController.RegisterRoutes(app)
	categoryController := categories.NewCategoryController(categoryService)
	categoryController.RegisterRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
