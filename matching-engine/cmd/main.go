package main

import (
	"log"
	"sync"

	event "github.com/Sam-Frost/contract/events"
	"github.com/Sam-Frost/contract/order"
	"github.com/Sam-Frost/matching-engine/internal/orderbook"
	"github.com/Sam-Frost/matching-engine/internal/services"
	"github.com/Sam-Frost/matching-engine/internal/util"
)

func main() {

	util.LoadEnv()

	envVariables, err := util.GetEnvVariables()
	if err != nil {
		log.Fatalf("Env variables not loaded... : %v", err)
	}

	// TODO : Start Threads for each market
	initEngine(envVariables, "BTC", util.SPOT)
}

func initEngine(envVariables util.EnvVariables, marketName string, marketType string) {

	// Kafka Consumer -> Engine
	ordersRingBuf := util.NewRingBuffer[event.Event[any]](65536)

	// Engine -> Kafka Producer
	fillsRingBuf := util.NewRingBuffer[[]services.Fill](65536)

	kafkaClient := services.CreateKafkaClient(envVariables.KafkaURL, marketName, marketType)

	var wg sync.WaitGroup
	wg.Add(3)
	go services.StartOrderConsumer(ordersRingBuf, kafkaClient)
	go StartMatchingEngine(ordersRingBuf, fillsRingBuf)
	go services.StartFillProducer(fillsRingBuf, kafkaClient)

	wg.Wait()
}

func StartMatchingEngine(ordersRingBuf *util.RingBuffer[event.Event[any]], fillsRingBuf *util.RingBuffer[[]services.Fill]) {
	orderbook := orderbook.CreateOrderbook()

	for {
		orderData := ordersRingBuf.PopWait()
		// Inefering types based on data

		switch orderData.Type {
		case event.CreateSpotOrder:
			createOrderData := orderData.Data.(order.CreateOrder) // Casting
			services.CreateOrder(createOrderData, &orderbook, fillsRingBuf)
		case event.CancelOrder:
			services.CancelOrder()
		default:
			log.Fatalf("Engine crashed, unknown event")
		}
	}
}
