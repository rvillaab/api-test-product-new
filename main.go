package main

import (
	"api-test-product-new/service"
	"api-test-product-new/transport"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	svc := service.ProductServiceImpl{}

	uppercaseHandler := httptransport.NewServer(
		transport.MakeCreateProductEndpoint(svc),
		transport.DecodeProductCreateRequest,
		transport.EncodeResponse,
	)

	countHandler := httptransport.NewServer(
		transport.MakeCountEndpoint(svc),
		transport.DecodeCountRequest,
		transport.EncodeResponse,
	)

	getAllHandler := httptransport.NewServer(
		transport.MakeGetallProductsEndpoint(svc),
		transport.DecodeCountRequest,
		transport.EncodeResponse,
	)

	updateHandler := httptransport.NewServer(
		transport.MakeUpdateProductEndpoint(svc),
		transport.DecodeUpdateRequest,
		transport.EncodeResponse,
	)

	deleteHandler := httptransport.NewServer(
		transport.MakeDeleteroductEndpoint(svc),
		transport.DecodeDeleteRequest,
		transport.EncodeResponse,
	)

	router.PathPrefix("/product").Handler(uppercaseHandler).Methods("POST")
	router.PathPrefix("/products/{id}").Handler(countHandler).Methods("GET")
	router.PathPrefix("/products/count").Handler(countHandler).Methods("GET")
	router.PathPrefix("/products").Handler(getAllHandler).Methods("GET")
	router.PathPrefix("/products/{id}").Handler(updateHandler).Methods("PUT")
	router.PathPrefix("/products/{id}").Handler(deleteHandler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))

}
