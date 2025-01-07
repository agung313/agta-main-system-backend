package models

import "gorm.io/gorm"

type Contacts struct {
	gorm.Model
	Email          string `json:"email"`
	Instagram      string `json:"instagram"`
	Linkedin       string `json:"linkedin"`
	Address        string `json:"address"`
	GoogleMapsLink string `json:"googleMapsLink"`
}
