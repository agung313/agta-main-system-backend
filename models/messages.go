package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email"`
	Content string `json:"content"`
}
