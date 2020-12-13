package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"

	"github.com/hackarmour/iota/entity"
	"github.com/hackarmour/iota/entityvalues"
	"github.com/hackarmour/iota/project"
)

func routes(app *fiber.App) {
	app.Get("/", project.GetProjects)
}

func main() {
	app := fiber.New()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&project.Project{})
	db.AutoMigrate(&entity.Entity{})
	db.AutoMigrate(&entityvalues.EntityValues{})

	routes(app)
	log.Fatal(app.Listen(":4201"))
}
