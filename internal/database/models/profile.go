package models

import (
	"gorm.io/gorm"
)

// Profile stores information about the user.
type Profile struct {
	gorm.Model
	ID          string
	Username    string
	FirstName   string
	SecondName  string
	Email       string
	Avatar32URL string
	UserID      string
}
