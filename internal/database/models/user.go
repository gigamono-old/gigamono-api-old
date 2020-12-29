package models

import (
	"gorm.io/gorm"
)

// User stores information about the user.
type User struct {
	gorm.Model
	ID        string
	ProfileID string
	// passwordCredID string
	// socialLoginID string
}
