package main

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/database/migration"
	"go-fiber-gorm/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DatabaseInit()
	migration.RunMigration()
	app := fiber.New()

	// INITIAL ROUTE
	route.RouteInit(app)

	app.Listen(":3000")
}
