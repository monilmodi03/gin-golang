package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Gin is working",
		})
	})
	server.Run(":8000") // listen and serve on localhost:8080/test
}
