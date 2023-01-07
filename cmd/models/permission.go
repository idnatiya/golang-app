package models

import (
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	ID   uint   `gorm:"primary-key"`
	Slug string `gorm:"index:permission_slug_unique_index,unique"`
	Name string
}
