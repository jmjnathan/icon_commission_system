package dto

type MasterSaintsRequest struct {
	Name string `json:"name" binding:"required"`
	Remark string `json:"remark"`
	Status string `json:"status"`
}