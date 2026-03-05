package main

import (
	"log"
	"log/slog"
	"net"

	"github.com/TheDurodola/go-email-service/src/config"
	pb "github.com/TheDurodola/go-email-service/src/config/proto_files"
	"google.golang.org/grpc"
)


type server struct {
	pb.UnimplementedEmailServiceServer
}



func main(){
	
	db, err := config.InitDB()
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }

    log.Println("Successfully connected to the database!")

}