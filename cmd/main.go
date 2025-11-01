package main

import (
	"log"
	"net"
	"net/http"

	"github.com/sorroche-m/desafio-clean-architecture/internal/config"
	"github.com/sorroche-m/desafio-clean-architecture/internal/delivery/graphql"
	"github.com/sorroche-m/desafio-clean-architecture/internal/delivery/grpc"
	httpHandler "github.com/sorroche-m/desafio-clean-architecture/internal/delivery/http"
	"github.com/sorroche-m/desafio-clean-architecture/internal/repository"
	"github.com/sorroche-m/desafio-clean-architecture/internal/usecase"
	pb "github.com/sorroche-m/desafio-clean-architecture/pkg/proto"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
	grpcServer "google.golang.org/grpc"
)

func main() {
	db, err := config.NewDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	orderRepo := repository.NewOrderRepository(db)
	orderUseCase := usecase.NewOrderUseCase(orderRepo)

	// gRPC configuration
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatal("Failed to listen:", err)
		}

		s := grpcServer.NewServer()
		pb.RegisterOrderServiceServer(s, grpc.NewOrderGRPCService(orderUseCase))

		log.Println("gRPC server listening on :50051")
		if err := s.Serve(lis); err != nil {
			log.Fatal("Failed to serve:", err)
		}
	}()

	r := gin.Default()

	orderHandler := httpHandler.NewOrderHandler(orderUseCase)

	// REST endpoints
	r.POST("/orders", orderHandler.CreateOrder)
	r.GET("/orders", orderHandler.ListOrders)
	r.GET("/orders/:id", orderHandler.GetOrder)

	// GraphQL endpoint
	schema, err := graphql.NewSchema(orderUseCase)
	if err != nil {
		log.Fatal("Failed to create GraphQL schema:", err)
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	r.POST("/graphql", gin.WrapH(h))
	r.GET("/graphql", gin.WrapH(h))

	log.Println("HTTP server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
