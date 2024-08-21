package handlers

import (
	// "fmt"
	"time"
	"encoding/base64"

	wrapperHelper "github.com/cunkz/go-product/bin/helpers/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/cunkz/go-product/bin/config"
)

var jwtIssuer = config.GetConfig().JwtIssuer
var jwtAudience = config.GetConfig().JwtAudience

func Login(c *fiber.Ctx) error {
	privateKeyString, err := base64.StdEncoding.DecodeString(config.GetConfig().JwtPrivateKey)
	if err != nil {
		return err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKeyString))
	if err != nil {
		return err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, 
		jwt.MapClaims{
			"user": map[string]string{ "userId": "sampleuser"	},
			"client": map[string]string{ "clientId": "sampleclient"	},
			"iat": time.Now().Unix(),
			"exp": time.Now().Add(time.Hour * 24).Unix(), 
			"aud": jwtAudience,
			"iss": jwtIssuer,
		})

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return err
	}

	item := fiber.Map{
		"accessToken": tokenString,
		"accessTokenExpiresAt": time.Now(),
		"refreshToken": tokenString,
		"refreshTokenExpiresAt": time.Now(),
		"user": map[string]string{ "userId": "sampleuser"	},
		"client": map[string]string{ "clientId": "sampleclient"	},
	}

	result := wrapperHelper.Success(item)
	return wrapperHelper.Response(c, "success", result, "Successfully Logged", 201)
}

func GetMe(c *fiber.Ctx) error {
	claims := c.Locals("claims").(*jwt.MapClaims)

	item := fiber.Map{
		"client": (*claims)["client"],
		"user": (*claims)["user"],
	}

	result := wrapperHelper.Success(item)
	return wrapperHelper.Response(c, "success", result, "Successfully Get Me", 200)
}
