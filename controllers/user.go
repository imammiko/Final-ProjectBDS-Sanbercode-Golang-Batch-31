package controllers

import (
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/user"
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/utils"
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/utils/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userControllers struct {
	userService user.Service
	authService auth.Service
}

func NewUserController(userService user.Service, authService auth.Service) *userControllers {
	return &userControllers{userService, authService}
}

// ForgetPassword godoc
// @Summary Forget Password a User.
// @Description Reset password will send to email.
// @Tags Auth
// @Produce json
// @Param        email    query     string  false  "email forgot password"  Format(email)
// @Success 200 {object} utils.Response
// @Router /forgotPassword [get]
func (h *userControllers) ForgetPassword(c *gin.Context) {
	emailInput, bolean := c.GetQuery("email")
	if bolean {
		response := utils.ApiResponse("Query tidak ditemukan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err := h.userService.ForgotPassword(emailInput)
	if err != nil {
		response := utils.ApiResponse("forgot password error", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("forget password succes", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)

}

// ChangePassword godoc
// @Summary Change Password an user.
// @Description Change Password User.
// @Tags Auth
// @Param Body body user.ChangePassword true "the body to register a user"
// @Produce json
// @Success 200 {object} utils.Response
// @Router /changePassword [post]
func (h *userControllers) ChangePassword(c *gin.Context) {
	var input user.ChangePassword

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationEror(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	_, err1 := h.userService.ChangePassword(input.Email, input.PasswordNew, input.PasswordOld)
	if err1 != nil {
		response := utils.ApiResponse("change password failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := utils.ApiResponse("Account has been changed password", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

// Register godoc
// @Summary Register a user.
// @Description registering a user from public access.
// @Tags Auth
// @Param Body body user.RegisterUserInput true "the body to register a user"
// @Produce json
// @Success 200 {object} user.UserFormatter
// @Router /register [post]
func (h *userControllers) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationEror(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := utils.ApiResponse("Register account hailed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := utils.ApiResponse("Register account failde", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
	}
	formatter := user.FormatUser(newUser, token)

	response := utils.ApiResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

	// err:= c.ShouldBindJSON(inpu)
}

// LoginUser godoc
// @Summary Login as as user.
// @Description Logging in to get jwt token to access admin or user api by roles.
// @Tags Auth
// @Param Body body user.LoginInput true "the body to login a user"
// @Produce json
// @Success 200 {object} user.UserFormatter
// @Router /login [post]
func (h *userControllers) Login(c *gin.Context) {
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationEror(err)
		errorMessage := gin.H{
			"errors": errors,
		}
		response := utils.ApiResponse("Login failed", http.StatusUnprocessableEntity, "errors", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := utils.ApiResponse("Loggin Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := utils.ApiResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(loggedinUser, token)
	response := utils.ApiResponse("succesfuly loggedin", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
