package repositories

import (
	"github.com/andres-dev4/go-api-certificate/src/internal/domain"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	DB *gorm.DB
}

func (r *GormUserRepository) Create(user *domain.User) error {
	return r.DB.Create(user).Error
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &GormUserRepository{DB: db}
}
