package main

import "github.com/gin-gonic/gin"

func routerSetup() *gin.Engine {
	r := gin.Default()

	r.GET("/", HomeController)
	return r
}

func main() {
	r := routerSetup()
	r.Run(":8080")
}