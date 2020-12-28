package main

import (
	"log"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"

	"github.com/hackarmour/iota/common"

	"github.com/hackarmour/iota/entity"
	"github.com/hackarmour/iota/project"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&project.Project{})
	db.AutoMigrate(&entity.Entity{})
}

func main() {
	app := fiber.New()
	db := common.Init()
	migrate(db)

	routes(app)
	log.Fatal(app.Listen(":4201"))
}

func routes(app *fiber.App) {
	app.Get("/", hello)

	app.Get("/projects", project.GetProjects)
	app.Post("/projects", project.PostProject)
	app.Get("/projects/:id", project.GetOne)
	app.Delete("/projects/:id", project.DeleteProject)
	app.Patch("/projects/:id", project.UpdateProject)

	app.Post("/entity", entity.CreateEntity)
	app.Delete("/entity/:id", entity.DeleteEntity)
}

func hello(c *fiber.Ctx) error {
	return c.SendString("welcome to the iota API server\n")
}
