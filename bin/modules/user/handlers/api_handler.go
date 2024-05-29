package handlers

import (
	"time"

	wrapperHelper "github.com/cunkz/go-product/bin/helpers/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey = []byte("jwt-secret-key")
var jwtIssuer = "jwt-issuer"
var jwtAudience = "jwt-audience"

func Login(c *fiber.Ctx) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
		jwt.MapClaims{
			"user": map[string]string{ "userId": "sampleuser"	},
			"client": map[string]string{ "clientId": "sampleclient"	},
			"iat": time.Now().Unix(),
			"exp": time.Now().Add(time.Hour * 24).Unix(), 
			"aud": jwtAudience,
			"iss": jwtIssuer,
		})

	tokenString, err := token.SignedString(jwtSecretKey)
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