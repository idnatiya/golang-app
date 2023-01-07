package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID          uint   `gorm:"primary-key"`
	Slug        string `gorm:"index:role_slug_unique_index,unique"`
	Name        string
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}
