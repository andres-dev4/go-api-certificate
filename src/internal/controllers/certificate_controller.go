package controllers

import (
	"net/http"

	"github.com/andres-dev4/go-api-certificate/src/internal/models"
	"github.com/andres-dev4/go-api-certificate/src/internal/services"
	"github.com/gin-gonic/gin"
)

type Certificate = models.Certificate
type CertificateController struct {
	service services.CertificateService
}

func NewCertificateController(service services.CertificateService) *CertificateController {
	return &CertificateController{service: service}
}

// GetAllCertificates godoc
// @Summary List all certificates
// @Description Get all certificates
// @Tags certificates
// @Accept json
// @Produce json
// @Success 200 {array} models.Certificate
// @Router /certificates [get]
func (c *CertificateController) GetAllCertificates(ctx *gin.Context) {
	certificates, err := c.service.GetAllCertificates(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, certificates)
}

// Implementa los demás métodos del controlador...
