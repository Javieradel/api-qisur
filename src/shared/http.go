package shared

import (
	"github.com/gofiber/fiber/v3"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
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

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewValidationErrorResponse(c fiber.Ctx, errors []ValidationError) error {
	return c.Status(fiber.StatusUnprocessableEntity).JSON(Response{
		Success: false,
		Data:    errors,
		Error:   "Validation failed",
	})
}
