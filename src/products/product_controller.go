package products

import (
	"github.com/gofiber/fiber/v3"
)

type ProductController struct{}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (pc *ProductController) RegisterRoutes(app *fiber.App) {
	app.Get("/products", pc.GetProducts)
}

func (pc *ProductController) GetProducts(c fiber.Ctx) error {
	return c.SendString("ok")
}
