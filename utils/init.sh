#Script bash para inicializar el proyecto creando la estructura

# Inicializar módulo Go
go mod init github.com/andres-dev4/go-api-certificate

# Crear estructura de directorios
# Estructura principal
mkdir -p go-api-certificate/cmd/api
mkdir -p go-api-certificate/api/docs
mkdir -p go-api-certificate/pkg/database

# Directorios internos
mkdir -p go-api-certificate/internal/config
mkdir -p go-api-certificate/internal/controllers
mkdir -p go-api-certificate/internal/domain
mkdir -p go-api-certificate/internal/interfaces
mkdir -p go-api-certificate/internal/routes
mkdir -p go-api-certificate/internal/utils

# Infraestructura
mkdir -p go-api-certificate/internal/infrastructure/repositories

# Instalar dependencias
go get -u \
  github.com/gin-gonic/gin \
  gorm.io/gorm \
  gorm.io/driver/postgres \
  github.com/go-playground/validator/v10 \
  github.com/swaggo/swag \
  github.com/swaggo/gin-swagger


#instalar swaggo 
go install github.com/swaggo/swag/cmd/swag@latest

