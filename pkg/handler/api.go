package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type APIError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type APISuccess struct {
	Data interface{} `json:"data"`
}

type APIResponse struct {
	Type   string      `json:"type"`
	Error  *APIError   `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Status int         `json:"status"`
}

func CreateErrorResponse(message string, code int) APIResponse {
	return APIResponse{
		Type: "error",
		Error: &APIError{
			Message: message,
			Code:    code,
		},
		Status: code,
	}
}

func HandleAPIError(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	// Check if it's a fiber.Error type
	var e *fiber.Error
	if errors.As(err, &e) {
		// Override status code if fiber.Error type
		code = e.Code
		message = e.Message
	}

	response := c.Status(code).JSON(CreateErrorResponse(message, code))
	return response
}

func CreateSuccessResponse(data interface{}) APIResponse {
	return APIResponse{
		Type:   "success",
		Data:   data,
		Status: fiber.StatusOK,
	}
}

func HandleAPISuccess(c *fiber.Ctx, data interface{}) error {
	response := c.Status(fiber.StatusOK).JSON(CreateSuccessResponse(data))
	return response
}
