package main

import (
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/config"
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/docs"
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/routes"
	"log"

	"github.com/joho/godotenv"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/
func main() {
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server ecommerce."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	// config.ConnectDataBase()
	// userRepository := user.NewReposiotry(db)
	// user, err := userRepository.Save(models.User{
	// 	Username:    "imamMiko1",
	// 	Email:       "imammiko@gmail.com",
	// 	Name:        "imam Sujatmiko",
	// 	Password:    "1234",
	// 	DateOfBirth: "2/desember/1995",
	// 	Gender:      "man",
	// 	PhoneNumber: "08955333",
	// })
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.SetupRouter(db)
	r.Run()
}
