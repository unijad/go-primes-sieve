package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Prime struct {
	Limit int `json:"limit" binding:"required"`
}

func HighestPrimeNumber(c *gin.Context) {
	var json Prime

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"highestPrime": nil,
	})
}