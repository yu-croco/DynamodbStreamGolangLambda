package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"yu-croco.com/DynamodbStreamGolangLambda/app/adapter/converter"
)

func Handler(request events.DynamoDBEvent) error {
	records, convertErr := converter.ToModel(request)
	if convertErr != nil {
		return convertErr
	}

	fmt.Print(records)

	return nil
}

func main() {
	lambda.Start(Handler)
}
