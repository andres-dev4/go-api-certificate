// @title Certificate API
// @version 1.0
// @description API para gestión de certificados
// @BasePath /api/v1
package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/andres-dev4/go-api-certificate/src/internal/config"
	"github.com/andres-dev4/go-api-certificate/src/internal/controllers"
	"github.com/andres-dev4/go-api-certificate/src/internal/infrastructure/database"
	"github.com/andres-dev4/go-api-certificate/src/internal/infrastructure/repositories"
	"github.com/andres-dev4/go-api-certificate/src/internal/routes"
	"github.com/andres-dev4/go-api-certificate/src/internal/services"
)

func main() {
	// Cargar configuración
	if err := godotenv.Load(); err != nil {
		log.Printf("No se encontró archivo .env: %v", err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Inicializar base de datos
	db, err := database.NewPostgresDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Ejecutar migraciones
	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Inicializar controladores
	baseController := controllers.NewBaseController()
	certRepo := repositories.NewCertificateRepository(db)
	certService := services.NewCertificateService(certRepo)
	certController := controllers.NewCertificateController(certService)

	// Configurar router
	router := routes.NewRouter(certController, baseController).Setup()

	// Iniciar servidor
	log.Printf("Server running on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
