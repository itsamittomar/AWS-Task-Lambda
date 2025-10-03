package main

import (
	"aws-task-api/pkg/dynamodb"
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
	// Get task ID from path parameters
	taskID, ok := request.PathParameters["id"]
	if !ok || taskID == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `{"error": "Task ID is required"}`,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}

	// Check if task exists before deleting
	task, err := dbClient.GetTask(taskID)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       `{"error": "Failed to get task"}`,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}

	if task == nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       `{"error": "Task not found"}`,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}

	// Delete task from DynamoDB
	if err := dbClient.DeleteTask(taskID); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       `{"error": "Failed to delete task"}`,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}

	// Return success response
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       `{"message": "Task deleted successfully"}`,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
