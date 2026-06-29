package services

import (
	"fmt"
	"log"
	"strings"

	"github.com/twmb/franz-go/pkg/kgo"
)

func CreateKafkaClient(kafkaUrl, marketName, marketType string) *kgo.Client {

	consumerGroup := fmt.Sprintf("%s-%s-matching-engine", strings.ToLower(marketName), strings.ToLower(marketType))
	consumerTopic := fmt.Sprintf("%s-%s-orders", strings.ToLower(marketName), strings.ToLower(marketType))

	client, err := kgo.NewClient(kgo.SeedBrokers(kafkaUrl), kgo.ConsumerGroup(consumerGroup), kgo.ConsumeTopics(consumerTopic))
	if err != nil {
		log.Fatalf("Trying to connect to kafka... : %v", err)
	}

	fmt.Println("Kafka client succesfully created...")
	fmt.Printf(`Consumer Group : %s
Topic : %s
`, consumerGroup, consumerTopic)

	return client
}
