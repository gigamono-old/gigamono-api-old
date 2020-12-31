package models

// Role represents roles with assigned permissions.
type Role struct {
	Base
	Name           string
	Description    string
	XApp           []*App           `gorm:"many2many:apps_x_roles"`
	XUser          []*User          `gorm:"many2many:users_x_roles"`
	XAccessControl []*AccessControl `gorm:"many2many:access_controls_x_roles"`
	XTheme         []*Theme         `gorm:"many2many:themes_x_roles"`
	XWorkflow      []*Workflow      `gorm:"many2many:workflow_x_roles"`
}
