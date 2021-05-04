package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeController(c *gin.Context) {
	c.String(http.StatusOK, "Homepage")
}