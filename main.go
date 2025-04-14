package main

import (
	"fmt"
	"log"
	"os"
	"stocks_etf/backend/alpaca_script"
	"stocks_etf/backend/routes"
	"time"

	"stocks_etf/backend/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
)

func main() {

	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	log.SetOutput(logFile)

	log.Println("Go app started")
	fmt.Println("server started")

	err = godotenv.Load()
	if err != nil {
		log.Print("No .env file found, will use environment variables")
	}

	database.Connect()

	r := gin.Default()
	routes.SetupStaticRoutes(r)
	routes.SetupAPIRoutes(r)

	currentUTC := time.Now()
	log.Printf("Current time: %s", currentUTC.Format("2006-01-02 15:04:05"))

	c := cron.New()
	// second, minute, hour, day, month, day of week
	// "0 0 17 * * 1-5" noon NY monday through friday
	err = c.AddFunc("0 10 12 * * 1-5", func() {
		log.Println("Cron job was triggered")
		err := alpaca_script.Run()
		if err != nil {
			log.Printf("Failed to run alpaca script: %v", err)
		}
	})
	if err != nil {
		log.Printf("Failed to schedule cron job: %v", err)
	}
	c.Start()

	r.Run(":3000")
}
