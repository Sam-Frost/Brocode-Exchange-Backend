package grpc

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/Sam-Frost/accounts-service/protobufs"
	"google.golang.org/grpc"
)

func StartGrpcServer(ctx context.Context) error {

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		return errors.New("Failed to listen ")
	}

	fmt.Println("Starting listening for gRPC server")
	grpcServer := grpc.NewServer()
	protobufs.RegisterUserBalanceServiceServer(grpcServer, &userBalance{})
	grpcServer.Serve(lis)

	return nil
}
