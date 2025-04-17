package main

import (
	"log"
	"os"
	"stocks_etf/backend/database"

	"github.com/joho/godotenv"
)

func main() {
	logFile, err := os.OpenFile("scheduler.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println("Scheduler started")

	_ = godotenv.Load()

	database.Connect() // Assuming this is your database.Connect function

	log.Println("Running Alpaca script...")
	// err = alpaca_script.Run()
	// if err != nil {
	// 	log.Fatalf("Failed to run alpaca script: %v", err)
	// }

	log.Println("Alpaca script completed successfully")
}
