package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

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
	//

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	errChan := make(chan error, 4)
	// Start gRPC server to communitcate with web-server
	go func() { errChan <- grpc.StartGrpcServer(ctx) }()

	// Start consuming events comming from the matching engine
	go func() { errChan <- service.ConsumeMatchingEngineEvents(ctx) }()

	// Start event logging
	go func() { errChan <- eventLogger.InitEventLogger(ctx) }()

	// Start kafka producer to send event log to kafka broker
	go func() { errChan <- eventLogger.SendEventToKafkaBroker(ctx) }()

	select {
	case err := <-errChan:
		fmt.Printf("Copmonent Failed. Shuting system down... \n Error :  %s", err)
		cancel()
	case <-ctx.Done():
		fmt.Printf("Shutting down....")
	}
	// internal.RunTest()
}
