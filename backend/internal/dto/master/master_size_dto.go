package dto

type MasterSizeRequest struct {
	Name string `json:"name" binding:"required"`
	Size string `json:"size" binding:"required"`
	Status string `json:"status"`
}