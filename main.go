package main

import (
	"tayoh-backend//migration"
	"tayoh-backend//route"
	"tayoh-backend/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DatabaseInit()
	migration.RunMigration()
	
	app := fiber.New()

	route.RouteInit(app)

	app.Listen(":8000")
}