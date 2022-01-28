package user

import "Final-ProjectBDS-Sanbercode-Golang-Batch-31/models"

type RegisterUserInput struct {
	Name        string        `json:"name" binding:"required"`
	Username    string        `json:"username" binding:"required"`
	Email       string        `json:"email" binding:"required,email"`
	Password    string        `json:"password" binding:"required"`
	DateOfBirth string        `json:"dateOfBirth"`
	Gender      models.Gender `gorm:"type:ENUM('man','woman');" json:"gender"`
	PhoneNumber string        `json:"phoneNumber"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type ChangePassword struct {
	Email       string `json:"email" binding:"required,email"`
	PasswordOld string `json:"passwordOld" binding:"required"`
	PasswordNew string `json:"passwordNew" binding:"required"`
}
