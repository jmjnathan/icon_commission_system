package master

import (
	"github.com/gin-gonic/gin"

	handler "github.com/nathanchristiawan02/icon-commission-system-backend/internal/handler/master"
)

func RegisterMaterialComponentRoutes(
	rg *gin.RouterGroup,
	h *handler.MaterialComponentHandler,
) {
	rg.GET(
		"/master/get-material-component",
		h.GetAll,
	)

	rg.POST(
		"/master/create-material-component",
		h.Create,
	)

	rg.PUT(
		"/master/material-component/edit/:id",
		h.Update,
	)

	rg.DELETE(
		"/master/material-component/delete/:id",
		h.Delete,
	)
}
