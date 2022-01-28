package models

import "time"

type (
	OrderDetail struct {
		ID        int `gorm:"primary_key"`
		Harga     int
		Jumlah    int
		ProductID int
		OrderID   int
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
