package domain

import (
	"context"
	"errors"
)

var (
	ErrCertificateNotFound = errors.New("certificate not found")
)

type Certificate struct {
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

// Resto de tu código (interfaces, etc.)...
type CertificateRepository interface {
	GetAll(ctx context.Context) ([]Certificate, error)
	GetByID(ctx context.Context, id string) (*Certificate, error)
	Create(ctx context.Context, certificate *Certificate) error
	Update(ctx context.Context, id string, certificate *Certificate) error
	Delete(ctx context.Context, id string) error
	GetByUserID(ctx context.Context, userID string) ([]Certificate, error)
}
