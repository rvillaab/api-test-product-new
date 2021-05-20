package data

import (
	"api-test-product-new/entity"
)

type repo struct{}

type allProducts []entity.Product

var Products = allProducts{
	{
		ID:    "1",
		Code:  "HPOT",
		Name:  "Harry Potter Book 1",
		Price: 15.0,
	},
}

func NewMemoryRepository() Repository {
	return &repo{}
}

func (r *repo) FindAll() (*[]entity.Product, error) {

	prods := []entity.Product{}

	return &prods, nil
}

func (r *repo) Find(id string) (*entity.Product, error) {

	return &Products[0], nil
}
