package master

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanchristiawan02/icon-commission-system-backend/internal/handler/master"
)

func RegisterMasterStyleRoutes(rg *gin.RouterGroup, h *handler.MasterStyleHandler) {
	rg.GET("/master/get-style", h.GetAll)
	rg.POST("/master/create-style", h.Create)
	rg.PUT("/master/style/edit/:id", h.Update)
	rg.DELETE("/master/style/delete/:id", h.Delete)
}
