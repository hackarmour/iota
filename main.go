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
	app.Get("/", hello)

	app.Get("/projects", project.GetProjects)
	app.Post("/projects", project.PostProject)

}

// Migrate all the models
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

func hello(c *fiber.Ctx) error {
	return c.SendString("welcome to the iota API server\n")
}
