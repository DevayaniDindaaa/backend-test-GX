package main

import (
	"log"
	"os"

	"github.com/DevayaniDindaaa/backend-test-GX/db"
	"github.com/DevayaniDindaaa/backend-test-GX/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}

	// initialize database
	db.Init()

	// initialize Fiber app
	app := fiber.New()

	// setup routes
	router.SetupRoutes(app)

	// start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
