package models

import "gorm.io/gorm"

type Certificate struct {
	gorm.Model
	ID          string `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID      string `json:"user_id" gorm:"type:uuid;not null;index"`
	Title       string `json:"title" gorm:"size:255;not null"`
	Description string `json:"description" gorm:"type:text"`
	IssueDate   string `json:"issue_date" gorm:"type:date;not null"`
	ExpiryDate  string `json:"expiry_date" gorm:"type:date"`
	Credential  string `json:"credential" gorm:"type:text;not null"`
	CreatedAt   int64  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   int64  `json:"updated_at" gorm:"autoUpdateTime"`
}

type User struct {
	gorm.Model
	Name  string `gorm:"size:255" validate:"required"`
	Email string `gorm:"size:255;uniqueIndex" validate:"required,email"`
}
