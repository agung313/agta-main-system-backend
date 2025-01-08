package models

import "gorm.io/gorm"

type TokenAdmin struct {
	gorm.Model
	Token string `gorm:"unique"`
}
