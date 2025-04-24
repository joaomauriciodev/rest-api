package repository

import (
	"database/sql"
	"log"

	"github.com/joaomauriciodev/rest-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(conn *sql.DB) ProductRepository {
	return ProductRepository{connection: conn}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, description, price FROM products"

	rows, err := pr.connection.Query(query)
	if err != nil {
		log.Fatalln(err)
		return []model.Product{}, err
	}

	var productsList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Description,
			&productObj.Price,
		)

		if err != nil {
			log.Fatalln(err)
			return []model.Product{}, err
		}

		productsList = append(productsList, productObj)
	}

	return productsList, nil

}
