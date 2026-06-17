package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/Sam-Frost/accounts-service/protobufs"
	"google.golang.org/grpc"
)

func StartGrpcServer() {

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	fmt.Println("Starting listening for gRPC server")
	grpcServer := grpc.NewServer()
	protobufs.RegisterUserBalanceServiceServer(grpcServer, &userBalance{})
	grpcServer.Serve(lis)

}
