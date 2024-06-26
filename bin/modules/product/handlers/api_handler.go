package handlers

import (
	"fmt"
	"time"
	"log"

	postgresqlHelper "github.com/cunkz/go-product/bin/helpers/db/postgresql"
	wrapperHelper "github.com/cunkz/go-product/bin/helpers/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type Product struct {
	ProductId string `db:"productid" json:"productid" xml:"productid" form:"productid"`
	Label string `db:"label" json:"label" xml:"label" form:"label"`
}

func AddProduct(c *fiber.Ctx) error {
	var err error
	var db *sqlx.DB
	db = postgresqlHelper.GetDB()

	fmt.Println("Successfully connected!")

	p := new(Product)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	p.ProductId = uuid.New().String()

	sqlStatement := `
	INSERT INTO products (productid, label, created_at, updated_at)
	VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(sqlStatement, p.ProductId, p.Label, time.Now(), time.Now())
	if err != nil {
		panic(err)
	}

	item := fiber.Map{
		"productid": p.ProductId,
	}

	result := wrapperHelper.Success(item)
	return wrapperHelper.Response(c, "success", result, "Product has been added", 201)
}

func EditProduct(c *fiber.Ctx) error {
	var err error
	var db *sqlx.DB
	db = postgresqlHelper.GetDB()

	product := Product{}
	sqlStatement := `SELECT productid, label FROM products WHERE productid=$1 LIMIT 1`
  rows, _ := db.Queryx(sqlStatement, c.Params("id"))
	for rows.Next() {
		err := rows.StructScan(&product)
		if err != nil {
				log.Fatalln(err)
		}
	}
	if product.ProductId == "" {
		return wrapperHelper.Response(c, "fail", nil, "Product not found", 404)
	}

	p := new(Product)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	sqlStatementUpdate := `
	UPDATE products 
	SET label = $2, updated_at = $3
	WHERE productid = $1`
	_, err = db.Exec(sqlStatementUpdate, c.Params("id"), p.Label, time.Now())
	if err != nil {
		panic(err)
	}
	
	return wrapperHelper.Response(c, "default", nil, "Product #" + c.Params("id") + " has been edited", 200)
}

func FetchProduct(c *fiber.Ctx) error {
	var db *sqlx.DB
	db = postgresqlHelper.GetDB()

	var products []Product
	product := Product{}
	sqlStatement := `SELECT productid, label FROM products`
  rows, _ := db.Queryx(sqlStatement)
	for rows.Next() {
		err := rows.StructScan(&product)
		if err != nil {
				log.Fatalln(err)
		}
		products = append(products, product)
	}

	items := fiber.Map{
		"items": products,
	}

	result := wrapperHelper.Success(items)
	return wrapperHelper.Response(c, "success", result, "Sucess fetch list of Product", 200)
}

func FetchProductById(c *fiber.Ctx) error {
	var db *sqlx.DB
	db = postgresqlHelper.GetDB()
	// var label string

	product := Product{}
	sqlStatement := `SELECT productid, label FROM products WHERE productid=$1 LIMIT 1`
  rows, _ := db.Queryx(sqlStatement, c.Params("id"))
	for rows.Next() {
		err := rows.StructScan(&product)
		if err != nil {
				log.Fatalln(err)
		}
	}
	if product.ProductId == "" {
		return wrapperHelper.Response(c, "fail", nil, "Product not found", 404)
	}

	item := fiber.Map{
		"productid": product.ProductId,
		"label": product.Label,
	}

	result := wrapperHelper.Success(item)
	return wrapperHelper.Response(c, "success", result, "Success get detail of Product #" + c.Params("id"), 200)
}

func RemoveProduct(c *fiber.Ctx) error {
	var db *sqlx.DB
	db = postgresqlHelper.GetDB()

	product := Product{}
	sqlStatement := `SELECT productid, label FROM products WHERE productid=$1 LIMIT 1`
  rows, _ := db.Queryx(sqlStatement, c.Params("id"))
	for rows.Next() {
		err := rows.StructScan(&product)
		if err != nil {
				log.Fatalln(err)
		}
	}
	if product.ProductId == "" {
		return wrapperHelper.Response(c, "fail", nil, "Product not found", 404)
	}

	_, err := db.Exec("DELETE FROM products WHERE productid=$1", product.ProductId)
	if err != nil {
		log.Fatalln(err)
	}
	return wrapperHelper.Response(c, "default", nil, "Product #" + c.Params("id") + " has been removed", 200)
}
