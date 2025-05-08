package domain

import (
	"context"
	"errors"

	"github.com/andres-dev4/go-api-certificate/src/internal/models"
)

type Certificate = models.Certificate

var (
	ErrCertificateNotFound = errors.New("certificate not found")
)

// Resto de tu código (interfaces, etc.)...
type CertificateRepository interface {
	GetAll(ctx context.Context) ([]Certificate, error)
	GetByID(ctx context.Context, id string) (*Certificate, error)
	Create(ctx context.Context, certificate *Certificate) error
	Update(ctx context.Context, id string, certificate *Certificate) error
	Delete(ctx context.Context, id string) error
	GetByUserID(ctx context.Context, userID string) ([]Certificate, error)
}
