package project

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hackarmour/iota/common"
	"gorm.io/gorm"
)

// Project model
type Project struct {
	gorm.Model
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
	return c.JSON(project)
}
