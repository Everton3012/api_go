package repository

import (
	"api_go/model"
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return []model.Product{}, err
	}

	var productsList []model.Product
	var productObj model.Product
	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)
		if err != nil {
			fmt.Println(err.Error())
			return []model.Product{}, err
		}

		productsList = append(productsList, productObj)
	}
	rows.Close()
	return productsList, nil
}
