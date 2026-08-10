package service

import (
	"context"
	"fmt"
	"time"
)

func ConsumeMatchingEngineEvents(ctx context.Context) error {
	for {
		fmt.Println("Consuming and reacting to event...")

		time.Sleep(2 * time.Second)
	}

	return nil
}
