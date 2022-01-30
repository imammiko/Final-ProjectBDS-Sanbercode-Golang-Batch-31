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

// GetAllProduct godoc
// @Summary Get all Product.
// @Description Get a list of Products.
// @Tags Product
// @Produce json
// @Success 200 {object} []models.Product
// @Router /allProducts [get]
func GetAllProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var products []models.Product
	db.Find(&products)

	response := utils.ApiResponse("All Products", http.StatusOK, "success", &products)
	c.JSON(http.StatusOK, response)
}

// GetProductsByUser a Rating godoc
// @Summary Get Products By User Id
// @Description Get list products refrence By userID
// @Tags Product
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Product
// @Router /products [get]
func GetProductsByUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var products []models.Product
	db.Where("User_id=?", int(c.GetUint("currentUser"))).Find(&products)

	response := utils.ApiResponse("All Products", http.StatusOK, "success", &products)
	c.JSON(http.StatusOK, response)
}

// CreateProduct godoc
// @Summary Create New Product.
// @Description Creating a new Product.
// @Tags Product
// @Param Body body ProductInput true "the body to create a new Product"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Product.
// @Router /products [post]
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

// UpdateProduct godoc
// @Summary Update Product.
// @Description Update Product by id.
// @Tags Product
// @Param Body body ProductInput true "the body to update a new Product"
// @Produce json
// @Param id path string true "Product id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Product
// @Router /products/{id} [patch]
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

// DeleteProduct godoc
// @Summary Delete one Product.
// @Description Delete a Product by id.
// @Tags Product
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path string true "Product id"
// @Success 200 {object} map[string]boolean
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var product models.Product
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		response := utils.ApiResponse("Product Not Found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	db.Delete(&product)
	response := utils.ApiResponse("product had delete", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
