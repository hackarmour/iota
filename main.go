package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/hackarmour/iota/project"
)

func routes(app *fiber.App) {
	app.Get("/", project.GetProjects)
}

func main() {
	app := fiber.New()
	routes(app)
	log.Fatal(app.Listen(":4201"))
}
