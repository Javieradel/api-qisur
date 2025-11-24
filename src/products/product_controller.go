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
	app.Get("/api/v1/products", pc.GetProducts)
}

// @Summary Get all products
// @Description Get a paginated list of products with optional filters
// @Tags products
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Param name query string false "Filter by product name (partial match)"
// @Param description query string false "Filter by product description (partial match)"
// @Param price_from query number false "Filter by minimum price"
// @Param price_to query number false "Filter by maximum price"
// @Param stock query int false "Filter by minimum stock"
// @Success 200 {object} shared.PaginatedResponse{data=[]Product} "OK with paginated products"
// @Failure 400 {object} shared.Response "Invalid query parameters"
// @Failure 404 {object} shared.Response "Products not found"
// @Failure 500 {object} shared.Response "Internal server error"
// @Router /products [get]
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
