package main

import (
	"log"
	"net"

	email "github.com/TheDurodola/go-email-service/api/email"
	"github.com/TheDurodola/go-email-service/internal/config"
	"github.com/TheDurodola/go-email-service/internal/data/repositories"
	"github.com/TheDurodola/go-email-service/internal/service"
	"google.golang.org/grpc"
)

func main() {

	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	log.Println("Successfully connected to the database!")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rabbitMQ := config.NewRabbitMQConnection()
	defer rabbitMQ.Close()
	brevo := config.NewBrevoClient()

	emailRepo := repositories.NewEmailRepository(db)
	grpcServer := grpc.NewServer()


	emailImpl := &service.EmailServer{Repo: emailRepo}
	email.RegisterEmailServiceServer(grpcServer, emailImpl)

	log.Printf("gRPC server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
