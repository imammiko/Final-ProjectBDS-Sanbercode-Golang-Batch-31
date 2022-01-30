package models

import (
	"time"
)

type (
	Product struct {
		ID           int           `gorm:"primary_key;" json:"id"`
		Name         string        `json:"name"`
		Condition    string        `json:"condition"`
		Description  string        `json:"description"`
		ImageUrl     string        `json:"ImageUrl"`
		Stock        int           `json:"stock"`
		Price        int           `json:"price"`
		Heavy        string        `json:"heavy"`
		UserID       int           `json:"userID"`
		CategoryID   int           `json:"categoryID"`
		Review       []Review      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
		Carts        []Cart        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
		OrderDetails []OrderDetail `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
		CreatedAt    time.Time     `json:"createdAT"`
		UpdatedAt    time.Time     `json:"updatedAT"`
	}
)
