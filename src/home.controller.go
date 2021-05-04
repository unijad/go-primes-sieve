package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeController(c *gin.Context) {
	c.HTML(http.StatusOK, "/views/home.html", gin.H{
		"Foo": "World",
	})
}