package project

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hackarmour/iota/common"
	"github.com/hackarmour/iota/entity"
	"gorm.io/gorm"

	"strconv"
	"sync"
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
	project.ID = 0
	db.Create(&project)
	if project.ID == 0 {
		return c.Status(404).SendString("Sorry what ?")
	}
	return c.JSON(project)
}

// GetOne project with all its entities
func GetOne(c *fiber.Ctx) error {
	db := common.GetDB()
	project := &Project{}
	var entities []entity.Entity
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendString("oops")
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		db.Where(&Project{ID: id}).Find(&project)
		wg.Done()
	}()

	go func() {
		db.Where(&entity.Entity{ProjectID: id}).Find(&entities)
		wg.Done()
	}()

	wg.Wait()
	if project.ID == 0 {
		return c.Status(404).SendString("NOT FOUND")
	}
	return c.JSON(&fiber.Map{
		"project":  project,
		"entities": entities,
	})
}

// UpdateProject updates
func UpdateProject(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"message": "updated",
	})
}

// DeleteProject goes brrrr
func DeleteProject(c *fiber.Ctx) error {
	db := common.GetDB()
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendString("oops")
	}
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		db.Delete(&Project{ID: id})
		wg.Done()
	}()
	go func() {
		var entities []entity.Entity
		db.Where(&entity.Entity{ProjectID: id}).Delete(&entities)
		wg.Done()
	}()
	wg.Wait()
	return c.JSON(&fiber.Map{
		"message": "deleted lol",
	})
}
