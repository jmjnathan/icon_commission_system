package commission

import (
	"github.com/gin-gonic/gin"

	handler "github.com/nathanchristiawan02/icon-commission-system-backend/internal/handler/commission"
)

func RegisterCommissionItemMaterialRoutes(
	rg *gin.RouterGroup,
	h *handler.CommissionItemMaterialHandler,
) {
	rg.GET(
		"/commission-items/:commissionItemId/materials",
		h.GetAll,
	)

	rg.POST(
		"/commission-items/:commissionItemId/materials",
		h.Create,
	)

	rg.PUT(
		"/commission-item-materials/:id",
		h.Update,
	)

	rg.DELETE(
		"/commission-item-materials/:id",
		h.Delete,
	)
}
