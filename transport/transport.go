package transport

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	s "api-test-product-new/service"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

type productUpdateRequest struct {
	S s.Product `json:"s"`
	V string    `json:"v"`
}

type productRequest struct {
	Str string `json:"str"`
}

type productResponse struct {
	Str string `json:"str"`
	Err string `json:"err,omitempty"`
}

type productsResponse struct {
	Str string
}

func MakeCountEndpoint(svc s.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		v := svc.Count()
		return v, nil
	}
}

func DecodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func MakeCreateProductEndpoint(svc s.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(s.Product)
		v, err := svc.CreateProduct(req)
		if err != nil {

			return productResponse{"", err.Error()}, nil
		}
		return v, nil
	}
}

func DecodeProductCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request s.Product
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
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

func DecodeUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request productUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&request.S); err != nil {
		return nil, err
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("id is missing in parameters")
	}

	request.V = id
	return request, nil
}

func MakeUpdateProductEndpoint(svc s.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(productUpdateRequest)
		v, err := svc.UpdateProduct(req.V, req.S)
		if err != nil {
			return productResponse{"", err.Error()}, nil
		}
		return v, nil
	}
}

func DecodeDeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request productRequest

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("id is missing in parameters")
	}

	request.Str = id
	return request, nil
}

func MakeDeleteroductEndpoint(svc s.ProductService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(productRequest)
		v, err := svc.DeleteProduct(req.Str)
		if err != nil {
			return productResponse{"", err.Error()}, nil
		}
		return v, nil
	}
}
