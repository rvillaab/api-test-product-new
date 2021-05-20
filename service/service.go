package service

import (
	data "api-test-product-new/data"
	ent "api-test-product-new/entity"
	"errors"
	"fmt"
)

type ProductService interface {
	CreateProduct(ent.Product) (ent.Product, error)
	GetAllProducts() ([]ent.Product, error)
	Count() int
	UpdateProduct(string, ent.Product) (interface{}, error)
	DeleteProduct(string) (string, error)
}

// stringService is a concrete implementation of StringService
type ProductServiceImpl struct{}

func NewService() ProductService {
	return ProductServiceImpl{}
}

func (prod ProductServiceImpl) CreateProduct(product ent.Product) (ent.Product, error) {

	data.Products = append(data.Products, product)
	return product, nil
}

func (ProductServiceImpl) Count() int {
	return len(data.Products)

}

func (prod ProductServiceImpl) GetAllProducts() ([]ent.Product, error) {

	return data.Products, nil
}

func (ProductServiceImpl) UpdateProduct(productID string, updatedProduct ent.Product) (interface{}, error) {

	for i, singleProduct := range data.Products {
		if singleProduct.ID == productID {
			singleProduct.ID = updatedProduct.ID
			singleProduct.Code = updatedProduct.Code
			singleProduct.Name = updatedProduct.Name
			singleProduct.Price = updatedProduct.Price
			data.Products = append(data.Products[:i], singleProduct)
			return updatedProduct, nil
		}
	}

	return nil, errors.New("Product not found")

}

func (ProductServiceImpl) DeleteProduct(productID string) (string, error) {

	for i, singleProduct := range data.Products {
		if singleProduct.ID == productID {
			data.Products = append(data.Products[:i], data.Products[i+1:]...)
			return fmt.Sprintf("The product with ID %v has been deleted successfully", productID), nil
		}
	}

	return "", errors.New("Product not found")
}
