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

	emailRepo := repositories.NewEmailRepository(db)
	grpcServer := grpc.NewServer()

	// 4. Register your Service Implementation
	// This connects your 'routeGuideServer' or 'emailServer' to the gRPC engine
	emailImpl := &service.EmailServer{Repo: emailRepo} 
	email.RegisterEmailServiceServer(grpcServer, emailImpl)

	log.Printf("gRPC server listening at %v", lis.Addr())

	// 5. Start serving (This is a blocking call)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}