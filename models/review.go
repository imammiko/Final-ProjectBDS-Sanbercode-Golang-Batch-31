package models

import "time"

type (
	Review struct {
		ID          int       `gorm:"primary_key"`
		Star        int       `json:"star"`
		Description string    `json:"description"`
		UserID      int       `json:"usersId"`
		ProductID   int       `json:"productId"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
	}
)
