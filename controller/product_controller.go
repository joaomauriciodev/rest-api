package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaomauriciodev/rest-api/usecase"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{productUsecase: usecase}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUsecase.GetProducts()
	if err != nil {
		log.Fatalln(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}
