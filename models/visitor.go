package models

import "gorm.io/gorm"

type Visitor struct {
	gorm.Model
	Countries string `json:"countries"`
}
