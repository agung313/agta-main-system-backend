package models

type Blacklist struct {
	Token string `gorm:"unique"`
}
