package user

import "Final-ProjectBDS-Sanbercode-Golang-Batch-31/models"

type UserFormatter struct {
	ID          int    `json:"id"`
	Name        string `json:"string"`
	Email       string `json:"email"`
	DateOfBirth string `json:"dateOfBirth"`
	Gender      string `json:"gender"`
	Role        string `json:"role"`
	PhoneNumber string `json:"phoneNumber"`
	Token       string `json:"Token"`
}

func FormatUser(user models.User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		DateOfBirth: user.DateOfBirth,
		Gender:      user.Gender,
		Role:        user.Role,
		PhoneNumber: user.PhoneNumber,
		Token:       token,
	}
	return formatter
}
