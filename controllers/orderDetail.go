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

// CreateOrderDetail godoc
// @Summary Create New OrderDetail.
// @Description Creating a new OrderDetail.
// @Tags OrderDetail
// @Param Body body OrderDetailInput true "the body to create a new OrderDetail"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.OrderDetail.
// @Router /orderDetails [post]
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

// UpdateOrderDetail godoc
// @Summary Update OrderDetail.
// @Description Update OrderDetail by id.
// @Tags OrderDetail
// @Param Body body OrderDetailInput true "the body to update a new OrderDetail"
// @Produce json
// @Param id path string true "OrderDetail id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.OrderDetail
// @Router /orderDetails/{id} [patch]
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

// DeleteOrderDetail godoc
// @Summary Delete one OrderDetail.
// @Description Delete a OrderDetail by id.
// @Tags OrderDetail
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path string true "OrderDetail id"
// @Success 200 {object} map[string]boolean
// @Router /orderDetails/{id} [delete]
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
