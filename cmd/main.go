package main

import (
	grpcLib "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"pos-go-cleanarch-listorders/graphql"
	"pos-go-cleanarch-listorders/grpc/pb"
	grpcService "pos-go-cleanarch-listorders/grpc/service"
	"pos-go-cleanarch-listorders/internal/repository"
	"pos-go-cleanarch-listorders/internal/service"

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
	orderService := service.NewOrderService(orderRepo)

	r.HandleFunc("/order", orderService.ListOrders).Methods("GET")
	r.HandleFunc("/order", orderService.CreateOrder).Methods("POST")

	schema := graphqlLib.MustParseSchema(graphql.Schema, graphql.NewResolver(orderRepo))
	r.Handle("/graphql", &relay.Handler{Schema: schema}).Methods("POST")

	go func() {
		log.Println("HTTP server (REST & GraphQL) listening on port 8080")
		log.Fatal(http.ListenAndServe(":8080", r))
	}()

	grpcServer := grpcLib.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, grpcService.NewOrderServiceServer(orderRepo))

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("gRPC server listening on port 50051")
	log.Fatal(grpcServer.Serve(lis))
}
