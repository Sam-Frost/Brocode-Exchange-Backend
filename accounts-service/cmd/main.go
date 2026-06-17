package main

import (
	"sync"

	eventlogger "github.com/Sam-Frost/accounts-service/internal/event-logger"
	"github.com/Sam-Frost/accounts-service/internal/grpc"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	// Start gRPC server to communitcate with web-server
	go grpc.StartGrpcServer()

	// Start consuming events comming from the matching engine
	// go service.ConsumeMatchingEngineEvents()

	// Start event logging
	go eventlogger.InitEventLogger()

	// Start state backup

	wg.Wait()
}
