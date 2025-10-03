package main

import (
	"aws-task-api/pkg/dynamodb"
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var dbClient *dynamodb.Client

func init() {
	tableName := os.Getenv("TABLE_NAME")
	if tableName == "" {
		tableName = "Tasks"
	}

	var err error
	dbClient, err = dynamodb.NewClient(tableName)
	if err != nil {
		panic(err)
	}
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// List all tasks from DynamoDB
	tasks, err := dbClient.ListTasks()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       `{"error": "Failed to list tasks"}`,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}

	// Return tasks
	responseBody, _ := json.Marshal(map[string]interface{}{
		"tasks": tasks,
		"count": len(tasks),
	})

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBody),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
