package controllers

import (
	"errors"
	"fmt"
	"log"

	"github.com/Aashish32/GORM-CRUD/database"
	"github.com/Aashish32/GORM-CRUD/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname" `
}

func CreateResponseUser(user models.User) User {
	return User{Id: user.Id, Firstname: user.Firstname, Lastname: user.Lastname}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}
func GetAllUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.Database.Db.Find(&users)

	//or just return the users as JSON. Everything will be same But created time will be exposed.

	responseUsers := []User{}
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)

	}

	return c.Status(200).JSON(responseUsers)

}
func GetUser(c *fiber.Ctx) error {
	user := User{}
	id, err := c.ParamsInt("id")
	if err != nil {
		errorMessage := fmt.Sprintf("Error fetching user with ID %d: %s", id, err)
		return c.Status(400).JSON("integer id parameter is not found", errorMessage)
	}

	database.Database.Db.Find(&user, "id=?", id)

	if user.Id == 0 {
		return errors.New(fmt.Sprintf("User with id %d does not exist", id))
	}
	return c.Status(200).JSON(user)

}

func UpdateUser(c *fiber.Ctx) error {
	user := models.User{}

	id, err := c.ParamsInt("id")
	if err != nil {
		errorMessage := fmt.Sprintf("Error fetching user with ID %d: %s", id, err)
		return c.Status(400).JSON("integer id parameter is not found", errorMessage)
	}
	database.Database.Db.Find(&user, "id=?", id)

	if user.Id == 0 {
		return errors.New(fmt.Sprintf("User with id %d does not exist", id))
	}

	if err := c.BodyParser(&user); err != nil {
		log.Println(err)
	}
	database.Database.Db.Save(&user)
	updatedUser := models.User{}
	database.Database.Db.First(&updatedUser, id)

	return c.Status(200).JSON(updatedUser)

}

func DeleteUser(c *fiber.Ctx) error {
	user := models.User{}

	id, err := c.ParamsInt("id")
	if err != nil {
		errorMessage := fmt.Sprintf("Error fetching user with ID %d: %s", id, err)
		return c.Status(400).JSON("integer id parameter is not found", errorMessage)
	}
	database.Database.Db.Find(&user, "id=?", id)

	if user.Id == 0 {
		return errors.New(fmt.Sprintf("User with id %d does not exist", id))
	}
	database.Database.Db.Delete(&user)
	message := fmt.Sprintf("User succesfully deleted : %v", user)
	return c.Status(200).JSON(message)

}
