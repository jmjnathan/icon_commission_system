package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathanchristiawan02/icon-commission-system-backend/internal/handler"
)

func RegisterAuthRoutes(router *gin.Engine, h *handler.AuthHandler) {
	router.POST("/login", h.Login)
}

func RegisterMeRoute(rg *gin.RouterGroup) {
	rg.GET("/me", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"admin_id": c.MustGet("admin_id"),
			"username": c.MustGet("username"),
		})
	})
}