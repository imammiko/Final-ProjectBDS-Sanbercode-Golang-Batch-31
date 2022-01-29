package routes

import (
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/controllers"
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/middleware"
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/user"
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/utils/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	userRepository := user.NewReposiotry(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userController := controllers.NewUserController(userService, authService)
	middleware := middleware.NewUserMiddleware(authService)
	// middleware := middleware.NewUserMiddleware(authService)
	r.POST("/register", userController.RegisterUser)
	r.POST("/login", userController.Login)
	r.GET("/forgotPassword", userController.ForgetPassword)
	r.POST("/changePassword", userController.ChangePassword)

	productMiddlewareRoute := r.Group("/products")
	r.GET("/allProducts", controllers.GetAllProduct)
	productMiddlewareRoute.Use(middleware.JwtAuthMiddleware())
	productMiddlewareRoute.GET("", controllers.GetProductsById)
	productMiddlewareRoute.POST("", controllers.CreateProduct)
	productMiddlewareRoute.PATCH("/:id", controllers.UpdateProduct)
	productMiddlewareRoute.DELETE("/:id", controllers.DeleteProduct)

	orderMiddlewareRoute := r.Group("/orders")
	orderMiddlewareRoute.Use(middleware.JwtAuthMiddleware())
	orderMiddlewareRoute.GET("", controllers.GetAllOrderByUser)
	orderMiddlewareRoute.POST("", controllers.MakeOrder)
	orderMiddlewareRoute.PATCH("/:id", controllers.UpdateOrder)
	orderMiddlewareRoute.DELETE("/:id", controllers.DeleteOrder)
	return r
}
