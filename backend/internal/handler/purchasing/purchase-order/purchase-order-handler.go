package purchase_order

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	poDto "github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto/purchasing/purchase-order"
	poUsecase "github.com/nathanchristiawan02/icon-commission-system-backend/internal/usecase/purchasing/purchase-order"
)

type PurchaseOrderHandler struct {
	usecase poUsecase.PurchaseOrderUsecase
}

func NewPurchaseOrderHandler(
	usecase poUsecase.PurchaseOrderUsecase,
) *PurchaseOrderHandler {
	return &PurchaseOrderHandler{
		usecase: usecase,
	}
}

func (h *PurchaseOrderHandler) GetAll(c *gin.Context) {

	data, err := h.usecase.GetAll(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (h *PurchaseOrderHandler) GetByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid purchase order id",
		})
		return
	}

	data, err := h.usecase.GetByID(
		c.Request.Context(),
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (h *PurchaseOrderHandler) GetItems(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid purchase order id",
		})
		return
	}

	data, err := h.usecase.GetItems(
		c.Request.Context(),
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (h *PurchaseOrderHandler) Create(c *gin.Context) {

	var req poDto.CreatePurchaseOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	data, err := h.usecase.Create(
		c.Request.Context(),
		req,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "purchase order created successfully",
		"data":    data,
	})
}

func (h *PurchaseOrderHandler) Update(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid purchase order id",
		})
		return
	}

	var req poDto.UpdatePurchaseOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	data, err := h.usecase.Update(
		c.Request.Context(),
		id,
		req,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "purchase order updated successfully",
		"data":    data,
	})
}

func (h *PurchaseOrderHandler) Submit(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid purchase order id",
		})
		return
	}

	err = h.usecase.Submit(
		c.Request.Context(),
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "purchase order submitted successfully",
		"status":  "ordered",
	})
}

func (h *PurchaseOrderHandler) Delete(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid purchase order id",
		})
		return
	}

	err = h.usecase.Delete(
		c.Request.Context(),
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "purchase order deleted successfully",
	})
}