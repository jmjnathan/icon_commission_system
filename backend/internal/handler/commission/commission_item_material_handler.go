package commission

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	dto "github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto/commission"
	usecase "github.com/nathanchristiawan02/icon-commission-system-backend/internal/usecase/commission"
)

type CommissionItemMaterialHandler struct {
	usecase usecase.CommissionItemMaterialUsecase
}

func NewCommissionItemMaterialHandler(
	usecase usecase.CommissionItemMaterialUsecase,
) *CommissionItemMaterialHandler {
	return &CommissionItemMaterialHandler{
		usecase: usecase,
	}
}

func (h *CommissionItemMaterialHandler) GetAll(
	c *gin.Context,
) {
	commissionItemID, err := strconv.Atoi(
		c.Param("commissionItemId"),
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID commission item tidak valid",
		})
		return
	}

	result, err := h.usecase.GetAll(
		uint(commissionItemID),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (h *CommissionItemMaterialHandler) Create(
	c *gin.Context,
) {
	commissionItemID, err := strconv.Atoi(
		c.Param("commissionItemId"),
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID commission item tidak valid",
		})
		return
	}

	var req dto.CommissionItemMaterialRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	username := c.MustGet("username").(string)

	result, err := h.usecase.Create(
		uint(commissionItemID),
		req,
		username,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": result,
	})
}

func (h *CommissionItemMaterialHandler) Update(
	c *gin.Context,
) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID material commission tidak valid",
		})
		return
	}

	var req dto.CommissionItemMaterialRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	username := c.MustGet("username").(string)

	result, err := h.usecase.Update(
		uint(id),
		req,
		username,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (h *CommissionItemMaterialHandler) Delete(
	c *gin.Context,
) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID material commission tidak valid",
		})
		return
	}

	if err := h.usecase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil dihapus",
	})
}
