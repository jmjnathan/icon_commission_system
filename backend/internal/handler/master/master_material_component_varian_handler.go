package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	dto "github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto/master"
	usecase "github.com/nathanchristiawan02/icon-commission-system-backend/internal/usecase/master"
)

type MaterialComponentVariantHandler struct {
	usecase usecase.MaterialComponentVariantUsecase
}

func NewMaterialComponentVariantHandler(
	usecase usecase.MaterialComponentVariantUsecase,
) *MaterialComponentVariantHandler {
	return &MaterialComponentVariantHandler{
		usecase: usecase,
	}
}

func (h *MaterialComponentVariantHandler) GetAll(
	c *gin.Context,
) {
	result, err := h.usecase.GetAll()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": result,
		},
	)
}

func (h *MaterialComponentVariantHandler) GetByMaterialComponentID(
	c *gin.Context,
) {
	id, err := strconv.Atoi(
		c.Param("materialComponentId"),
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "ID material component tidak valid",
			},
		)

		return
	}

	result, err := h.usecase.GetByMaterialComponentID(
		uint(id),
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": result,
		},
	)
}

func (h *MaterialComponentVariantHandler) Create(
	c *gin.Context,
) {
	id, err := strconv.Atoi(
		c.Param("materialComponentId"),
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "ID material component tidak valid",
			},
		)

		return
	}

	var req dto.MaterialComponentVariantRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	username := c.MustGet("username").(string)

	result, err := h.usecase.Create(
		uint(id),
		req,
		username,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"data": result,
		},
	)
}

func (h *MaterialComponentVariantHandler) Update(
	c *gin.Context,
) {
	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "ID variant tidak valid",
			},
		)

		return
	}

	var req dto.MaterialComponentVariantRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	username := c.MustGet("username").(string)

	result, err := h.usecase.Update(
		uint(id),
		req,
		username,
	)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "Data variant tidak ditemukan",
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": result,
		},
	)
}

func (h *MaterialComponentVariantHandler) Delete(
	c *gin.Context,
) {
	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "ID variant tidak valid",
			},
		)

		return
	}

	if err := h.usecase.Delete(uint(id)); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Berhasil dihapus",
		},
	)
}