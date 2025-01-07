package models

import "gorm.io/gorm"

type Visitor struct {
	gorm.Model
	Location string `json:"location"`
}
