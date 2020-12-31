package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

// Workflow represents a workflow.
type Workflow struct {
	Base
	Name      string
	Code      datatypes.JSON
	IsActive  bool
	IsDraft   bool
	FolderID  uuid.UUID
	CreatorID uuid.UUID
	XRole     []*Role `gorm:"many2many:workflow_x_roles"`
}
