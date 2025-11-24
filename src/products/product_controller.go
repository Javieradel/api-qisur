package products

import (
	"github.com/Javieradel/api-qisur.git/src/shared"
	"github.com/gofiber/fiber/v3"
)

type ProductController struct {
	service *ProductService
}

func NewProductController(service *ProductService) *ProductController {
	return &ProductController{service: service}
}

func (pc *ProductController) RegisterRoutes(app *fiber.App) {
	app.Get("/products", pc.GetProducts)
}

func (pc *ProductController) GetProducts(c fiber.Ctx) error {
	var q ProductQueryDTO
	errQuery := c.Bind().Query(&q)
	if errQuery != nil {
		return shared.NewErrorResponse(c, fiber.StatusBadRequest, "Invalids query params")
	}
	//TODO check all filters, price field not work
	filters := q.ToCriterions()

	products, err := pc.service.FindAll(filters)
	if err != nil {
		return shared.NewErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch products")
	}

	if len(products) == 0 {
		return shared.NewErrorResponse(c, fiber.StatusNotFound, "Products not found")
	}

	if q.Limit > 0 || q.Page > 0 {
		return shared.NewPaginatedResponse(c, fiber.StatusFound, products, q.Page, q.Limit)
	}

	return shared.NewSuccessResponse(c, fiber.StatusFound, products)
}

/*
~~GET /api/products - Lista paginada de productos
GET /api/products/{id} - Detalle de producto
POST /api/products - Crear producto
 Qisur Challenge API REST y Webscoket para gestión de productos.
PUT /api/products/{id} - Actualizar producto
DELETE /api/products/{id} - Eliminar producto
GET /api/products/{id}/history?start={date}&end={date} – Historial del producto GET /api/categories - Lista de categorías
POST /api/categories - Crear categoría
PUT /api/categories/{id} - Actualizar categoría
DELETE /api/categories/{id} - Eliminar categoría GET/api/search?{product/category}&[params] – Buscar productos o categorías

*/
