package service

import (
	"fmt"
	"time"
)

func ConsumeMatchingEngineEvents() {
	for {
		fmt.Println("Consuming and reacting to event...")

		time.Sleep(2 * time.Second)
	}
}
