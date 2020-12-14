package project

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hackarmour/iota/common"
	"gorm.io/gorm"

	"strconv"
)

// Project model
type Project struct {
	gorm.Model
	ID          int    `json:"ID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetProjects gets all the projects
func GetProjects(c *fiber.Ctx) error {
	db := common.GetDB()
	var projects []Project
	db.Where(&Project{}).Find(&projects)
	return c.JSON(projects)
}

// PostProject posts a project
func PostProject(c *fiber.Ctx) error {
	db := common.GetDB()
	project := &Project{}
	if err := c.BodyParser(project); err != nil {
		return c.SendString("Unprocessable Entity")
	}
	db.Create(&project)
	if project.ID == 0 {
		return c.Status(404).SendString("Sorry what ?")
	}
	return c.JSON(project)
}

// GetOne project
func GetOne(c *fiber.Ctx) error {
	db := common.GetDB()
	project := &Project{}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendString("oops")
	}

	db.Where(&Project{ID: id}).Find(&project)
	if project.ID == 0 {
		return c.Status(404).SendString("NOT FOUND")
	}
	return c.JSON(project)
}
