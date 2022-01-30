package models

import (
	"database/sql/driver"
	"time"
)

type Gender string

const (
	Man   Gender = "man"
	Woman Gender = "woman"
)

func (e *Gender) Scan(value interface{}) error {
	*e = Gender(value.([]byte))
	return nil

}

func (e Gender) Value() (driver.Value, error) {
	return string(e), nil
}

type Role string

const (
	Cutomer Role = "customer"
	Admin   Role = "admin"
)

func (e *Role) Scan(value interface{}) error {
	*e = Role(value.([]byte))
	return nil

}

func (e Role) Value() (driver.Value, error) {
	return string(e), nil
}

type (
	User struct {
		ID            int            `json:"id" gorm:"primary_key"`
		Username      string         `gorm:"not null;unique" json:"username"`
		Email         string         `gorm:"not null;unique" json:"email"`
		Name          string         `gorm:"not null" json:"name"`
		Password      string         `gorm:"not null" json:"password"`
		DateOfBirth   string         `gorm:"not null" json:"dateOfBirth"`
		Gender        Gender         `gorm:"type:ENUM('man','woman');"`
		PhoneNumber   string         `json:"phoneNumber"`
		Role          Role           `gorm:"type:ENUM('customer','admin');"`
		Products      []Product      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		Orders        []Order        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		Confrimations []Confrimation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		Categories    []Category     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		Carts         []Cart         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		Review        []Review       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
		CreatedAt     time.Time      `json:"created_at"`
		UpdatedAt     time.Time      `json:"updated_at"`
	}
)
