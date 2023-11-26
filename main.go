package main

import (
	"log"

	"github.com/Aashish32/GORM-CRUD/controllers"
	"github.com/Aashish32/GORM-CRUD/database"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connectdb()
	app := fiber.New()

	SetupRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}

}

func welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("welcome to the api....")

}

func SetupRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users", controllers.GetAllUsers)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/update/:id", controllers.UpdateUser)

}
