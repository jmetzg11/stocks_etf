package models

import "time"

type Transaction struct {
	ID             uint `gorm:"primaryKey"`
	Stock          string
	CreatedAt      time.Time
	Percent        float64
	Value          float64
	MarketValue    float64
	UnrealizedPLPC float64
}

type Value struct {
	ID        uint   `gorm:"primaryKey"`
	Stock     string `gorm:"uniqueIndex"`
	CreatedAt time.Time
	Value     float64
}
