package usecase

import (
	"github.com/joaomauriciodev/rest-api/model"
	"github.com/joaomauriciodev/rest-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	products, err := pu.repository.GetProducts()
	if err != nil {
		return []model.Product{}, err
	}
	return products, nil
}
