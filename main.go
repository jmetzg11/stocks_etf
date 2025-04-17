package main

import (
	"fmt"
	"log"
	"os"
	"stocks_etf/backend/routes"

	"stocks_etf/backend/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	log.SetOutput(logFile)

	log.Println("Go app started")
	fmt.Println("server started")

	_ = godotenv.Load()

	database.Connect()

	r := gin.Default()
	routes.SetupStaticRoutes(r)
	routes.SetupAPIRoutes(r)

	r.Run(":3000")
}
