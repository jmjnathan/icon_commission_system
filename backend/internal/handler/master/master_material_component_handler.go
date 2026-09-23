package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	dto "github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto/master"
	usecase "github.com/nathanchristiawan02/icon-commission-system-backend/internal/usecase/master"
)

type MaterialComponentHandler struct {
	usecase usecase.MaterialComponentUsecase
}

func NewMaterialComponentHandler(
	usecase usecase.MaterialComponentUsecase,
) *MaterialComponentHandler {
	return &MaterialComponentHandler{
		usecase: usecase,
	}
}

func (h *MaterialComponentHandler) GetAll(c *gin.Context) {
	data, err := h.usecase.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (h *MaterialComponentHandler) Create(c *gin.Context) {
	var req dto.MaterialComponentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	username := c.MustGet("username").(string)

	data, err := h.usecase.Create(req, username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": data,
	})
}

func (h *MaterialComponentHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID tidak valid",
		})
		return
	}

	var req dto.MaterialComponentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	username := c.MustGet("username").(string)

	data, err := h.usecase.Update(
		uint(id),
		req,
		username,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (h *MaterialComponentHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID tidak valid",
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
