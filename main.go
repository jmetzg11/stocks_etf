package main

import (
	"fmt"
	"log"
	"os"
	"stocks_etf/backend/alpaca_script"
	"stocks_etf/backend/routes"

	"stocks_etf/backend/database"

	"github.com/gin-gonic/gin"
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
	database.Connect()

	r := gin.Default()
	routes.SetupStaticRoutes(r)
	routes.SetupAPIRoutes(r)

	log.Println("Adding cron job to run at 14:10 ET, Monday to Friday")
	c := cron.New()
	// second, minute, hour, day, month, day of week
	// sceduled to run 11AM ET, Monday - Friday
	err = c.AddFunc("0 0 16 * * 1-5", func() {
		log.Println("Running scheduled alpaca script...")
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
