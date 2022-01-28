package models

import "time"

type (
	Cart struct {
		ID        int `gorm:"primary_key"`
		Price     int
		Total     int
		Date      time.Time
		UsersID   int
		ProductID int
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
