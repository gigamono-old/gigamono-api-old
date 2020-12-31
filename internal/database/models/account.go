package models

import (
	"github.com/google/uuid"
)

// Account represents credentials needed for an app to authorize access.
type Account struct {
	Base
	App               App
	AccessTokenCredID uuid.UUID `gorm:"unique; type:uuid"`
	AppID             uuid.UUID
	UserID            uuid.UUID
}
