package util

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	KafkaURL string
}

var envVariables EnvVariables

var isEnvLoaded bool = false

func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Trying to load env file : %v", err)
	}

	envVariables = EnvVariables{
		KafkaURL: os.Getenv("KAFKA_URL"),
	}

	isEnvLoaded = true
	fmt.Println("Env variables loaded!")
	fmt.Println(envVariables)
}

func GetEnvVariables() (EnvVariables, error) {
	if isEnvLoaded {
		return envVariables, nil
	}

	return EnvVariables{}, errors.New("Env variables not loaded yet, cannot access")
}
