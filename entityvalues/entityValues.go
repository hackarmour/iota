package entityvalues

import (
	"gorm.io/gorm"
)

type EntityValues struct {
	gorm.Model
	Name     string
	ID       int
	EntityID int
}
