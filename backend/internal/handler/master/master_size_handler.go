package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	dto "github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto/master"
	usecase "github.com/nathanchristiawan02/icon-commission-system-backend/internal/usecase/master"
)

type MasterSizeHandler struct {
	usecase usecase.MasterSizeUsecase
}

func NewMasterSizeHandler(usecase usecase.MasterSizeUsecase) *MasterSizeHandler {
	return &MasterSizeHandler{usecase: usecase}
}

func (h *MasterSizeHandler) GetAll(c *gin.Context) {
	sizes, err := h.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sizes})
}

func (h *MasterSizeHandler) Create(c *gin.Context) {
	var req dto.MasterSizeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := c.MustGet("username").(string)

	size, err := h.usecase.Create(req, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": size})
}

func (h *MasterSizeHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req dto.MasterSizeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := c.MustGet("username").(string)

	size, err := h.usecase.Update(uint(id), req, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": size})
}
func (h *MasterSizeHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.usecase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil dihapus"})
}