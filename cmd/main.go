package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joaomauriciodev/rest-api/controller"
	"github.com/joaomauriciodev/rest-api/usecase"
)

func main() {
	s := gin.Default()

	productUsecase := usecase.NewProductUsecase()
	productController := controller.NewProductController(productUsecase)

	s.GET("/api/products", productController.GetProducts)

	s.Run(":8000")
}
