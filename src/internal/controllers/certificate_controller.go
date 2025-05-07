package controllers

import (
	"net/http"

	"github.com/andres-dev4/go-api-certificate/src/internal/services"
	"github.com/gin-gonic/gin"
)

type CertificateController struct {
	service services.CertificateService
}

func NewCertificateController(service services.CertificateService) *CertificateController {
	return &CertificateController{service: service}
}

func (c *CertificateController) GetAllCertificates(ctx *gin.Context) {
	certificates, err := c.service.GetAllCertificates(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, certificates)
}

// Implementa los demás métodos del controlador...
