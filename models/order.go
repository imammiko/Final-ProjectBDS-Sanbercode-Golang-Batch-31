package models

import (
	"database/sql/driver"
	"time"
)

type StatusPayment string

const (
	Paid   StatusPayment = "paid"
	UnPaid StatusPayment = "unpaid"
)

func (e *StatusPayment) Scan(value interface{}) error {
	*e = StatusPayment(value.([]byte))
	return nil

}

func (e StatusPayment) Value() (driver.Value, error) {
	return string(e), nil
}

type (
	Order struct {
		ID             int            `gorm:"primary_key" json:"id"`
		RecipientsName string         `json:"recipientsName"`
		OrderDate      time.Time      `json:"orderDate"`
		City           string         `json:"city"`
		Address        string         `json:"address"`
		StatusPayment  StatusPayment  `gorm:"type:ENUM('paid','unpaid');" json:"statusPayment"`
		PhoneNumber    string         `json:"phoneNumber"`
		UserID         int            `json:"userID"`
		OrderDetails   []OrderDetail  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"orderDetails"`
		Confrimations  []Confrimation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"confrimations"`
		CreatedAt      time.Time      `json:"createdAt"`
		UpdatedAt      time.Time      `json:"updatedAt"`
	}
)
