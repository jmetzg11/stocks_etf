package models

import "time"

type Transaction struct {
	ID             int
	Stock          string
	CreatedAt      time.Time
	Percent        float64
	Value          float64
	MarketValue    float64
	UnrealizedPLPC float64
}

type Value struct {
	Stock     string
	CreatedAt time.Time
	Value     float64
}
