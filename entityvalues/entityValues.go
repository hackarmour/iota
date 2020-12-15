package entityvalues

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hackarmour/iota/common"
	"gorm.io/gorm"
)

// EntityValue model
type EntityValue struct {
	gorm.Model
	Name     string `json:"name"`
	ID       int
	EntityID int `json:"entityID"`
}

// PostEntityValue stuff
func PostEntityValue(c *fiber.Ctx) error {
	db := common.GetDB()
	entityValue := &EntityValue{}
	if err := c.BodyParser(entityValue); err != nil {
		return c.SendString("oops")
	}
	entityValue.ID = 0
	db.Create(&entityValue)
	return c.JSON(entityValue)
}
