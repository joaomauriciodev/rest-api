package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaomauriciodev/rest-api/model"
	"github.com/joaomauriciodev/rest-api/usecase"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{productUsecase: usecase}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID:          1,
			Description: "Produto X",
			Price:       10.00,
		},
	}

	ctx.JSON(http.StatusOK, products)
}
