package routes

import (
	_ "github.com/andres-dev4/go-api-certificate/src/api/docs" // Documentación Swagger
	"github.com/andres-dev4/go-api-certificate/src/internal/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // Swagger: archivos estáticos
	ginSwagger "github.com/swaggo/gin-swagger" // Swagger: integración con Gin
)

// Router estructura que contiene los controladores necesarios para las rutas
type Router struct {
	certController *controllers.CertificateController
	baseController *controllers.BaseController
}

// NewRouter crea una nueva instancia del enrutador con los controladores inyectados
func NewRouter(
	certController *controllers.CertificateController,
	baseController *controllers.BaseController,
) *Router {
	return &Router{
		certController: certController,
		baseController: baseController,
	}
}

// Setup configura y devuelve el enrutador Gin con todas las rutas definidas
func (r *Router) Setup() *gin.Engine {
	router := gin.Default()

	// Configuración de Swagger
	// Ruta para acceder a la documentación interactiva de la API
	// URL típica: http://localhost:8080/docs/index.html
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Rutas base de la aplicación
	router.GET("/", r.baseController.HolaMundo)

	// API versión 1
	v1 := router.Group("/api/v1")
	{
		v1.GET("/", r.baseController.HolaMundo)
		// Grupo de rutas para certificados
		certificates := v1.Group("/certificates")
		{
			// Obtener todos los certificados
			// GET /api/v1/certificates
			certificates.GET("/", r.certController.GetAllCertificates)

			// Aquí se pueden agregar más rutas relacionadas con certificados:
			// certificates.POST("/", r.certController.CreateCertificate)
			// certificates.GET("/:id", r.certController.GetCertificateByID)
			// certificates.PUT("/:id", r.certController.UpdateCertificate)
			// certificates.DELETE("/:id", r.certController.DeleteCertificate)
		}

		// Grupo de rutas para usuarios (comentado como ejemplo futuro)
		// users := v1.Group("/users")
		// {
		//     users.GET("/", r.userController.GetAllUsers)
		//     users.POST("/", r.userController.CreateUser)
		// }
	}

	return router
}
