package main

import (
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/config"
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/docs"
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/routes"
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/utils"
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
	//programmatically set swagger info
	err := godotenv.Load(".env")
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Movie."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = utils.Getenv("SWAGGER_HOST", "localhost:8080")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.SetupRouter(db)
	r.Run()
}
