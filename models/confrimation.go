package models

import "time"

type (
	Confrimation struct {
		ID             int       `gorm:"primary_key"`
		TransferAmount int       `json:"transferAmount"`
		ImageUrl       string    `json:"imageUrl"`
		Description    string    `json:"description"`
		Date           time.Time `json:"date"`
		OrderID        int       `json:"orderId"`
		UserID         int       `json:"userId"`
		CreatedAt      time.Time `json:"createdAt"`
		UpdatedAt      time.Time `json:"updatedAt"`
	}
)
