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

// @Summary Delete item
// @Description Delete item
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "ID"
// @Success 204
// @Failure 400 {string} string "400 StatusBadRequest"
// @Failure 404 {string} string "404 not found"
// @Router /User/{id} [delete]
// if err != nil {
// 	errors := utils.FormatValidationEror(err)
// 	errorMessage := gin.H{"errors": errors}
// 	response := utils.ApiResponse("register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
// 	c.JSON(http.StatusUnprocessableEntity, response)
// 	return
// }

// response := utils.ApiResponse("Account has been registered", http.StatusOK, "success", formatter)
// c.JSON(http.StatusOK, response)

type ProductInput struct {
	Name        string `json:"name"`
	Condition   string `json:"condition"`
	Description string `json:"description"`
	ImageUrl    string `json:"ImageUrl"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	Heavy       string `json:"heavy"`
	CategoryID  int    `json:"categoryID"`
}

func GetAllProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var products []models.Product
	db.Find(&products)

	response := utils.ApiResponse("All Products", http.StatusOK, "success", &products)
	c.JSON(http.StatusOK, response)
}

func GetProductsById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var products []models.Product
	db.Where("User_id=?", int(c.GetUint("currentUser"))).Find(&products)

	response := utils.ApiResponse("All Products", http.StatusOK, "success", &products)
	c.JSON(http.StatusOK, response)
}

func CreateProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input ProductInput
	var categoryProduct models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Where("id = ?", input.CategoryID).First(&categoryProduct).Error; err != nil {
		response := utils.ApiResponse("Category Product not found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	_, err := url.ParseRequestURI(input.ImageUrl)
	if err != nil {
		errors := utils.FormatValidationEror(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("Create Product failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	product := models.Product{
		Name:        input.Name,
		Condition:   input.Condition,
		Description: input.Description,
		ImageUrl:    input.ImageUrl,
		Stock:       input.Stock,
		Price:       input.Price,
		Heavy:       input.Heavy,
		UserID:      int(c.GetUint("currentUser")),
		CategoryID:  input.CategoryID,
	}
	db.Create(&product)
	response := utils.ApiResponse("product success create ", http.StatusOK, "success", &product)
	c.JSON(http.StatusOK, response)
}

func UpdateProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var product models.Product
	var categoryProduct models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		response := utils.ApiResponse("Product not found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	if product.UserID != int(c.GetUint("currentUser")) {
		response := utils.ApiResponse("This User can,t access update Products", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	var input ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Where("id = ?", input.CategoryID).First(&categoryProduct).Error; err != nil {
		response := utils.ApiResponse("ID Category not Found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	updatedInput := models.Product{
		Name:        input.Name,
		Condition:   input.Condition,
		Description: input.Description,
		ImageUrl:    input.ImageUrl,
		Stock:       input.Stock,
		Price:       input.Price,
		Heavy:       input.Heavy,
		UserID:      int(c.GetUint("currentUser")),
		CategoryID:  input.CategoryID,
		UpdatedAt:   time.Now(),
	}
	db.Model(&product).Updates(updatedInput)
	response := utils.ApiResponse("product success update ", http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)
}

func DeleteProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var product models.Product
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		response := utils.ApiResponse("Product Dont Delete", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	db.Delete(&product)
	response := utils.ApiResponse("product had delete", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
