package domain

import (
	"github.com/andres-dev4/go-api-certificate/src/internal/models"
)

type User = models.User

// UserRepository interface (puerto)
type UserRepository interface {
	Create(user *User) error
}
