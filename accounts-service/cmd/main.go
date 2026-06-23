package main

import (
	"sync"

	eventLogger "github.com/Sam-Frost/accounts-service/internal/event-logger"
	"github.com/Sam-Frost/accounts-service/internal/grpc"
	"github.com/Sam-Frost/accounts-service/internal/service"
)

func main() {

	// -------------------------CAN BE WRAPPED IN INIT FUNCTION---------------------------

	// Recreate database incase this isn' the first start of service
	service.RecreateDatabase()

	// Create event log and offset tracker file if not exist
	eventLogger.InitEventLogFile()
	eventLogger.InitOffsetTracker()

	// -------------------------CAN BE WRAPPED IN INIT FUNCTION---------------------------

	var wg sync.WaitGroup
	wg.Add(3)

	// Start gRPC server to communitcate with web-server
	go grpc.StartGrpcServer()

	// Start consuming events comming from the matching engine
	go service.ConsumeMatchingEngineEvents()

	// Start event logging
	go eventLogger.InitEventLogger()

	// Start kafka producer to send event log to kafka broker
	go eventLogger.SendEventToKafkaBroker()

	wg.Wait()

	// internal.RunTest()
}
