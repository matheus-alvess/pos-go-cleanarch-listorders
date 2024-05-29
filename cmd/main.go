package main

import (
	"encoding/json"
	grpcLib "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"pos-go-cleanarch-listorders/graphql"
	"pos-go-cleanarch-listorders/grpc/pb"
	"pos-go-cleanarch-listorders/grpc/service"
	"pos-go-cleanarch-listorders/internal/model"
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
		w.Header().Set("Content-Type", "application/json")
		orders, err := orderRepo.ListOrders()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(orders)
	}).Methods("GET")

	r.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var order model.Order
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := orderRepo.CreateOrder(order.Price(), order.Tax())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		order.IDField = int32(id)
		json.NewEncoder(w).Encode(order)
	}).Methods("POST")

	schema := graphqlLib.MustParseSchema(graphql.Schema, graphql.NewResolver(orderRepo))
	r.Handle("/graphql", &relay.Handler{Schema: schema}).Methods("POST")

	go func() {
		log.Println("HTTP server (REST & GraphQL) listening on port 8080")
		log.Fatal(http.ListenAndServe(":8080", r))
	}()

	grpcServer := grpcLib.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, service.NewOrderServiceServer(orderRepo))

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("gRPC server listening on port 50051")
	log.Fatal(grpcServer.Serve(lis))
}
