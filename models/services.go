package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Description    *ServiceDescription `json:"description" gorm:"embedded;embeddedPrefix:description_"`
	TechnologyList []TechnologyList    `json:"technologiesList" gorm:"foreignKey:ServiceId"`
}

type ServiceDescription struct {
	Id string `json:"id"`
	En string `json:"en"`
}

type TechnologyList struct {
	gorm.Model
	ServiceId   uint                       `json:"ServiceId"`
	Icont       string                     `json:"icont"`
	Title       string                     `json:"title"`
	Link        string                     `json:"link"`
	Description *TechnologyListDescription `json:"description" gorm:"embedded;embeddedPrefix:description_"`
}

type TechnologyListDescription struct {
	Id string `json:"id"`
	En string `json:"en"`
}
