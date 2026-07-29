package master

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanchristiawan02/icon-commission-system-backend/internal/handler/master"
)

func RegisterMasterSizeRoutes(rg *gin.RouterGroup, h *handler.MasterSizeHandler) {
	rg.GET("/master/get-sizes", h.GetAll)
	rg.POST("/master/create-sizes", h.Create)
	rg.PUT("/master/sizes/edit/:id", h.Update)
	rg.DELETE("/master/size/delete/:id", h.Delete)
}