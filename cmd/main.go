package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joaomauriciodev/rest-api/controller"
	"github.com/joaomauriciodev/rest-api/db"
	"github.com/joaomauriciodev/rest-api/repository"
	"github.com/joaomauriciodev/rest-api/usecase"
)

func main() {

	conn, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	s := gin.Default()
	productRepository := repository.NewProductRepository(conn)
	productUsecase := usecase.NewProductUsecase(productRepository)
	productController := controller.NewProductController(productUsecase)

	s.GET("/api/products", productController.GetProducts)

	s.Run(":8000")
}
