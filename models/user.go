package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID               uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Name             string         `json:"name"`
	Username         string         `json:"username"`
	Email            string         `json:"email"`
	Role             string         `json:"role"`
	Password         string         `json:"password"`
	OriginalPassword string         `json:"original_password"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}
	return
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if user.Password != "" {
		user.OriginalPassword = user.Password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	return
}
