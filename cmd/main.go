package main

import (
	"encoding/json"
	grpcLib "google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"pos-go-cleanarch-listorders/graphql"
	"pos-go-cleanarch-listorders/grpc/pb"
	"pos-go-cleanarch-listorders/grpc/service"
	"pos-go-cleanarch-listorders/internal/repository"

	"github.com/gorilla/mux"
	graphqlLib "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	r := mux.NewRouter()

	db, err := repository.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	orderRepo := repository.NewOrderRepository(db)

	r.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		orders, err := orderRepo.ListOrders()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(orders)
	}).Methods("GET")

	schema := graphqlLib.MustParseSchema(graphql.Schema, graphql.NewResolver(orderRepo))
	r.Handle("/graphql", &relay.Handler{Schema: schema})

	go func() {
		log.Println("HTTP server listening on port 8080")
		log.Fatal(http.ListenAndServe(":8080", r))
	}()

	grpcServer := grpcLib.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, service.NewOrderServiceServer(orderRepo))

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("gRPC server listening on port 50051")
	log.Fatal(grpcServer.Serve(lis))
}
