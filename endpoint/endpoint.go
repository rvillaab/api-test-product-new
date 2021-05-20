package endpoint

import (
	ent "api-test-product-new/entity"
	s "api-test-product-new/service"
	t "api-test-product-new/transport"
	"context"
	"log"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints are exposed
type Endpoints struct {
	CreateProductEndpoint  endpoint.Endpoint
	GetAllProductsEndpoint endpoint.Endpoint
	CountEndpoint          endpoint.Endpoint
	UpdateProductEndpoint  endpoint.Endpoint
	DeleteProductendpoint  endpoint.Endpoint
}

func MakeCountEndpoint(svc s.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		v := svc.Count()
		return v, nil
	}
}

func MakeCreateProductEndpoint(svc s.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(ent.Product)
		v, err := svc.CreateProduct(req)
		if err != nil {

			return t.ProductResponse{Str: "", Err: err.Error()}, nil
		}
		return v, nil
	}
}

func MakeGetallProductsEndpoint(svc s.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		v, err := svc.GetAllProducts()
		if err != nil {
			log.Fatal(err)
			return v, err
		}
		return v, nil
	}
}

func MakeUpdateProductEndpoint(svc s.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(t.ProductUpdateRequest)
		v, err := svc.UpdateProduct(req.V, req.S)
		if err != nil {
			return t.ProductResponse{Str: "", Err: err.Error()}, nil
		}
		return v, nil
	}
}

func MakeDeleteroductEndpoint(svc s.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(t.ProductRequest)
		v, err := svc.DeleteProduct(req.Str)
		if err != nil {
			return t.ProductResponse{Str: "", Err: err.Error()}, nil
		}
		return v, nil
	}
}

func (e Endpoints) Count(ctx context.Context) (string, error) {

	resp, err := e.CountEndpoint(ctx, nil)
	if err != nil {
		return "", err
	}

	response := resp.(string)
	return response, nil
}

func (e Endpoints) GetAllProducts(ctx context.Context) (string, error) {

	resp, err := e.GetAllProductsEndpoint(ctx, nil)
	if err != nil {
		return "", err
	}

	response := resp.(string)
	return response, nil
}

func (e Endpoints) CreateProduct(ctx context.Context) (string, error) {

	resp, err := e.CreateProductEndpoint(ctx, nil)
	if err != nil {
		return "", err
	}

	response := resp.(string)
	return response, nil
}

func (e Endpoints) UpdateProduct(ctx context.Context) (string, error) {

	resp, err := e.UpdateProductEndpoint(ctx, nil)
	if err != nil {
		return "", err
	}

	response := resp.(string)
	return response, nil
}

func (e Endpoints) DeleteProductProduct(ctx context.Context) (string, error) {

	resp, err := e.DeleteProductendpoint(ctx, nil)
	if err != nil {
		return "", err
	}

	response := resp.(string)
	return response, nil
}
