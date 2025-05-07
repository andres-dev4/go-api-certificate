package database

import (
	"log"

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

type Certificate struct {
	gorm.Model
	UserID      string `gorm:"type:uuid;not null;index"`
	Title       string `gorm:"size:255;not null"`
	Description string `gorm:"type:text"`
	IssueDate   string `gorm:"type:date;not null"`
	ExpiryDate  string `gorm:"type:date"`
	Credential  string `gorm:"type:text;not null"`
}

// User define el modelo de usuario en la base de datos
type User struct {
	gorm.Model
	Name  string `gorm:"size:255;not null"`
	Email string `gorm:"size:255;not null;uniqueIndex"`
	// Agrega más campos según sea necesario
}