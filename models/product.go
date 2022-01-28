package models

import (
	"time"
)

type (
	Product struct {
		ID          int `gorm:"primary_key;"`
		Name        string
		Condition   string
		Description string
		ImageUrl    string
		Stock       int
		Price       int
		Heavy       string
		UserID      int
		CategoryID  int
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)
