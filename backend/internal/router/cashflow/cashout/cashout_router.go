package cashout

import (
	"github.com/gin-gonic/gin"
	handler "github.com/nathanchristiawan02/icon-commission-system-backend/internal/handler/cashflow/cashout"
)

func RegisterCashOutRoutes(rg *gin.RouterGroup, h *handler.CashOutHandler) {
	rg.GET("/cashflow/cashout/get-list", h.GetAll)
	rg.POST("/cashflow/cashout/create", h.Create)
	rg.PUT("/cashflow/cashout/edit/:id", h.Update)
	rg.DELETE("/cashflow/cashout/delete/:id", h.Delete)
}
