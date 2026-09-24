package product

import (
	"github.com/gin-gonic/gin"

	productHandler "github.com/nathanchristiawan02/icon-commission-system-backend/internal/handler/master/product"
)

func RegisterProductRoutes(
	router *gin.RouterGroup,
	handler *productHandler.ProductHandler,
) {

	product := router.Group("/master/products")

	{
		product.GET("", handler.GetAll)

		product.GET("/search", handler.Search)

		product.GET("/:id", handler.GetByID)

		product.POST("", handler.Create)

		product.PUT("/:id", handler.Update)

		product.DELETE("/:id", handler.Delete)
	}
}