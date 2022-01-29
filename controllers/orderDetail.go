package controllers

import (
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/models"
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderDetailInput struct {
	Jumlah    int `json:"jumlah"`
	ProductID int `json:"productID"`
	OrderID   int `json:"orderID"`
}

func CreateOrderDetail(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input OrderDetailInput
	var Product models.Product
	var Order models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.ProductID).First(&Product).Error; err != nil {
		response := utils.ApiResponse("ID product not found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	if err := db.Where("id = ?", input.OrderID).First(&Order).Error; err != nil {
		response := utils.ApiResponse("ID order not found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	harga := input.Jumlah * Product.Price
	OrderDetail := models.OrderDetail{
		Harga:     harga,
		Jumlah:    input.Jumlah,
		ProductID: input.ProductID,
		OrderID:   input.OrderID,
	}
	db.Create(&OrderDetail)
	response := utils.ApiResponse("Order Detail success create ", http.StatusOK, "success", OrderDetail)
	c.JSON(http.StatusOK, response)
}

func UpdateOrderDetail(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input OrderDetailInput
	var Product models.Product
	var OrderDetail models.OrderDetail
	var Order models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Where("id = ?", c.Param("id")).First(&OrderDetail).Error; err != nil {
		response := utils.ApiResponse("Product not found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	if err := db.Where("id = ?", input.ProductID).First(&Product).Error; err != nil {
		response := utils.ApiResponse("ID product not found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	if err := db.Where("id = ?", input.OrderID).First(&Order).Error; err != nil {
		response := utils.ApiResponse("ID order not found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	harga := input.Jumlah * Product.Price
	OrderDetailUpdate := models.OrderDetail{
		Harga:     harga,
		Jumlah:    input.Jumlah,
		ProductID: input.ProductID,
		OrderID:   input.OrderID,
	}
	db.Model(&OrderDetail).Updates(OrderDetailUpdate)
	response := utils.ApiResponse("Order Detail success Update ", http.StatusOK, "success", OrderDetail)
	c.JSON(http.StatusOK, response)
}

func DeleteOrderDetail(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var OrderDetail models.OrderDetail
	if err := db.Where("id = ?", c.Param("id")).First(&OrderDetail).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&OrderDetail)
	response := utils.ApiResponse("Order Detail had delete", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
