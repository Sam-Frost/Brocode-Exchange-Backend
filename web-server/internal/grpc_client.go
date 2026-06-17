package internal

import (
	"context"
	"fmt"
	"log"

	"github.com/Sam-Frost/accounts-service/protobufs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateGrpcClient() {

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient("localhost:50051", opts...)
	if err != nil {
		fmt.Printf("Creating gRPC client: %v", err)
		return
	}
	defer conn.Close()

	client := protobufs.NewUserBalanceServiceClient(conn)

	fmt.Println("Hitting gRPC server")

	response, err := client.CreateUserBalance(context.Background(), &protobufs.CreateUserBalanceRequest{
		UserId: 1,
	})

	if err != nil {
		log.Fatalf("Error creating user balance : %v", err)
	}

	fmt.Println(response)

	response2, err := client.GetUserBalance(context.Background(), &protobufs.GetUserBalanceRequest{
		UserId: 1,
	})

	if err != nil {
		log.Fatalf("Error getting user balance  : %v", err)
	}
	fmt.Println("This is empty")
	fmt.Println(response2)
	fmt.Println(response2.AvailableBalance)
	fmt.Println(response2.LockedBalance)

	// response4, err := client.IncreaseUserBalance(context.Background(), &protobufs.IncreaseUserBalanceRequest{
	// 	UserId: 1,
	// 	Amount: 100,
	// })

	// if err != nil {
	// 	log.Fatalf("Error increasing user balance  : %v", err)
	// }

	// fmt.Println(response4)

	// response3, err := client.LockUserBalance(context.Background(), &protobufs.LockUserBalanceRequest{
	// 	UserId: 1,
	// 	Amount: 10,
	// })

	// if err != nil {
	// 	log.Fatalf("Error locking user balance  : %v", err)
	// }

	// fmt.Println(response3)

}
