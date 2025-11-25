package categories

import (
	"errors"
	"strconv"

	"github.com/Javieradel/api-qisur.git/src/shared"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type CategoryController struct {
	service   *CategoryService
	validator *shared.XValidator
}

func NewCategoryController(service *CategoryService) *CategoryController {
	return &CategoryController{
		service:   service,
		validator: shared.NewValidator(),
	}
}

func (cc *CategoryController) RegisterRoutes(app *fiber.App) {
	app.Get("/api/v1/categories", cc.GetCategories)
	app.Get("/api/v1/categories/:id", cc.GetCategoryByID)
	app.Post("/api/v1/categories", cc.CreateCategory)
	app.Put("/api/v1/categories/:id", cc.UpdateCategory)
	app.Patch("/api/v1/categories/:id", cc.PatchCategory)
	app.Delete("/api/v1/categories/:id", cc.DeleteCategory)
}

// @Summary Get all categories
// @Description Get a paginated list of categories with optional filters
// @Tags categories
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Param name query string false "Filter by category name (partial match)"
// @Param description query string false "Filter by category description (partial match)"
// @Success 200 {object} shared.PaginatedResponse{data=[]Categories} "OK with paginated categories"
// @Failure 400 {object} shared.Response "Invalid query parameters"
// @Failure 404 {object} shared.Response "Categories not found"
// @Failure 500 {object} shared.Response "Internal server error"
// @Router /categories [get]
func (cc *CategoryController) GetCategories(c fiber.Ctx) error {
	var q CategoryQueryDTO
	if err := c.Bind().Query(&q); err != nil {
		return shared.NewErrorResponse(c, fiber.StatusBadRequest, "Invalids query params")
	}

	filters := q.ToCriterions()
	categories, err := cc.service.FindAll(filters)
	if err != nil {
		return shared.NewErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch categories")
	}

	if len(categories) == 0 {
		return shared.NewErrorResponse(c, fiber.StatusNotFound, "Categories not found")
	}

	if q.Limit > 0 || q.Page > 0 {
		return shared.NewPaginatedResponse(c, fiber.StatusFound, categories, q.Page, q.Limit)
	}

	return shared.NewSuccessResponse(c, fiber.StatusFound, categories)
}

// @Summary Get category by ID
// @Description Get a single category by its ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} shared.Response{data=Categories} "OK with category data"
// @Failure 400 {object} shared.Response "Invalid category ID"
// @Failure 404 {object} shared.Response "Category not found"
// @Failure 500 {object} shared.Response "Internal server error"
// @Router /categories/{id} [get]
func (cc *CategoryController) GetCategoryByID(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return shared.NewErrorResponse(c, fiber.StatusBadRequest, "Invalid category ID")
	}

	category, err := cc.service.FindByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return shared.NewErrorResponse(c, fiber.StatusNotFound, "Category not found")
		}
		return shared.NewErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch category")
	}

	return shared.NewSuccessResponse(c, fiber.StatusOK, category)
}

