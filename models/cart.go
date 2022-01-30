package models

import "time"

type (
	Cart struct {
		ID        int       `gorm:"primary_key"`
		Price     int       `json:"price"`
		Total     int       `json:"total"`
		Date      time.Time `json:"date"`
		UserID    int       `json:"usersId"`
		ProductID int       `json:"productId"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}
)
