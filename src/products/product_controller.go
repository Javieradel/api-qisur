package products

import (
	"errors"
	"strconv"

	"github.com/Javieradel/api-qisur.git/src/shared"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type ProductController struct {
	service   *ProductService
	validator *shared.XValidator
}

func NewProductController(service *ProductService) *ProductController {
	return &ProductController{
		service:   service,
		validator: shared.NewValidator(),
	}
}

func (pc *ProductController) RegisterRoutes(app *fiber.App) {
	app.Get("/api/v1/products", pc.GetProducts)
	app.Get("/api/v1/products/:id", pc.GetProductByID)
	app.Post("/api/v1/products", pc.CreateProduct)
	app.Put("/api/v1/products/:id", pc.UpdateProduct)
	app.Patch("/api/v1/products/:id", pc.PatchProduct)
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
	if err := c.Bind().Query(&q); err != nil {
		return shared.NewErrorResponse(c, fiber.StatusBadRequest, "Invalids query params")
	}

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

// @Summary Get product by ID
// @Description Get a single product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} shared.Response{data=Product} "OK with product data"
// @Failure 400 {object} shared.Response "Invalid product ID"
// @Failure 404 {object} shared.Response "Product not found"
// @Failure 500 {object} shared.Response "Internal server error"
// @Router /products/{id} [get]
func (pc *ProductController) GetProductByID(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return shared.NewErrorResponse(c, fiber.StatusBadRequest, "Invalid product ID")
	}

	product, err := pc.service.FindByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return shared.NewErrorResponse(c, fiber.StatusNotFound, "Product not found")
		}
		return shared.NewErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch product")
	}

	return shared.NewSuccessResponse(c, fiber.StatusOK, product)
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with the given data
// @Tags products
// @Accept json
// @Produce json
// @Param product body CreateProductDTO true "Product data"
// @Success 201 {object} shared.Response{data=Product} "Product created successfully"
// @Failure 400 {object} shared.Response "Invalid request body"
// @Failure 422 {object} shared.Response "Validation failed"
// @Failure 500 {object} shared.Response "Internal server error"
// @Router /products [post]
func (pc *ProductController) CreateProduct(c fiber.Ctx) error {
	var dto CreateProductDTO
	if err := c.Bind().Body(&dto); err != nil {
		return shared.NewErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if errs := pc.validator.Validate(dto); len(errs) > 0 {
		return shared.NewValidationErrorResponse(c, errs)
	}

	product := dto.ToProduct()
	if err := pc.service.Create(product); err != nil {
		return shared.NewErrorResponse(c, fiber.StatusInternalServerError, "Failed to create product")
	}
	if len(dto.CategoriesID) > 0 {
		if err := pc.service.UpdateCategories(product, dto.CategoriesID); err != nil {
			return shared.NewErrorResponse(c, fiber.StatusInternalServerError, "Failed to set categories for product")
		}
	}
	return shared.NewSuccessResponse(c, fiber.StatusCreated, product)
}

// @Summary Update a product
// @Description Update an existing product with the given data
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body UpdateProductDTO true "Product data"
// @Success 200 {object} shared.Response{data=Product} "Product updated successfully"
// @Failure 400 {object} shared.Response "Invalid request body or product ID"
// @Failure 404 {object} shared.Response "Product not found"
// @Failure 422 {object} shared.Response "Validation failed"
// @Failure 500 {object} shared.Response "Internal server error"
// @Router /products/{id} [put]
func (pc *ProductController) UpdateProduct(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return shared.NewErrorResponse(c, fiber.StatusBadRequest, "Invalid product ID")
	}

	var dto UpdateProductDTO
	if err := c.Bind().Body(&dto); err != nil {
		return shared.NewErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if errs := pc.validator.Validate(dto); len(errs) > 0 {
		return shared.NewValidationErrorResponse(c, errs)
	}

	product, err := pc.service.FindByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return shared.NewErrorResponse(c, fiber.StatusNotFound, "Product not found")
		}
		return shared.NewErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch product")
	}

	product.Name = dto.Name
	product.Description = dto.Description
	product.Price = dto.Price
	product.Stock = dto.Stock

	if err := pc.service.Update(product); err != nil {
		return shared.NewErrorResponse(c, fiber.StatusInternalServerError, "Failed to update product")
	}
	if err := pc.service.UpdateCategories(product, dto.CategoriesID); err != nil {
		return shared.NewErrorResponse(c, fiber.StatusInternalServerError, "Failed to update categories for product")
	}
	return shared.NewSuccessResponse(c, fiber.StatusOK, product)
}

// PatchProduct godoc
// @Summary Partially update a product
// @Description Partially update an existing product with the given data
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body PatchProductDTO true "Product data"
// @Success 200 {object} shared.Response{data=Product} "Product updated successfully"
// @Failure 400 {object} shared.Response "Invalid request body or product ID"
// @Failure 404 {object} shared.Response "Product not found"
// @Failure 422 {object} shared.Response "Validation failed"
// @Failure 500 {object} shared.Response "Internal server error"
// @Router /products/{id} [patch]
func (pc *ProductController) PatchProduct(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return shared.NewErrorResponse(c, fiber.StatusBadRequest, "Invalid product ID")
	}

	var dto PatchProductDTO
	if err := c.Bind().Body(&dto); err != nil {
		return shared.NewErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if errs := pc.validator.Validate(dto); len(errs) > 0 {
		return shared.NewValidationErrorResponse(c, errs)
	}

	product, err := pc.service.FindByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return shared.NewErrorResponse(c, fiber.StatusNotFound, "Product not found")
		}
		return shared.NewErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch product")
	}

	if dto.Name != nil {
		product.Name = *dto.Name
	}
	if dto.Description != nil {
		product.Description = *dto.Description
	}
	if dto.Price != nil {
		product.Price = *dto.Price
	}
	if dto.Stock != nil {
		product.Stock = *dto.Stock
	}

	if err := pc.service.Update(product); err != nil {
		return shared.NewErrorResponse(c, fiber.StatusInternalServerError, "Failed to update product")
	}

	if dto.CategoriesID != nil {
		if err := pc.service.UpdateCategories(product, *dto.CategoriesID); err != nil {
			return shared.NewErrorResponse(c, fiber.StatusInternalServerError, "Failed to update categories for product")
		}
	}

	return shared.NewSuccessResponse(c, fiber.StatusOK, product)
}
