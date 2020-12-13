package project

import (
	"github.com/gofiber/fiber/v2"
)

type Project struct {
	Name        string
	Description string
	ID          int
}

func GetProjects(c *fiber.Ctx) error {
	return c.SendString("Get the projects")
}
