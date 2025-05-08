package database

import (
	"log"

	"github.com/andres-dev4/go-api-certificate/src/internal/models"
	"gorm.io/gorm"
)

type Migrator struct {
	db *gorm.DB
}

func NewMigrator(db *gorm.DB) *Migrator {
	return &Migrator{db: db}
}

func (m *Migrator) RunMigrations() error {
	log.Println("Starting database migrations...")

	// Habilitar extensión UUID si usas PostgreSQL
	if err := m.db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		return err
	}

	// Ejecutar AutoMigrate para todos tus modelos
	if err := m.autoMigrateModels(); err != nil {
		return err
	}

	log.Println("Database migrations completed successfully")
	return nil
}

func (m *Migrator) autoMigrateModels() error {
	// Registra aquí todos tus modelos
	models := []interface{}{
		&Certificate{}, // Modelo de certificado
		&User{},        // Si tienes modelo de usuario
		// Agrega otros modelos aquí
	}

	for _, model := range models {
		if err := m.db.AutoMigrate(model); err != nil {
			return err
		}
		log.Printf("Migrated model: %T", model)
	}

	return nil
}

type User = models.User
type Certificate = models.Certificate
