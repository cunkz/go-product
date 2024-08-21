package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/cunkz/go-product/bin/config"
	postgresqlHelper "github.com/cunkz/go-product/bin/helpers/db/postgresql"
	wrapperHelper "github.com/cunkz/go-product/bin/helpers/utils"

	auth "github.com/cunkz/go-product/bin/auth"

	productHandler "github.com/cunkz/go-product/bin/modules/product/handlers"
	userHandler "github.com/cunkz/go-product/bin/modules/user/handlers"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return wrapperHelper.Response(c, "default", nil, "This service is running properly", 200)
	})

	// ----- user route -----
	app.Post("/api/user/v1/login", userHandler.Login)
	app.Get("/api/user/v1/me", auth.AuthenticateJWT, userHandler.GetMe)

	// ----- product route -----
	app.Post("/api/product/v1", auth.AuthenticateJWT, productHandler.AddProduct)
	app.Get("/api/product/v1", auth.AuthenticateJWT, productHandler.FetchProduct)
	app.Get("/api/product/v1/:id", auth.AuthenticateJWT, productHandler.FetchProductById)
	app.Put("/api/product/v1/:id", auth.AuthenticateJWT, productHandler.EditProduct)
	app.Delete("/api/product/v1/:id", auth.AuthenticateJWT, productHandler.RemoveProduct)

	// ----- init section -----
	postgresqlHelper.InitConnection()

	listenerPort := fmt.Sprintf("127.0.0.1:%s", config.GetConfig().AppPort)
	app.Listen(listenerPort)
}
