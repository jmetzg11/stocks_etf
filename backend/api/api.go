package api

import (
	"net/http"
	"stocks_etf/backend/models"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetData(c *gin.Context) {
	var transactions []models.Transaction
	var values []models.Value

	threeMonthsAgo := time.Now().AddDate(0, -3, 0)

	if err := h.db.Where("created_at > ?", threeMonthsAgo).
		Order("created_at DESC").
		Find(&transactions).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
	}

	if err := h.db.Find(&values).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch values"})
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions, "values": values})
}
