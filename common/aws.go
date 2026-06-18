package common

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AwsClient struct {
	Client *s3.Client
}

func (a *AwsClient) UploadSnapshot(filePath string) {
	fmt.Println("Uploading snapshot to S3 bucket...")
}

func (a *AwsClient) GetLatestSnaspshot() {
	fmt.Println("Retreiving lastest snapshot from S3 bucket...")
}

func CreateClient() AwsClient {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	return AwsClient{
		Client: client,
	}
}
