package database

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"stocks_etf/backend/models"
	"strconv"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("data/data.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	DB.AutoMigrate(&models.Transaction{}, &models.Value{})

	seedDB(DB)
	return nil
}

func seedDB(db *gorm.DB) {
	var count int64
	db.Model(&models.Transaction{}).Count(&count)
	if count > 0 {
		log.Println("Database already has data, skipping seed")
		return
	}

	_, filename, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(filename)
	transactionsFile, err := os.Open(filepath.Join(basepath, "stock_transactions.csv"))
	if err != nil {
		log.Println("Error opening stock_transactions.csv:", err)
		return
	}
	defer transactionsFile.Close()

	transReader := csv.NewReader(transactionsFile)
	// Skip header
	_, err = transReader.Read()
	if err != nil {
		log.Println("Error reading CSV header:", err)
		return
	}

	var transactions []models.Transaction
	for {
		record, err := transReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error reading transaction record:", err)
			continue
		}

		id, _ := strconv.Atoi(record[0])
		createdAt, _ := time.Parse("2006-01-02", record[2])
		percent, _ := strconv.ParseFloat(record[3], 64)
		value, _ := strconv.ParseFloat(record[4], 64)
		marketValue, _ := strconv.ParseFloat(record[5], 64)
		unrealizedPLPC, _ := strconv.ParseFloat(record[6], 64)

		transaction := models.Transaction{
			ID:             id,
			Stock:          record[1],
			CreatedAt:      createdAt,
			Percent:        percent,
			Value:          value,
			MarketValue:    marketValue,
			UnrealizedPLPC: unrealizedPLPC,
		}
		transactions = append(transactions, transaction)
	}

	if len(transactions) > 0 {
		result := db.Create(&transactions)
		if result.Error != nil {
			log.Println("Error inserting transactions:", result.Error)
		} else {
			log.Printf("Inserted %d transactions\n", len(transactions))
		}
	}

	valuesFile, err := os.Open(filepath.Join(basepath, "stock_values.csv"))
	if err != nil {
		log.Println("Error opening stock_values.csv", err)
		return
	}
	defer valuesFile.Close()

	valueReader := csv.NewReader(valuesFile)
	// Skip header
	_, err = valueReader.Read()
	if err != nil {
		log.Println("Error reading CSV header:", err)
		return
	}

	var values []models.Value
	for {
		record, err := valueReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error reading value record:", err)
			continue
		}

		createdAt, _ := time.Parse("2006-01-02", record[1])
		value, _ := strconv.ParseFloat(record[2], 64)

		stockValue := models.Value{
			Stock:     record[0],
			CreatedAt: createdAt,
			Value:     value,
		}

		values = append(values, stockValue)
	}
	if len(values) > 0 {
		result := db.Create(&values)
		if result.Error != nil {
			log.Println("Error insertin stock values:", result.Error)
		} else {
			log.Printf("Inserted %d stock values\n", len(values))
		}
	}
}
