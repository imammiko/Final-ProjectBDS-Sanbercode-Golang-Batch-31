package models

import (
	"time"
)

type (
	User struct {
		ID            int    `json:"id" gorm:"primary_key"`
		Username      string `gorm:"not null;unique" json:"username"`
		Email         string `gorm:"not null;unique" json:"email"`
		Name          string `gorm:"not null" json:"name"`
		Password      string `gorm:"not null" json:"password"`
		DateOfBirth   string `gorm:"not null" json:"dateOfBirth"`
		Gender        string
		PhoneNumber   string `json:"phoneNumber"`
		Role          string
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
