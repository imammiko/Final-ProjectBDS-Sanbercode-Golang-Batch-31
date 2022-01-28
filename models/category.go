package models

import "time"

type (
	Category struct {
		ID          int `gorm:"primary_key"`
		Name        string
		Description string
		Products    []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)
