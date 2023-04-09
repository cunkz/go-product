package handlers

import (
	"fmt"
	"time"

	postgresqlHelper "github.com/cunkz/go-product/bin/helpers/db/postgresql"
	wrapperHelper "github.com/cunkz/go-product/bin/helpers/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"database/sql"

	_ "github.com/lib/pq"
)

type Product struct {
	Label string `json:"label" xml:"label" form:"label"`
}

func AddProduct(c *fiber.Ctx) error {
	var db *sql.DB
	var err error
	db = postgresqlHelper.GetDB()

	fmt.Println("Successfully connected!")

	p := new(Product)
	if err := c.BodyParser(p); err != nil {
		return err
	}

	itemUUID := uuid.New()
	sqlStatement := `
	INSERT INTO products (productid, label, created_at, updated_at)
	VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(sqlStatement, itemUUID, p.Label, time.Now(), time.Now())
	if err != nil {
		panic(err)
	}

	item := fiber.Map{
		"productid": itemUUID,
	}

	result := wrapperHelper.Success(item)
	return wrapperHelper.Response(c, "success", result, "Product has been added", 200)
}

func EditProduct(c *fiber.Ctx) error {
	return c.SendString("Product has been edited #" + c.Params("id"))
}

func FetchProduct(c *fiber.Ctx) error {
	return c.SendString("Sucess fetch list of Product")
}

func FetchProductById(c *fiber.Ctx) error {
	var db *sql.DB
	db = postgresqlHelper.GetDB()
	sqlStatement := `SELECT label FROM products WHERE productid=$1;`
	var label string
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	row := db.QueryRow(sqlStatement, c.Params("id"))
	switch err := row.Scan(&label); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(label)
	default:
		panic(err)
	}

	item := fiber.Map{
		"label": label,
	}

	result := wrapperHelper.Success(item)
	return wrapperHelper.Response(c, "success", result, fmt.Sprintf("Success get a single Product #%s", c.Params("id")), 200)
}

func DeleteProduct(c *fiber.Ctx) error {
	return c.SendString("Product has been deleted #" + c.Params("id"))
}
