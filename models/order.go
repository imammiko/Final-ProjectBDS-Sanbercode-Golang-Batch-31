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
		ID             int `gorm:"primary_key"`
		RecipientsName string
		OrderDate      time.Time
		City           string
		Address        string
		StatusPayment  StatusPayment `gorm:"type:ENUM('paid','unpaid');"`
		PhoneNumber    string
		UserID         int
		CreatedAt      time.Time
		UpdatedAt      time.Time
	}
)
