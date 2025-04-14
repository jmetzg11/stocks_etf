package database

import (
	"errors"
	"log"
	"math"
	"stocks_etf/backend/models"

	"gorm.io/gorm"
)

func GetTotalReserves() float64 {
	var values []models.Value
	var total float64 = 0.0

	result := DB.Find(&values)
	if result.Error != nil {
		log.Printf("Failed to query values: %v", result.Error)
		return total
	}
	log.Println("result", result)

	for _, value := range values {
		total += value.Value
	}

	return total
}

func UpdateReserves(reserves, balance float64) map[string]float64 {
	realBalance := balance - reserves
	availableFunds := Min(55, .05*realBalance)

	amountEachETF := availableFunds / 11
	amountEachETF = math.Round(amountEachETF*100) / 100

	etfValues := make(map[string]float64)

	for _, etf := range EtfList {
		// Attempt to find existing record
		log.Println(etf)
		var value models.Value
		result := DB.Where("stock = ?", etf).First(&value)

		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				// Create new record if not found
				newValue := models.Value{
					Stock: etf,
					Value: amountEachETF,
				}
				if err := DB.Create(&newValue).Error; err != nil {
					log.Printf("Failed to create value for stock %s: %v", etf, err)
					continue
				}
				etfValues[etf] = amountEachETF
			} else {
				log.Printf("Failed to query value for stock %s: %v", etf, result.Error)
				continue
			}
		} else {
			// Update existing record
			value.Value += amountEachETF
			if err := DB.Save(&value).Error; err != nil {
				log.Printf("Failed to update value for stock %s: %v", etf, err)
				continue
			}
			etfValues[etf] = value.Value
		}
	}

	return etfValues
}

func UpdateDatabase(etf string, amountSpent, percentChange, unrealizedPlpc, marketValue float64) {
	// Update stock value
	if err := DB.Model(&models.Value{}).
		Where("stock = ?", etf).
		Update("value", gorm.Expr("value - ?", amountSpent)).Error; err != nil {
		log.Printf("Failed to update value for stock %s: %v", etf, err)
	}

	// Create new transaction record
	transaction := models.Transaction{
		Stock:          etf,
		Percent:        percentChange,
		Value:          amountSpent,
		MarketValue:    marketValue,
		UnrealizedPLPC: unrealizedPlpc,
	}

	if err := DB.Create(&transaction).Error; err != nil {
		log.Printf("Failed to update stock_transactions in UpdateDatabase etf: %s, %v", etf, err)
	}
}

func Min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

var EtfList = []string{"IXJ", "RXI", "KXI", "IXG", "EXI", "IXN", "IXC", "MXI", "JXI", "REET", "IXP"}
