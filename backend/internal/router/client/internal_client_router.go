package master

import (
	"github.com/gin-gonic/gin"
	handler "github.com/nathanchristiawan02/icon-commission-system-backend/internal/handler/client"
)

func RegisterClientRoutes(rg *gin.RouterGroup, h *handler.ClientHandler) {
	rg.GET("/master/get-clients", h.GetAll)
	rg.POST("/master/create-clients", h.Create)
	rg.PUT("/master/clients/edit/:id", h.Update)
	rg.DELETE("/master/clients/delete/:id", h.Delete)
}
