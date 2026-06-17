package grpc

import (
	"context"
	"fmt"

	"github.com/Sam-Frost/accounts-service/internal/db"
	"github.com/Sam-Frost/accounts-service/protobufs"
)

type userBalance struct {
	protobufs.UnimplementedUserBalanceServiceServer
}

func (u *userBalance) CreateUserBalance(ctx context.Context, request *protobufs.CreateUserBalanceRequest) (*protobufs.CreateUserBalanceResponse, error) {

	db.CreateNewUser(request.UserId)

	return &protobufs.CreateUserBalanceResponse{
		Success: true,
	}, nil
}

func (u *userBalance) GetUserBalance(ctx context.Context, request *protobufs.GetUserBalanceRequest) (*protobufs.GetUserBalanceResponse, error) {

	userBalanceData, err := db.GetUserBalanceData(request.UserId)
	if err != nil {
		fmt.Println(err)
		// TODO : Send error in grpc header
	}

	availableBalance, lockedBalance := userBalanceData.GetUserBalance()

	return &protobufs.GetUserBalanceResponse{
		AvailableBalance: availableBalance,
		LockedBalance:    lockedBalance,
	}, nil
}

func (u *userBalance) IncreaseUserBalance(ctx context.Context, request *protobufs.IncreaseUserBalanceRequest) (*protobufs.IncreaseUserBalanceResponse, error) {

	userBalanceData, err := db.GetUserBalanceData(request.UserId)
	if err != nil {
		// TODO : Send error in grpc header
	}

	udpatedBalance := userBalanceData.IncreaseAvailableBalance(request.Amount, request.UserId)

	return &protobufs.IncreaseUserBalanceResponse{
		UpdatedBalance: udpatedBalance,
	}, nil
}

func (u *userBalance) LockUserBalance(ctx context.Context, request *protobufs.LockUserBalanceRequest) (*protobufs.LockUserBalanceResponse, error) {

	userBalanceData, err := db.GetUserBalanceData(request.UserId)
	if err != nil {
		// TODO : Send error in grpc header
	}

	availableBalance, lockedBalance, err := userBalanceData.LockBalance(request.Amount, request.UserId)
	if err != nil {
		// TODO : Send error in grpc header
	}

	return &protobufs.LockUserBalanceResponse{
		AvailableBalance: availableBalance,
		LockedBalance:    lockedBalance,
	}, nil
}
