package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

// Theme represents a theme used to change the look of the application.
type Theme struct {
	Base
	Name      string
	Code      datatypes.JSON
	PublicID  uuid.UUID `gorm:"unique; type:uuid"`
	CreatorID uuid.UUID
	XRole     []*Role `gorm:"many2many:themes_x_roles"`
}
