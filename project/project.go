package project

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hackarmour/iota/common"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name        string
	Description string
}

func GetProjects(c *fiber.Ctx) error {
	db := common.GetDB()
	var projects []Project
	db.Where(&Project{}).Find(&projects)
	return c.JSON(projects)
}

func CreateDummyProject(c *fiber.Ctx) error {
	db := common.GetDB()
	db.Create(&Project{Name: "Project One", Description: "woooooo"})
	return c.SendString("wooo")
}