// @Summary Create a new category
// @Description Create a new category with the given data
// @Tags categories
// @Accept json
// @Produce json
// @Param category body CreateCategoryDTO true "Category data"
// @Success 201 {object} shared.Response{data=Categories} "Category created successfully"
// @Failure 400 {object} shared.Response "Invalid request body"
// @Failure 422 {object} shared.Response "Validation failed"
// @Failure 500 {object} shared.Response "Internal server error"
// @Router /categories [post]
func (cc *CategoryController) CreateCategory(c fiber.Ctx) error {
	var dto CreateCategoryDTO
	if err := c.Bind().Body(&dto); err != nil {
		return shared.NewErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if errs := cc.validator.Validate(dto); len(errs) > 0 {
		return shared.NewValidationErrorResponse(c, errs)
	}

	category := dto.ToCategory()
	if err := cc.service.Create(category); err != nil {
		return shared.NewErrorResponse(c, fiber.StatusInternalServerError, "Failed to create category")
	}

	return shared.NewSuccessResponse(c, fiber.StatusCreated, category)
}

// @Summary Update a category
// @Description Update an existing category with the given data
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body UpdateCategoryDTO true "Category data"
// @Success 200 {object} shared.Response{data=Categories} "Category updated successfully"
// @Failure 400 {object} shared.Response "Invalid request body or category ID"
// @Failure 404 {object} shared.Response "Category not found"
// @Failure 422 {object} shared.Response "Validation failed"
// @Failure 500 {object} shared.Response "Internal server error"
// @Router /categories/{id} [put]
func (cc *CategoryController) UpdateCategory(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return shared.NewErrorResponse(c, fiber.StatusBadRequest, "Invalid category ID")
	}

	var dto UpdateCategoryDTO
	if err := c.Bind().Body(&dto); err != nil {
		return shared.NewErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if errs := cc.validator.Validate(dto); len(errs) > 0 {
		return shared.NewValidationErrorResponse(c, errs)
	}

	category, err := cc.service.FindByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return shared.NewErrorResponse(c, fiber.StatusNotFound, "Category not found")
		}
		return shared.NewErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch category")
	}

	category.Name = dto.Name
	category.Description = dto.Description

	if err := cc.service.Update(category); err != nil {
		return shared.NewErrorResponse(c, fiber.StatusInternalServerError, "Failed to update category")
	}

	return shared.NewSuccessResponse(c, fiber.StatusOK, category)
}

// @Summary Partially update a category
// @Description Partially update an existing category with the given data
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body PatchCategoryDTO true "Category data"
// @Success 200 {object} shared.Response{data=Categories} "Category updated successfully"
// @Failure 400 {object} shared.Response "Invalid request body or category ID"
// @Failure 404 {object} shared.Response "Category not found"
// @Failure 422 {object} shared.Response "Validation failed"
// @Failure 500 {object} shared.Response "Internal server error"
// @Router /categories/{id} [patch]
func (cc *CategoryController) PatchCategory(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return shared.NewErrorResponse(c, fiber.StatusBadRequest, "Invalid category ID")
	}

	var dto PatchCategoryDTO
	if err := c.Bind().Body(&dto); err != nil {
		return shared.NewErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if errs := cc.validator.Validate(dto); len(errs) > 0 {
		return shared.NewValidationErrorResponse(c, errs)
	}

	category, err := cc.service.FindByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return shared.NewErrorResponse(c, fiber.StatusNotFound, "Category not found")
		}
		return shared.NewErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch category")
	}

	if dto.Name != nil {
		category.Name = *dto.Name
	}
	if dto.Description != nil {
		category.Description = *dto.Description
	}

	if err := cc.service.Update(category); err != nil {
		return shared.NewErrorResponse(c, fiber.StatusInternalServerError, "Failed to update category")
	}

	return shared.NewSuccessResponse(c, fiber.StatusOK, category)
}

// @Summary Delete a category
// @Description Delete an existing category by its ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} shared.Response "Category deleted successfully"
// @Failure 400 {object} shared.Response "Invalid category ID"
// @Failure 404 {object} shared.Response "Category not found"
// @Failure 500 {object} shared.Response "Internal server error"
// @Router /categories/{id} [delete]
func (cc *CategoryController) DeleteCategory(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return shared.NewErrorResponse(c, fiber.StatusBadRequest, "Invalid category ID")
	}

	_, err = cc.service.FindByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return shared.NewErrorResponse(c, fiber.StatusNotFound, "Category not found")
		}
		return shared.NewErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch category")
	}

	if err := cc.service.Delete(uint(id)); err != nil {
		return shared.NewErrorResponse(c, fiber.StatusInternalServerError, "Failed to delete category")
	}

	return shared.NewSuccessResponse(c, fiber.StatusOK, "Category deleted successfully")
}