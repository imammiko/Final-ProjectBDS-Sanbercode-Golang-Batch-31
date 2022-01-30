package controllers

import (
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/models"
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CartInput struct {
	Total     int `json:"total"`
	ProductID int `json:"productId"`
}

func GetAllCartByUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var cart []models.Cart
	db.Where("Users_id=?", int(c.GetUint("currentUser"))).Find(&cart)

	response := utils.ApiResponse("All Products", http.StatusOK, "success", &cart)
	c.JSON(http.StatusOK, response)
}

func CreateCart(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input CartInput
	var Product models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Where("id = ?", input.ProductID).First(&Product).Error; err != nil {
		response := utils.ApiResponse("Category Product not found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	cart := models.Cart{
		Price:     Product.Price * input.Total,
		Total:     input.Total,
		Date:      time.Now(),
		UserID:    int(c.GetUint("currentUser")),
		ProductID: input.ProductID,
	}
	db.Create(&cart)
	response := utils.ApiResponse("Cart success create ", http.StatusOK, "success", &cart)
	c.JSON(http.StatusOK, response)
}
func UpdateCart(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var cart models.Cart
	var product models.Product
	if err := db.Where("id = ?", c.Param("id")).First(&cart).Error; err != nil {
		response := utils.ApiResponse("cart not found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var input CartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Where("id = ?", input.ProductID).First(&product).Error; err != nil {
		response := utils.ApiResponse("ID Product not Found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	if product.UserID != int(c.GetUint("currentUser")) {
		response := utils.ApiResponse("This user can,t access update", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	updateCart := models.Cart{
		Price:     product.Price * input.Total,
		Total:     input.Total,
		ProductID: input.ProductID,
		UserID:    int(c.GetUint("currentUser")),
	}
	db.Model(&cart).Updates(updateCart)
	response := utils.ApiResponse("cart success update ", http.StatusOK, "success", cart)
	c.JSON(http.StatusOK, response)
}

func DeleteCart(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var cart models.Cart
	if err := db.Where("id = ?", c.Param("id")).First(&cart).Error; err != nil {
		response := utils.ApiResponse("Product Dont Delete", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	db.Delete(&cart)
	response := utils.ApiResponse("product had delete", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func CartAddToOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	orderId := c.Param("orderId")
	chartId := c.Param("cartId")
	var order models.Order
	var cart models.Cart

	if err := db.Where("id = ?", orderId).First(&order).Error; err != nil {
		response := utils.ApiResponse("Order Id not found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	if err := db.Where("id = ?", chartId).First(&cart).Error; err != nil {
		response := utils.ApiResponse("cart Id not found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	if cart.UserID != int(c.GetUint("currentUser")) {
		response := utils.ApiResponse("This User can,t access cart to order", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	num, _ := strconv.Atoi(orderId)
	orderInput := models.OrderDetail{
		Harga:     cart.Price,
		Jumlah:    cart.Total,
		ProductID: cart.ProductID,
		OrderID:   num,
	}
	db.Create(&orderInput)
	db.Delete(&cart)
	response := utils.ApiResponse("Order Detail success create ", http.StatusOK, "success", orderInput)
	c.JSON(http.StatusOK, response)

}
