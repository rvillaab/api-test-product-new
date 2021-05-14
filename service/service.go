package service

import (
	"errors"
	"fmt"
)

type ProductService interface {
	CreateProduct(Product) (Product, error)
	GetAllProducts() ([]Product, error)
	Count() int
	UpdateProduct(string, Product) (interface{}, error)
	DeleteProduct(string) (string, error)
}

// stringService is a concrete implementation of StringService
type ProductServiceImpl struct{}

type Product struct {
	ID    string  `json:"id"`
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type allProducts []Product

var products = allProducts{
	{
		ID:    "1",
		Code:  "HPOT",
		Name:  "Harry Potter Book 1",
		Price: 15.0,
	},
}

func (ProductServiceImpl) CreateProduct(product Product) (Product, error) {

	products = append(products, product)
	return product, nil
}

func (ProductServiceImpl) Count() int {
	return len(products)

}

func (ProductServiceImpl) GetAllProducts() ([]Product, error) {
	return products, nil
}

func (ProductServiceImpl) UpdateProduct(productID string, updatedProduct Product) (interface{}, error) {

	for i, singleProduct := range products {
		if singleProduct.ID == productID {
			singleProduct.ID = updatedProduct.ID
			singleProduct.Code = updatedProduct.Code
			singleProduct.Name = updatedProduct.Name
			singleProduct.Price = updatedProduct.Price
			products = append(products[:i], singleProduct)
			return updatedProduct, nil
		}
	}

	return nil, errors.New("Product not found")

}

func (ProductServiceImpl) DeleteProduct(productID string) (string, error) {

	for i, singleProduct := range products {
		if singleProduct.ID == productID {
			products = append(products[:i], products[i+1:]...)
			return fmt.Sprintf("The product with ID %v has been deleted successfully", productID), nil
		}
	}

	return "", errors.New("Product not found")
}
