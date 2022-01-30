package controllers

import (
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/models"
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RatingInput struct {
	Star        int    `json:"star"`
	Description string `json:"description"`
	ProductID   int    `json:"productId"`
}

func GetRatingByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var rating []models.Review
	db.Where("User_id=?", int(c.GetUint("currentUser"))).Find(&rating)

	response := utils.ApiResponse("All Review by User", http.StatusOK, "success", &rating)
	c.JSON(http.StatusOK, response)
}

func CreateRating(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input RatingInput
	var product models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.ProductID).First(&product).Error; err != nil {
		response := utils.ApiResponse(" Product not found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	if product.UserID == int(c.GetUint("currentUser")) {
		response := utils.ApiResponse("user Dont Access user's self", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	if input.Star > 5 {
		response := utils.ApiResponse("Star must 0-5 Dont Access", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	rating := models.Review{
		Star:        input.Star,
		Description: input.Description,
		ProductID:   input.ProductID,
		UserID:      int(c.GetUint("currentUser")),
	}

	db.Create(&rating)
	response := utils.ApiResponse("Review success create ", http.StatusOK, "success", &rating)
	c.JSON(http.StatusOK, response)
}

func UpdateRating(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var rating models.Review
	var product models.Product
	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		response := utils.ApiResponse("Review not found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	if rating.UserID != int(c.GetUint("currentUser")) {
		response := utils.ApiResponse("This User can,t access update Products", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	var input RatingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Where("id = ?", input.ProductID).First(&product).Error; err != nil {
		response := utils.ApiResponse("ID Product not Found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	if product.UserID == int(c.GetUint("currentUser")) {
		response := utils.ApiResponse("user Dont Access user's self", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	if input.Star > 5 {
		response := utils.ApiResponse("Star must 0-5 Dont Access", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	updateRating := models.Review{
		Star:        input.Star,
		Description: input.Description,
		UserID:      int(c.GetUint("currentUser")),
		ProductID:   input.ProductID,
	}
	db.Model(&rating).Updates(updateRating)
	response := utils.ApiResponse("product success update ", http.StatusOK, "success", rating)
	c.JSON(http.StatusOK, response)
}

func DeleteRating(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var rating models.Review
	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		response := utils.ApiResponse("Review Dont Delete", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	var product models.Product
	if err := db.Where("id = ?", rating.ProductID).First(&product).Error; err != nil {
		response := utils.ApiResponse("Review not found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	if product.UserID == int(c.GetUint("currentUser")) {
		response := utils.ApiResponse("user Dont Access user's self", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	db.Delete(&rating)
	response := utils.ApiResponse("Review had delete", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)

}
