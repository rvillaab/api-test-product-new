package server

import (
	e "api-test-product-new/endpoint"
	t "api-test-product-new/transport"
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer is a good little server
func NewHTTPServer(ctx context.Context, endpoints e.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("GET").Path("/count").Handler(httptransport.NewServer(
		endpoints.CountEndpoint,
		t.DecodeCountRequest,
		t.EncodeResponse,
	))

	r.Methods("GET").Path("/products").Handler(httptransport.NewServer(
		endpoints.GetAllProductsEndpoint,
		t.DecodeCountRequest,
		t.EncodeResponse,
	))

	r.Methods("DELETE").Path("/products/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteProductendpoint,
		t.DecodeDeleteRequest,
		t.EncodeResponse,
	))

	r.Methods("PUT").Path("/products/{id}").Handler(httptransport.NewServer(
		endpoints.UpdateProductEndpoint,
		t.DecodeUpdateRequest,
		t.EncodeResponse,
	))

	r.Methods("POST").Path("/product").Handler(httptransport.NewServer(
		endpoints.CreateProductEndpoint,
		t.DecodeProductCreateRequest,
		t.EncodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
