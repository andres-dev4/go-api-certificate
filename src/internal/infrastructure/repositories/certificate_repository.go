package repositories

import (
	"context"
	"errors"

	"github.com/andres-dev4/go-api-certificate/src/internal/domain"
	"gorm.io/gorm"
)

type certificateRepository struct {
	db *gorm.DB
}

func NewCertificateRepository(db *gorm.DB) domain.CertificateRepository {
	return &certificateRepository{db: db}
}

func (r *certificateRepository) GetAll(ctx context.Context) ([]domain.Certificate, error) {
	var certificates []domain.Certificate
	if err := r.db.WithContext(ctx).Find(&certificates).Error; err != nil {
		return nil, err
	}
	return certificates, nil
}

func (r *certificateRepository) GetByID(ctx context.Context, id string) (*domain.Certificate, error) {
	var certificate domain.Certificate
	if err := r.db.WithContext(ctx).First(&certificate, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrCertificateNotFound
		}
		return nil, err
	}
	return &certificate, nil
}

func (r *certificateRepository) Create(ctx context.Context, certificate *domain.Certificate) error {
	return r.db.WithContext(ctx).Create(certificate).Error
}

func (r *certificateRepository) Update(ctx context.Context, id string, certificate *domain.Certificate) error {
	result := r.db.WithContext(ctx).Model(&domain.Certificate{}).Where("id = ?", id).Updates(certificate)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrCertificateNotFound
	}
	return nil
}

func (r *certificateRepository) Delete(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(&domain.Certificate{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrCertificateNotFound
	}
	return nil
}

func (r *certificateRepository) GetByUserID(ctx context.Context, userID string) ([]domain.Certificate, error) {
	var certificates []domain.Certificate
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&certificates).Error; err != nil {
		return nil, err
	}
	return certificates, nil
}
