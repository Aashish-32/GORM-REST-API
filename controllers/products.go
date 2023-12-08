package controllers

import (
	"strconv"

	"github.com/Aashish32/GORM-CRUD/database"
	"github.com/Aashish32/GORM-CRUD/models"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID           uint   `json:"id" `
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(product models.Product) Product {
	return Product{ID: product.Id, Name: product.Name, SerialNumber: product.SerialNumber}

}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		c.Status(400).JSON(err)
	}

	database.Database.Db.Create(&product)
	ResponseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(ResponseProduct)
}

func GetAllProducts(c *fiber.Ctx) error {

	products := []models.Product{}
	database.Database.Db.Find(&products)

	ResponseProducts := []Product{}

	for _, product := range products {
		ResponseProduct := CreateResponseProduct(product)
		ResponseProducts = append(ResponseProducts, ResponseProduct)

	}

	return c.Status(400).JSON(ResponseProducts)

}

func GetProduct(c *fiber.Ctx) error {
	var product Product
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	database.Database.Db.Find(&product, "id=?", id)
	if product.ID == 0 {
		return c.Status(400).JSON("user not found")
	}
	return c.Status(200).JSON(product)

}
func DeleteProduct(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	idstring := strconv.Itoa(id)
	if err != nil {
		return c.Status(400).JSON(err)
	}
	database.Database.Db.Delete(&models.Product{})
	return c.JSON("Deleted id:", idstring)
}
