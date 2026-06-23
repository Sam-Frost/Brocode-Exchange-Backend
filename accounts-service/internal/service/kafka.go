package service

import (
	"log"

	"github.com/Sam-Frost/accounts-service/internal/util"
	"github.com/twmb/franz-go/pkg/kgo"
)

var kafkaProducer *kgo.Client

func CreateProducer() *kgo.Client {

	p, err := kgo.NewClient(
		kgo.SeedBrokers(util.EnvVariable.KafkaUrl),
		kgo.DefaultProduceTopic("accounts-service-events"),
	)
	if err != nil {
		log.Fatalf("Error creating kafka producer %v", err)
	}

	return p
}

func GetConsumer() *kgo.Client {
	return kafkaProducer
}
