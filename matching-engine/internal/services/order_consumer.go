package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	event "github.com/Sam-Frost/contract/events"
	"github.com/Sam-Frost/matching-engine/internal/util"
	"github.com/twmb/franz-go/pkg/kgo"
)

func StartOrderConsumer(ringBuffer *util.RingBuffer[event.Event[any]], kafkaClient *kgo.Client) {
	fmt.Println("Starting order consumer...")

	for {
		fmt.Println("Polling...")
		polledRecords := kafkaClient.PollRecords(context.Background(), 10000)

		if polledRecords.IsClientClosed() {
			log.Fatalf("Kafka client got closed")

		}
		if polledRecords.Err() != nil {
			log.Fatalf("Error occuered while consuing orders: %v", polledRecords.Err())
		}

		polledRecords.EachRecord(func(record *kgo.Record) {

			fmt.Println("Get a record : %s", record)
			var order event.Event[any]

			err := json.Unmarshal(record.Value, &order)
			if err != nil {
				// log.Fatalf("Error parsing kafka event : %v", err)
				fmt.Println("Error parsing kafka event : %v", err)
			}

			// ringBuffer.PushWait(order)
		})
	}
}
