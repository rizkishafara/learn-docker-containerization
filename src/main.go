package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"

	"src/config"
	"src/models"
	"src/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	} 

	engine := html.New("./views", ".html")

	// cek jika enviroment dev
	if os.Getenv("ENVIRONMENT") == "development" || os.Getenv("ENV") == "development" {
		engine.Reload(true)
		log.Println("Running in development mode: template reload enabled")
	}

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	config.InitDB()

	config.DB.AutoMigrate(&models.Users{})
	config.DB.AutoMigrate(&models.Products{})

	// app.Static("/public", "./public")

	routes.MainRoute(app)
	routes.Users(app)
	routes.Products(app)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
