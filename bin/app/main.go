package main

import (
	"fmt"

	"github.com/cunkz/go-product/bin/config"
	postgresqlHelper "github.com/cunkz/go-product/bin/helpers/db/postgresql"
	productHandler "github.com/cunkz/go-product/bin/modules/product/handlers"
	wrapperHelper "github.com/cunkz/go-product/bin/helpers/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return wrapperHelper.Response(c, "default", nil, "This service is running properly", 200)
	})

	app.Post("/api/product/v1", productHandler.AddProduct)
	app.Get("/api/product/v1", productHandler.FetchProduct)
	app.Get("/api/product/v1/:id", productHandler.FetchProductById)
	app.Put("/api/product/v1/:id", productHandler.EditProduct)
	app.Delete("/api/product/v1/:id", productHandler.DeleteProduct)

	// ----- init section -----
	postgresqlHelper.InitConnection()

	listenerPort := fmt.Sprintf("127.0.0.1:%s", config.GetConfig().AppPort)
	app.Listen(listenerPort)
}
