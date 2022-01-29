package models

import "time"

type (
	Category struct {
		ID          int       `gorm:"primary_key" json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Products    []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
		UserID      int
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
	}
)
