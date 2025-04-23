package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaomauriciodev/rest-api/model"
)

type productController struct {
}

func NewProductController() productController {
	return productController{}
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
