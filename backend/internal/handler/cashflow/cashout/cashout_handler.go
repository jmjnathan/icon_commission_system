package cashout

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	dto "github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto/cashflow/cashout"
	usecase "github.com/nathanchristiawan02/icon-commission-system-backend/internal/usecase/cashflow/cashout"
)

type CashOutHandler struct {
	usecase usecase.CashOutUsecase
}

func NewCashOutHandler(usecase usecase.CashOutUsecase) *CashOutHandler {
	return &CashOutHandler{usecase: usecase}
}

func (h *CashOutHandler) GetAll(c *gin.Context) {
	items, err := h.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": items})
}

func (h *CashOutHandler) Create(c *gin.Context) {
	var req dto.CashOutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	username := c.MustGet("username").(string)

	item, err := h.usecase.Create(req, username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": item})
}

func (h *CashOutHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req dto.CashOutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	username := c.MustGet("username").(string)

	item, err := h.usecase.Update(uint(id), req, username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": item})
}

func (h *CashOutHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.usecase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil dihapus"})
}