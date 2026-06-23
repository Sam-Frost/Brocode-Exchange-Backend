package util

import (
	"log"
	"os"
)

type Env struct {
	KafkaUrl string
}

var EnvVariable Env = Env{}

func LoadEnv() {
	value, exists := os.LookupEnv("KAFKA_URL")
	if !exists {
		log.Fatal("KAFKLA_URL doesn't exist in env")
	}

	EnvVariable.KafkaUrl = value
}
