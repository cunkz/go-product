package handlers

import (
	"fmt"

	wrapperHelper "github.com/cunkz/go-product/bin/helpers/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/cunkz/go-product/bin/config"
	"encoding/base64"
	"strings"
)

func AuthenticateJWT(c *fiber.Ctx) error {
	tokenString := strings.ReplaceAll(c.Get("Authorization"), "Bearer ", "")	
	publicKeyString, err := base64.StdEncoding.DecodeString(config.GetConfig().JwtPublicKey)
	if err != nil {
		return wrapperHelper.Response(c, "fail", nil, "Failed Get Public Key", 500)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKeyString))
	if err != nil {
		return wrapperHelper.Response(c, "fail", nil, "Failed Parse Public Key", 500)
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		return wrapperHelper.Response(c, "fail", nil, "Failed Parse Token", 401)
		// return fmt.Errorf("could not parse token: %w", err)
	}

	if !token.Valid {
		return wrapperHelper.Response(c, "fail", nil, "Invalid Token", 401)
		// return fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return wrapperHelper.Response(c, "fail", nil, "Failed Get Detail Token", 401)
		// return fmt.Errorf("could not parse claims")
	}

	c.Locals("claims", claims)
	return c.Next()
}
