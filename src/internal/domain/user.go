package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `gorm:"size:255" validate:"required"`
	Email string `gorm:"size:255;uniqueIndex" validate:"required,email"`
}

// UserRepository interface (puerto)
type UserRepository interface {
	Create(user *User) error
}
