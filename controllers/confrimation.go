package controllers

import (
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/models"
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/utils"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ConfrimationInput struct {
	TransferAmount int    `json:"transferAmount"`
	ImageUrl       string `json:"imageUrl"`
	Description    string `json:"description"`
	OrderID        int    `json:"orderId"`
}

// GetConfrimationsByUser a Rating godoc
// @Summary Get Confrimations By User Id
// @Description Get list Confrimations refrence By userID
// @Tags Confrimation
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Confrimation
// @Router /confrimation [get]
func GeatAllConfrimationByUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var confrimation []models.Confrimation
	db.Where("User_id=?", int(c.GetUint("currentUser"))).Find(&confrimation)

	response := utils.ApiResponse("All Confrimation", http.StatusOK, "success", &confrimation)
	c.JSON(http.StatusOK, response)
}

// CreateConfrimation godoc
// @Summary Create New Confrimation.
// @Description Creating a new Confrimation.
// @Tags Confrimation
// @Param Body body ConfrimationInput true "the body to create a new Confrimation"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Confrimation.
// @Router /confrimation [post]
func CreateConfrimation(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input ConfrimationInput
	var Order models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Where("id = ?", input.OrderID).First(&Order).Error; err != nil {
		response := utils.ApiResponse("Order not found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	_, err := url.ParseRequestURI(input.ImageUrl)
	if err != nil {
		errors := utils.FormatValidationEror(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("Confrimation failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	confrimation := models.Confrimation{
		TransferAmount: input.TransferAmount,
		ImageUrl:       input.ImageUrl,
		Description:    input.Description,
		Date:           time.Now(),
		OrderID:        input.OrderID,
		UserID:         int(c.GetUint("currentUser")),
	}
	db.Create(&confrimation)
	response := utils.ApiResponse("Confrimation success create ", http.StatusOK, "success", &confrimation)
	c.JSON(http.StatusOK, response)

}

// UpdateConfrimation godoc
// @Summary Update Confrimation.
// @Description Update Confrimation by id.
// @Tags Confrimation
// @Param Body body ConfrimationInput true "the body to update a new Confrimation"
// @Produce json
// @Param id path string true "Confrimation id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Confrimation
// @Router /confrimation/{id} [patch]
func UpdateConfrimation(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var confrimation models.Confrimation
	var order models.Order
	if err := db.Where("id = ?", c.Param("id")).First(&confrimation).Error; err != nil {
		response := utils.ApiResponse("Confrimation not found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	if confrimation.UserID != int(c.GetUint("currentUser")) {
		response := utils.ApiResponse("This User can,t access update Confrimation", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	var input ConfrimationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Where("id = ?", input.OrderID).First(&order).Error; err != nil {
		response := utils.ApiResponse("ID Order not Found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	updateConfirmation := models.Confrimation{
		TransferAmount: input.TransferAmount,
		ImageUrl:       input.ImageUrl,
		Description:    input.Description,
		OrderID:        input.OrderID,
		UserID:         int(c.GetUint("currentUser")),
		UpdatedAt:      time.Now(),
	}

	db.Model(&confrimation).Updates(updateConfirmation)
	response := utils.ApiResponse("Confrimation success update ", http.StatusOK, "success", confrimation)
	c.JSON(http.StatusOK, response)
}

// DeleteConfrimation godoc
// @Summary Delete one Confrimation.
// @Description Delete a Confrimation by id.
// @Tags Confrimation
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path string true "Confrimation id"
// @Success 200 {object} map[string]boolean
// @Router /confrimation/{id} [delete]
func DeleteConfrimation(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Confrimation models.Confrimation
	if err := db.Where("id = ?", c.Param("id")).First(&Confrimation).Error; err != nil {
		response := utils.ApiResponse("Confrimation Dont Delete", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	db.Delete(&Confrimation)
	response := utils.ApiResponse("Order had delete", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

// ApproveConfrimation  godoc
// @Summary Approve Confrimation Payment.
// @Description Approve Confrimation Payment order only Role admin.
// @Tags Confrimation
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path string true "Confrimation id"
// @Success 200 {object} map[string]boolean
// @Router /confrimation/approve/{id} [get]
func ApproveConfrimation(c *gin.Context) {
	//ambil id yang akan di confrimation
	// check id confrimation
	//ambil order id
	//update order id jadi status payent paid
	db := c.MustGet("db").(*gorm.DB)

	var confrimation models.Confrimation
	var order models.Order
	var user models.User
	if err := db.Where("id = ?", int(c.GetUint("currentUser"))).First(&user).Error; err != nil {
		response := utils.ApiResponse("User not found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	if user.Role != "admin" {
		response := utils.ApiResponse("Must an admin", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	if err := db.Where("id = ?", c.Param("id")).First(&confrimation).Error; err != nil {
		response := utils.ApiResponse("Confrimation not found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	if err := db.Where("id = ?", confrimation.OrderID).First(&order).Error; err != nil {
		response := utils.ApiResponse("OrderId not found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	db.Model(&order).Update("StatusPayment", "paid")
	response := utils.ApiResponse("Confrimation success Approve ", http.StatusOK, "success", confrimation)
	c.JSON(http.StatusOK, response)
}
