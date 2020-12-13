package project

import (
	"github.com/gofiber/fiber/v2"
)

func GetProjects(c *fiber.Ctx) error {
	return c.SendString("Get the projects")
}
