package entity

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hackarmour/iota/common"
	"gorm.io/gorm"
)

// Entity model
type Entity struct {
	gorm.Model
	ID        int
	Key       string `json:"key"`
	Value     string `json:"value"`
	Note      string `json:"note"`
	ProjectID int    `json:"projectID"`
}

// CreateEntity with projectid
func CreateEntity(c *fiber.Ctx) error {
	db := common.GetDB()
	entity := &Entity{}
	if err := c.BodyParser(entity); err != nil {
		return c.JSON(err)
	}
	entity.ID = 0
	db.Create(&entity)
	return c.JSON(entity)
}

// DeleteEntity deletes
func DeleteEntity(c *fiber.Ctx) error {
	db := common.GetDB()
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendString("oops")
	}
	db.Delete(&Entity{ID: id})
	return c.JSON(&fiber.Map{
		"message": "Deleted lol",
	})
}
