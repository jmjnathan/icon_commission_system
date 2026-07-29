package master

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanchristiawan02/icon-commission-system-backend/internal/handler/master"
)

func RegisterMasterSaintsRoutes(rg *gin.RouterGroup, h *handler.MasterSaintsHandler) {
	rg.GET("/master/get-saints", h.GetAll)
	rg.POST("/master/create-saints", h.Create)
	rg.PUT("/master/saints/edit/:id", h.Update)
	rg.DELETE("/master/saints/delete/:id", h.Delete)
}