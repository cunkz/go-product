package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func AddProduct(c *fiber.Ctx) error {
	return c.SendString("Product has been added")
}

func EditProduct(c *fiber.Ctx) error {
	return c.SendString("Product has been edited #" + c.Params("id"))
}

func FetchProduct(c *fiber.Ctx) error {
	return c.SendString("Sucess fetch list of Product")
}

func FetchProductById(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf("Success get a single Product #%s", c.Params("id")))
}

func DeleteProduct(c *fiber.Ctx) error {
	return c.SendString("Product has been deleted #" + c.Params("id"))
}
