package commission

type CommissionPhoto struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	CommissionID uint   `json:"commission_id"`
	FileURL      string `json:"file_url"`
}
