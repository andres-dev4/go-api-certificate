# Go API Certificate 🚀
Una API robusta para gestión de certificados digitales construida con Go y Gin.

## Características Clave
- ✅ Autenticación JWT
- ✅ Migraciones automáticas con GORM
- ✅ Documentación Swagger integrada
- ✅ Arquitectura en capas

## Estructura de Directorios

````bash
.
├── go.mod
├── go.sum
├── init.sh
├── LICENSE
├── README.md
├── src
│   ├── api
│   │   └── docs # Documentación Swagger
│   ├── cmd
│   │   └── api # Punto de entrada principal
│   ├── internal 
│   │   ├── config # Configuración de la aplicación
│   │   ├── controllers  # Manejo de requests HTTP
│   │   ├── domain  # Modelos de negocio e interfaces
│   │   ├── infrastructure # Configuración DB y migraciones
│   │   │   ├── database
│   │   │   └── repositories
│   │   ├── interfaces
│   │   ├── models #Modelos ORM
│   │   ├── routes # Definición de endpoints
│   │   ├── services  # Lógica de negocio
│   │   └── utils # Utilidades compartidas
└── utils
    ├── create_db_docker.sh # Creacion Base de datos
    ├── createmod.sh # Creacion mod
    └── init.sh # comando de creación de proyecto
````

# 🚀 Guía de Deployment

## Entorno Local

```bash
# Asignar permisos a init.sh
chmod +x init.sh

# Ejecutar
./init.sh
````
