package main

import (
	"tayoh-backend/database"
	"tayoh-backend/database/migration"
	"tayoh-backend/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DatabaseInit()
	migration.RunMigration()
	
	app := fiber.New()

	route.RouteInit(app)

	app.Listen(":8000")
}