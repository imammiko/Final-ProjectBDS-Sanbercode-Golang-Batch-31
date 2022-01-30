package controllers

import (
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/models"
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GeCateygoriesById a Rating godoc
// @Summary Get Category By User Id
// @Description Get list Category refrence By userID
// @Tags Category
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Category
// @Router /categories [get]
func GetAllCategoryByUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var categories []models.Category
	db.Where("user_id = ?", c.GetUint("currentUser")).Preload("Products").Find(&categories)
	response := utils.ApiResponse("All Category", http.StatusOK, "success", &categories)
	c.JSON(http.StatusOK, response)
}

// CreateCategory godoc
// @Summary Create New Category.
// @Description Creating a new Category.
// @Tags Category
// @Param Body body CategoryInput true "the body to create a new Category"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Category.
// @Router /categories [post]
func CreateCategories(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	categories := models.Category{
		Name:        input.Name,
		Description: input.Description,
		UserID:      int(c.GetUint("currentUser")),
	}
	db.Preload("Products").Create(&categories)
	response := utils.ApiResponse("categories success create ", http.StatusOK, "success", &categories)
	c.JSON(http.StatusOK, response)
}

// UpdateCategory godoc
// @Summary Update Category.
// @Description Update Category by id.
// @Tags Category
// @Param Body body CategoryInput true "the body to update a new Category"
// @Produce json
// @Param id path string true "Category id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Category
// @Router /categories/{id} [patch]
func UpdateCategories(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var categories models.Category
	var input CategoryInput
	if err := db.Where("id = ?", c.Param("id")).Preload("Products").Find(&categories).First(&categories).Error; err != nil {
		response := utils.ApiResponse("Categories not found", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	if categories.UserID != int(c.GetUint("currentUser")) {
		response := utils.ApiResponse("This User can,t access update category", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	UpdateCategories := models.Category{
		Name:        input.Name,
		Description: input.Description,
		UserID:      int(c.GetUint("currentUser")),
	}
	db.Model(&categories).Updates(UpdateCategories)
	response := utils.ApiResponse("category success update ", http.StatusOK, "success", categories)
	c.JSON(http.StatusOK, response)
}

// DeleteCategory godoc
// @Summary Delete one Category.
// @Description Delete a category by id.
// @Tags Category
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path string true "Category id"
// @Success 200 {object} map[string]boolean
// @Router /categories/{id} [delete]
func DeleteCategories(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Categories models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&Categories).Error; err != nil {
		response := utils.ApiResponse("Category Dont Delete", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	db.Delete(&Categories)
	response := utils.ApiResponse("Category had delete", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
