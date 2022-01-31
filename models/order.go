package models

import (
	"time"
)

type (
	Order struct {
		ID             int            `gorm:"primary_key" json:"id"`
		RecipientsName string         `json:"recipientsName"`
		OrderDate      time.Time      `json:"orderDate"`
		City           string         `json:"city"`
		Address        string         `json:"address"`
		StatusPayment  string         `json:"statusPayment"`
		PhoneNumber    string         `json:"phoneNumber"`
		UserID         int            `json:"userID"`
		OrderDetails   []OrderDetail  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"orderDetails"`
		Confrimations  []Confrimation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"confrimations"`
		CreatedAt      time.Time      `json:"createdAt"`
		UpdatedAt      time.Time      `json:"updatedAt"`
	}
)
