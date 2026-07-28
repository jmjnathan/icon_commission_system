package entity

type MasterSize struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Size string `json:"size"`
	BaseModel
}