package master

import (
	"github.com/gin-gonic/gin"

	handler "github.com/nathanchristiawan02/icon-commission-system-backend/internal/handler/master"
)

func RegisterMaterialComponentVariantRoutes(
	rg *gin.RouterGroup,
	h *handler.MaterialComponentVariantHandler,
) {

	rg.GET(
		"/master/material-component/variants",
		h.GetAll,
	)

	rg.GET(
		"/master/material-component/:materialComponentId/variants",
		h.GetByMaterialComponentID,
	)

	rg.POST(
		"/master/material-component/:materialComponentId/variants",
		h.Create,
	)

	rg.PUT(
		"/master/material-component/variant/edit/:id",
		h.Update,
	)

	rg.DELETE(
		"/master/material-component/variant/delete/:id",
		h.Delete,
	)
}
