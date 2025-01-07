package models

import (
	"gorm.io/gorm"
)

type About struct {
	gorm.Model
	OpeningText    *OpeningText    `json:"openingText" gorm:"embedded;embeddedPrefix:openingText_"`
	ClosingText    *ClosingText    `json:"closingText" gorm:"embedded;embeddedPrefix:closingText_"`
	ComitmentLists []ComitmentList `json:"comitmentLists" gorm:"foreignKey:AboutID"`
}

type OpeningText struct {
	Id string `json:"id"`
	En string `json:"en"`
}

type ClosingText struct {
	Id string `json:"id"`
	En string `json:"en"`
}

type ComitmentList struct {
	gorm.Model
	AboutID         uint             `json:"aboutId"`
	TitleText       *TitleText       `json:"titleText" gorm:"embedded;embeddedPrefix:titleText_"`
	DescriptionText *DescriptionText `json:"descriptionText" gorm:"embedded;embeddedPrefix:descriptionText_"`
}

type TitleText struct {
	Id string `json:"id"`
	En string `json:"en"`
}
type DescriptionText struct {
	Id string `json:"id"`
	En string `json:"en"`
}
