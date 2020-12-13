package entity

import (
	"gorm.io/gorm"
)

type Entity struct {
	gorm.Model
	Name      string
	ProjectID string
	ID        int
}
