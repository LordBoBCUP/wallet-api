package models

type Card struct {
	Id string `json:"id"`
	Merchant string `json:"merchant"`
	Name string `json:"name"`
	Number string `json:"number"`
	Expiry string  `json:"expiryDate"`
	Format string `json:"format"`
	Notes string `json:"notes"`
	Colour string `json:"colour"`
}