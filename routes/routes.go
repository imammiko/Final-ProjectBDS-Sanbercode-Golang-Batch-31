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
	productMiddlewareRoute.Use(middleware.JwtAuthMiddleware())
	productMiddlewareRoute.GET("/", controllers.TestId)
	return r
}
