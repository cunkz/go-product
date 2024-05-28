package handlers

import (
	"time"

	wrapperHelper "github.com/cunkz/go-product/bin/helpers/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	item := fiber.Map{
		"accessToken": "sampleaccesstoken",
		"accessTokenExpiresAt": time.Now(),
		"refreshToken": "samplerefreshtoken",
		"refreshTokenExpiresAt": time.Now(),
		"user": map[string]string{ "userId": "sampleuser"	},
		"client": map[string]string{ "clientId": "sampleclient"	},
	}

	result := wrapperHelper.Success(item)
	return wrapperHelper.Response(c, "success", result, "Successfully Logged", 201)
}