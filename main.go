package main

import (
	"api-test-product-new/endpoint"
	"api-test-product-new/server"
	"api-test-product-new/service"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()
	srv := service.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := endpoint.Endpoints{
		CountEndpoint:          endpoint.MakeCountEndpoint(srv),
		GetAllProductsEndpoint: endpoint.MakeGetallProductsEndpoint(srv),
		CreateProductEndpoint:  endpoint.MakeCreateProductEndpoint(srv),
		UpdateProductEndpoint:  endpoint.MakeUpdateProductEndpoint(srv),
		DeleteProductendpoint:  endpoint.MakeDeleteroductEndpoint(srv),
	}

	// HTTP transport
	go func() {
		log.Println("service is listening on port:", *httpAddr)
		handler := server.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
