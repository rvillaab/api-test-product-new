package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	ent "api-test-product-new/entity"

	"github.com/gorilla/mux"
)

type ProductUpdateRequest struct {
	S ent.Product `json:"s"`
	V string      `json:"v"`
}

type ProductRequest struct {
	Str string `json:"str"`
}

type ProductResponse struct {
	Str string `json:"str"`
	Err string `json:"err,omitempty"`
}

type ProductsResponse struct {
	Str string
}

func DecodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func DecodeProductCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request ent.Product
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func DecodeUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request ProductUpdateRequest
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

func DecodeDeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request ProductRequest

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("id is missing in parameters")
	}

	request.Str = id
	return request, nil
}
