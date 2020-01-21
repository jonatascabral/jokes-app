package main

import (
	"github.com/joho/godotenv"
	"github.com/jonatascabral/jokes-app/pkg/routes"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	router := routes.LoadRoutes()

	port := ":" + os.Getenv("APP_PORT")
	router.Run(port)
}