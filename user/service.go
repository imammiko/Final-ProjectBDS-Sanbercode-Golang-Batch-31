package user

import (
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/models"
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/utils"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (models.User, error)
	Login(input LoginInput) (models.User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	GetUserById(ID int) (models.User, error)
	ForgotPassword(email string) (models.User, error)
	ChangePassword(email string, passwordNew string, passwordOld string) (models.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (models.User, error) {
	user := models.User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Name = input.Name
	user.DateOfBirth = input.Password
	user.Gender = input.Gender
	user.PhoneNumber = input.PhoneNumber
	user.Username = input.Username
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(password)
	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}
	return newUser, nil

}

func (s *service) Login(input LoginInput) (models.User, error) {
	email := input.Email
	password := input.Password
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("no user found on that email")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if user.ID == 0 {
		return true, nil
	}
	return false, nil
}

func (s *service) GetUserById(ID int) (models.User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("no user found on with that ID")
	}
	return user, nil
}

func (s *service) ChangePassword(email string, passwordNew string, passwordOld string) (models.User, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("no email user found on with that ID")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordOld))
	if err != nil {
		return user, err
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(passwordNew), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(password)
	userUpdated, err := s.repository.Update(user)
	if err != nil {
		return user, err
	}
	message := fmt.Sprintf("<h1>Your New password</h1><P>password baru anda:  %s</P>", string(passwordNew))
	go utils.SendingEmail(user.Email, "New Password", message)
	return userUpdated, nil
}

func (s *service) ForgotPassword(email string) (models.User, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("no email user found on with that ID")
	}
	newPassword := utils.StringWithCharset(9, utils.Charset)
	password, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.MinCost)
	user.Password = string(password)
	userModel, err := s.repository.Update(user)
	if err != nil {
		return user, err
	}
	message := fmt.Sprintf("<h1>Your forgot password</h1><P>password baru anda:  %s</P>", newPassword)
	go utils.SendingEmail(user.Email, "New Password because Forgot Password", message)
	return userModel, nil

}
