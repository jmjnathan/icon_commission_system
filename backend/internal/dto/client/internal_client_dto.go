package dto

type ClientRequest struct {
	Name                   string `json:"name" binding:"required"`
	Nickname               string `json:"nickname"`
	Phone                  string `json:"phone" binding:"required"`
	AlternatePhone         string `json:"alternate_phone"`
	Email                  string `json:"email"`
	InstagramHandle        string `json:"instagram_handle"`
	Address                string `json:"address"`
	City                   string `json:"city"`
	Province               string `json:"province"`
	PostalCode             string `json:"postal_code"`
	Country                string `json:"country"`
	Denomination           string `json:"denomination"`
	PatronSaintPreference  string `json:"patron_saint_preference"`
	ParishName             string `json:"parish_name"`
	PreferredPaymentMethod string `json:"preferred_payment_method"`
	Notes                  string `json:"notes"`
}
