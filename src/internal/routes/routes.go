package routes

import (
	_ "github.com/andres-dev4/go-api-certificate/src/api/docs"
	"github.com/andres-dev4/go-api-certificate/src/internal/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	certController *controllers.CertificateController
	baseController *controllers.BaseController
}

func NewRouter(
	certController *controllers.CertificateController,
	baseController *controllers.BaseController,
) *Router {
	return &Router{
		certController: certController,
		baseController: baseController,
	}
}

func (r *Router) Setup() *gin.Engine {
	router := gin.Default()

	// Swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Rutas base
	router.GET("/", r.baseController.HolaMundo)

	// API v1
	v1 := router.Group("/api/v1")
	{
		// Certificados
		certificates := v1.Group("/certificates")
		{
			certificates.GET("/", r.certController.GetAllCertificates)
			// Agregar más rutas...
		}

		// Usuarios
		v1.GET("/users", r.baseController.HolaMundo)
	}

	return router
}
