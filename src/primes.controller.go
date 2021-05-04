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
		"highestPrime": getHighestPrimeNumber(json.Limit),
	})
}

func getHighestPrimeNumber(n int) int {
	// make a slice array with n length
	// default value false when the index is a prime number
	var slice = make([]bool, n+1)
	// 0, 1, are not prime numbers
	// we will leave 2 as false, as it's our first prime number
	slice[0] = true
	slice[1] = true
	// we need to record last prime as our highest prime less than n
	var highestPrime int
	// loop slices from first prime which is 2
	for i := 2; i <= n; i++ {
		// as sieves algorithm check if the number is not checked as not prime first
		if slice[i] == false {
			// loop through multiple of our prime number
			for j := 2; i*j <= n; j++ {
				// mark the number as neglicted from our primes list
				slice[i*j] = true
			}
			// set the prime over iteration
			highestPrime = i
		}
	}
	// return highest prime number
	return highestPrime
}
