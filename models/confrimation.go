package models

import "time"

type (
	Confrimation struct {
		ID             int `gorm:"primary_key"`
		TransferAmount int
		ImageUrl       string
		Description    string
		Date           string
		OrderID        int
		UserID         int
		CreatedAt      time.Time
		UpdatedAt      time.Time
	}
)
