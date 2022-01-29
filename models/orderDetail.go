package models

import "time"

type (
	OrderDetail struct {
		ID        int       `gorm:"primary_key" json:"id"`
		Harga     int       `json:"harga"`
		Jumlah    int       `json:"jumlah"`
		ProductID int       `json:"productId"`
		OrderID   int       `json:"orderId"`
		CreatedAt time.Time `json:"createdAT"`
		UpdatedAt time.Time `json:"updatedAT"`
	}
)
