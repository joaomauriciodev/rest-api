package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joaomauriciodev/rest-api/controller"
)

func main() {
	s := gin.Default()

	productController := controller.NewProductController()

	s.GET("/api/products", productController.GetProducts)

	s.Run(":8000")
}
