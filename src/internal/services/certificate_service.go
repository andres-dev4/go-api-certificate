package services

import (
	"context"

	"github.com/andres-dev4/go-api-certificate/src/internal/domain"
)

type CertificateService interface {
	GetAllCertificates(ctx context.Context) ([]domain.Certificate, error)
	GetCertificateByID(ctx context.Context, id string) (*domain.Certificate, error)
	CreateCertificate(ctx context.Context, certificate *domain.Certificate) error
	UpdateCertificate(ctx context.Context, id string, certificate *domain.Certificate) error
	DeleteCertificate(ctx context.Context, id string) error
	GetCertificatesByUserID(ctx context.Context, userID string) ([]domain.Certificate, error)
}

type certificateService struct {
	certRepo domain.CertificateRepository
}

func NewCertificateService(certRepo domain.CertificateRepository) CertificateService {
	return &certificateService{
		certRepo: certRepo,
	}
}

func (s *certificateService) GetAllCertificates(ctx context.Context) ([]domain.Certificate, error) {
	return s.certRepo.GetAll(ctx)
}

func (s *certificateService) GetCertificateByID(ctx context.Context, id string) (*domain.Certificate, error) {
	return s.certRepo.GetByID(ctx, id)
}

func (s *certificateService) CreateCertificate(ctx context.Context, certificate *domain.Certificate) error {
	return s.certRepo.Create(ctx, certificate)
}

func (s *certificateService) UpdateCertificate(ctx context.Context, id string, certificate *domain.Certificate) error {
	return s.certRepo.Update(ctx, id, certificate)
}

func (s *certificateService) DeleteCertificate(ctx context.Context, id string) error {
	return s.certRepo.Delete(ctx, id)
}

func (s *certificateService) GetCertificatesByUserID(ctx context.Context, userID string) ([]domain.Certificate, error) {
	return s.certRepo.GetByUserID(ctx, userID)
}
