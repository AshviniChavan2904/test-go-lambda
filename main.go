package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

// Handler is your Lambda function handler
func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Received request: %+v", event)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello from Lambda!",
	}, nil
}

func main() {
	lambda.Start(Handler)
}