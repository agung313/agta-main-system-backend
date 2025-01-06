package models

import "gorm.io/gorm"

type Slogan struct {
	gorm.Model
	FirstText   string       `json:"firstText"`
	SecondText  string       `json:"secondText"`
	ThirdText   string       `json:"thirdText"`
	Description *Description `json:"description" gorm:"embedded;embeddedPrefix:description_"`
}

type Description struct {
	Id string `json:"id"`

	En string `json:"en"`
}
