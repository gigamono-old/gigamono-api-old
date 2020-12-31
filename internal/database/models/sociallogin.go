package models

import (
	"github.com/google/uuid"
)

// SocialLogin for social login information.
type SocialLogin struct {
	Base
	AppName string
	UserID  uuid.UUID
}
