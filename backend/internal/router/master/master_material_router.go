package master

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanchristiawan02/icon-commission-system-backend/internal/handler/master"
)

func RegisterMasterMaterialRoutes(rg *gin.RouterGroup, h *handler.MasterMaterialHandler) {
	rg.GET("/master/get-material", h.GetAll)
	rg.POST("/master/create-material", h.Create)
	rg.PUT("/master/material/edit/:id", h.Update)
	rg.DELETE("/master/material/delete/:id", h.Delete)
}