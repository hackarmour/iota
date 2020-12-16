package entity

import (
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
