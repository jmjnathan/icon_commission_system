package commission

import (
	"github.com/gin-gonic/gin"
	handler "github.com/nathanchristiawan02/icon-commission-system-backend/internal/handler/commission"
)

func RegisterCommissionRoutes(rg *gin.RouterGroup, h *handler.CommissionHandler) {
	rg.GET("/commissions/get-list", h.GetAll)
	rg.POST("/commissions/create", h.Create)
	rg.PUT("/commissions/edit/:id", h.Update)
	rg.PATCH("/commissions/edit/:id/status", h.UpdateStatus)
	rg.DELETE("/commissions/delete/:id", h.Delete)
}