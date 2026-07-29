package commission

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	dto "github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto/commission"
	usecase "github.com/nathanchristiawan02/icon-commission-system-backend/internal/usecase/commission"
)

type CommissionHandler struct {
	usecase usecase.CommissionUsecase
}

func NewCommissionHandler(usecase usecase.CommissionUsecase) *CommissionHandler {
	return &CommissionHandler{usecase: usecase}
}

func (h *CommissionHandler) GetAll(c *gin.Context) {
	commissions, err := h.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": commissions})
}

func (h *CommissionHandler) Create(c *gin.Context) {
	var req dto.CommissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := c.MustGet("username").(string)

	result, err := h.usecase.Create(req, username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": result})
}

func (h *CommissionHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req dto.CommissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := c.MustGet("username").(string)

	result, err := h.usecase.Update(uint(id), req, username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (h *CommissionHandler) UpdateStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req dto.CommissionStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := c.MustGet("username").(string)

	result, err := h.usecase.UpdateStatus(uint(id), req, username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (h *CommissionHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.usecase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil dihapus"})
}