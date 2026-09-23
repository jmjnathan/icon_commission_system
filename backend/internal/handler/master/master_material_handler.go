package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	dto "github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto/master"
	usecase "github.com/nathanchristiawan02/icon-commission-system-backend/internal/usecase/master"
)

type MasterMaterialHandler struct {
	usecase usecase.MasterMaterialUsecase
}

func NewMasterMaterialHandler(usecase usecase.MasterMaterialUsecase) *MasterMaterialHandler {
	return &MasterMaterialHandler{usecase: usecase}
}

func (h *MasterMaterialHandler) GetAll(c *gin.Context) {
	sizes, err := h.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sizes})
}

func (h *MasterMaterialHandler) Create(c *gin.Context) {
	var req dto.MasterMaterialRequest
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

func (h *MasterMaterialHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req dto.MasterMaterialRequest
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
func (h *MasterMaterialHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.usecase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil dihapus"})
}
