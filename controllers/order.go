package controllers

import (
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/models"
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderInput struct {
	ID             int                  `gorm:"primary_key" json:"id"`
	RecipientsName string               `json:"recipientsName"`
	City           string               `json:"city"`
	Address        string               `json:"address"`
	StatusPayment  models.StatusPayment `json:"statusPayment"`
	PhoneNumber    string               `json:"phoneNumber"`
}

func GetAllOrderByUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var orders []models.Order
	db.Where("User_id=?", int(c.GetUint("currentUser"))).Find(&orders)

	response := utils.ApiResponse("All Orders by ID User", http.StatusOK, "success", &orders)
	c.JSON(http.StatusOK, response)
}

func MakeOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input OrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order := models.Order{
		RecipientsName: input.RecipientsName,
		OrderDate:      time.Now(),
		City:           input.City,
		Address:        input.Address,
		StatusPayment:  input.StatusPayment,
		PhoneNumber:    input.PhoneNumber,
		UserID:         int(c.GetUint("currentUser")),
	}
	db.Create(&order)
	response := utils.ApiResponse("Order success create ", http.StatusOK, "success", &order)
	c.JSON(http.StatusOK, response)
}

func UpdateOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var order models.Order
	if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		response := utils.ApiResponse("Order not found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if order.UserID != int(c.GetUint("currentUser")) {
		response := utils.ApiResponse("This User can,t access update Order", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	var input OrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderUpdate := models.Order{
		RecipientsName: input.RecipientsName,
		OrderDate:      time.Now(),
		City:           input.City,
		Address:        input.Address,
		StatusPayment:  input.StatusPayment,
		PhoneNumber:    input.PhoneNumber,
		UserID:         int(c.GetUint("currentUser")),
		UpdatedAt:      time.Now(),
	}
	db.Model(&order).Updates(orderUpdate)
	response := utils.ApiResponse("Order success update ", http.StatusOK, "success", orderUpdate)
	c.JSON(http.StatusOK, response)
}

func DeleteOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var order models.Order
	if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		response := utils.ApiResponse("Order Dont Delete", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	db.Delete(&order)
	response := utils.ApiResponse("Order had delete", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
