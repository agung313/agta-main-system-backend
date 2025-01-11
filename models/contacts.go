package models

import "gorm.io/gorm"

type Contacts struct {
	gorm.Model
	Email        string `json:"email"`
	Instagram    string `json:"instagram"`
	LinkedinLink string `json:"linkedinLink"`
	Address      string `json:"address"`
	AddressLink  string `json:"addressLink"`
	Phone        string `json:"phone"`
}
