package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

func NewBaseController() *BaseController {
	return &BaseController{}
}

// HolaMundo godoc
// @Summary Endpoint de ejemplo
// @Description Devuelve un Hola Mundo
// @Tags ejemplo
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Hola Mundo!"
// @Router / [get]
func (bc *BaseController) HolaMundo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hola Mundo!",
	})
}
