package product

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	productDto "github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto/master/product"
	productUsecase "github.com/nathanchristiawan02/icon-commission-system-backend/internal/usecase/master/product"
)

type ProductHandler struct {
	usecase productUsecase.ProductUsecase
}

func NewProductHandler(
	usecase productUsecase.ProductUsecase,
) *ProductHandler {
	return &ProductHandler{
		usecase: usecase,
	}
}

func (h *ProductHandler) GetAll(c *gin.Context) {

	data, err := h.usecase.GetAll(
		c.Request.Context(),
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

func (h *ProductHandler) GetByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid product id",
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

func (h *ProductHandler) Search(c *gin.Context) {

	keyword := c.Query("keyword")

	data, err := h.usecase.Search(
		c.Request.Context(),
		keyword,
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

func (h *ProductHandler) Create(c *gin.Context) {

	var req productDto.CreateProductRequest

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
		"message": "product created successfully",
		"data":    data,
	})
}

func (h *ProductHandler) Update(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid product id",
		})
		return
	}

	var req productDto.UpdateProductRequest

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
		"message": "product updated successfully",
		"data":    data,
	})
}

func (h *ProductHandler) Delete(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid product id",
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
		"message": "product deleted successfully",
	})
}