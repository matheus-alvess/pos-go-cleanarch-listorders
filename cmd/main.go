package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	pb "order-app/grpc"
	"pos-go-cleanarch-listorders/internal/repository"

	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"google.golang.org/grpc"
)

func main() {
	r := mux.NewRouter()

	db, err := repository.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	orderRepo := repository.NewOrderRepository(db)

	// Configurar o serviço REST
	orderService := grpcService.NewOrderServiceServer(orderRepo)
	r.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		orders, err := orderRepo.ListOrders()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(orders)
	}).Methods("GET")

	// Configurar servidor GraphQL
	schema := graphql.MustParseSchema(graphql.Schema, graphql.NewResolver(orderRepo))
	r.Handle("/graphql", &relay.Handler{Schema: schema})

	// Iniciar o servidor HTTP
	go func() {
		log.Println("HTTP server listening on port 8080")
		log.Fatal(http.ListenAndServe(":8080", r))
	}()

	// Iniciar o servidor gRPC
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, orderService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("gRPC server listening on port 50051")
	log.Fatal(grpcServer.Serve(lis))
}
