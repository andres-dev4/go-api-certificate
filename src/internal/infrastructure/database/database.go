package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDB() (*gorm.DB, error) {
	// Cargar variables con valores por defecto
	host := getEnv("DB_HOST", "localhost")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "mydb")
	port := getEnv("DB_PORT", "5432")
	sslmode := getEnv("DB_SSLMODE", "disable")

	// Validar puerto
	if _, err := strconv.Atoi(port); err != nil {
		return nil, fmt.Errorf("DB_PORT debe ser numérico: %v", err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host,
		user,
		password,
		dbname,
		port,
		sslmode,
	)

	log.Printf("Intentando conectar a PostgreSQL...")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("error al conectar a PostgreSQL: %w\nDSN usado: %s", err, maskPassword(dsn))
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error al obtener instancia de DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("¡Conexión a PostgreSQL establecida correctamente!")
	return db, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func maskPassword(dsn string) string {
	const passwordLabel = "password="
	start := strings.Index(dsn, passwordLabel)
	if start == -1 {
		return dsn
	}
	start += len(passwordLabel)
	end := strings.Index(dsn[start:], " ")
	if end == -1 {
		return dsn[:start] + "xxxxx"
	}
	return dsn[:start] + "xxxxx" + dsn[start+end:]
}

// Función de conveniencia para usar directamente
func RunMigrations(db *gorm.DB) error {
	migrator := NewMigrator(db)
	return migrator.RunMigrations()
}
