package models

import "time"

type User struct {
	ID              uint   `gorm:"primary_key"`
	FirstName       string `gorm:"index"`
	LastName        string `gorm:"index"`
	Email           string `gorm:"index:user_email_unique_index,unique"`
	Password        string
	EmailVerifiedAt time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
