package models

import (
	"github.com/google/uuid"
)

// App represents the application a workflow step runs.
type App struct {
	Base
	Name                string
	PublicID            uuid.UUID `gorm:"unique; type:uuid"`
	IsSecurityReviewed  bool
	IsOnAppEntityBehalf bool
	CreatorID           uuid.UUID
	AuthInfoID          uuid.UUID
	Account             []Account
	XRole               []*Role `gorm:"many2many:apps_x_roles"`
}
