package shared

import (
	"github.com/gofiber/fiber/v3"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type PaginatedResponse struct {
	Response
	Page  int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
	//TotalItems int64 `json:"totalItems,omitempty"`
	//TotalPages int   `json:"totalPages,omitempty"`
}

func NewSuccessResponse(c fiber.Ctx, status int, data interface{}) error {
	return c.Status(status).JSON(Response{
		Success: true,
		Data:    data,
	})
}

func NewErrorResponse(c fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(Response{
		Success: false,
		Error:   message,
	})
}

type ErrorResponse struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
}

func NewValidationErrorResponse(c fiber.Ctx, errors []ErrorResponse) error {
	return c.Status(fiber.StatusUnprocessableEntity).JSON(Response{
		Success: false,
		Data:    errors,
		Error:   "Validation failed",
	})
}

func NewPaginatedResponse(c fiber.Ctx, status int, data interface{}, page, limit int) error {
	return c.Status(status).JSON(PaginatedResponse{
		Response: Response{
			Success: true,
			Data:    data,
		},
		Page:  page,
		Limit: limit,
		//TotalItems: totalItems,
		//TotalPages: totalPages,
	})
}
