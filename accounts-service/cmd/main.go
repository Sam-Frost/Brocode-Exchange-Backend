package main

import (
	"sync"

	eventLogger "github.com/Sam-Frost/accounts-service/internal/event-logger"
	"github.com/Sam-Frost/accounts-service/internal/grpc"
	"github.com/Sam-Frost/accounts-service/internal/service"
)

func main() {

	// Recreate database incase this isn' the first start of service
	service.RecreateDatabase()

	var wg sync.WaitGroup
	wg.Add(3)

	// Start gRPC server to communitcate with web-server
	go grpc.StartGrpcServer()

	// Start consuming events comming from the matching engine
	go service.ConsumeMatchingEngineEvents()

	// Start event logging
	go eventLogger.InitEventLogger()

	wg.Wait()

	// internal.RunTest()
}
