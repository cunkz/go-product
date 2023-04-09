package utils

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2"
)

type ResponseModel struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    fiber.Map `json:"data"`
	Code    int       `json:"code"`
}

func Success(item fiber.Map) fiber.Map {
	return fiber.Map{
		"err":  nil,
		"data": item,
	}
}

func Response(c *fiber.Ctx, responseType string, responseData fiber.Map, responseMessage string, responseCode int) error {
	var success = true
	if responseType == "fail" {
		success = false
		responseData = fiber.Map{}
		if responseMessage == "" {
			responseMessage = fiber.ErrInternalServerError.Message
		}
	} else {
		responseCode = 200
		_, checkData := responseData["data"]
		if !checkData {
			responseCode = fiber.ErrNotFound.Code
			responseMessage = fiber.ErrNotFound.Message
			responseData = fiber.Map{}
		} else {
			responseData = responseData["data"].(fiber.Map)
		}
	}

	return c.Status(responseCode).JSON(ResponseModel{
		Success: success,
		Message: responseMessage,
		Data:    responseData,
		Code:    responseCode,
	})
}
