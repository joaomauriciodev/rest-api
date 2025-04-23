package main

import "github.com/gin-gonic/gin"

func main() {
	s := gin.Default()
	s.GET("/teste", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "teste gin",
		})
	})

	s.Run(":8000")
}
