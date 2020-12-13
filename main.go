package main

import (
	"log"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"

	"github.com/hackarmour/iota/common"

	"github.com/hackarmour/iota/entity"
	"github.com/hackarmour/iota/entityvalues"
	"github.com/hackarmour/iota/project"
)

func routes(app *fiber.App) {
	app.Get("/", project.GetProjects)
	app.Get("/foo", project.CreateDummyProject)
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&project.Project{})
	db.AutoMigrate(&entity.Entity{})
	db.AutoMigrate(&entityvalues.EntityValues{})
}

func main() {
	app := fiber.New()
	db := common.Init()
	Migrate(db)

	routes(app)
	log.Fatal(app.Listen(":4201"))
}
