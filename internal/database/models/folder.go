package models

import (
	"github.com/google/uuid"
)

// Folder represents the folder workflows and the likes can be kept.
type Folder struct {
	Base
	Name        string
	Description string
	Avatar32URL string `gorm:"column:avatar_32_url"`
	CreatorID   uuid.UUID
	Workflow    []Workflow
}
