package purchase_order

import (
	"github.com/gin-gonic/gin"

	purchaseOrderHandler "github.com/nathanchristiawan02/icon-commission-system-backend/internal/handler/purchasing/purchase-order"
)

func RegisterPurchaseOrderRoutes(
	router *gin.RouterGroup,
	handler *purchaseOrderHandler.PurchaseOrderHandler,
) {

	purchaseOrder := router.Group("/purchase-order")

	{
		purchaseOrder.GET("", handler.GetAll)

		purchaseOrder.GET("/:id", handler.GetByID)

		purchaseOrder.GET("/:id/items", handler.GetItems)

		purchaseOrder.POST("", handler.Create)

		purchaseOrder.PUT("/:id", handler.Update)

		// Submit Draft → Ordered
		purchaseOrder.POST("/:id/submit", handler.Submit)

		purchaseOrder.DELETE("/:id", handler.Delete)
	}
}