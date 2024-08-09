package main

import (
	"log"
	"os"
	"github.com/Hailemari/task_manager/data"
	"github.com/Hailemari/task_manager/router"
	"github.com/joho/godotenv"

)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize the database connection
	mongoURI := os.Getenv("MONGODB_URI")
	err = data.ConnectDB(mongoURI)  
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Set up the router and start the server
	r := router.SetupRouter()
	r.Run(":8000") 
}
