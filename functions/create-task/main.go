package main

import (
	"aws-task-api/pkg/dynamodb"
	"aws-task-api/pkg/models"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
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
	// Parse request body
	var req models.CreateTaskRequest
	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `{"error": "Invalid request body"}`,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}

	// Validate required fields
	if req.Title == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `{"error": "Title is required"}`,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}

	// Set default status if not provided
	if req.Status == "" {
		req.Status = "pending"
	}

	// Create task
	task := &models.Task{
		ID:          uuid.New().String(),
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Save to DynamoDB
	if err := dbClient.PutTask(task); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       `{"error": "Failed to create task"}`,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	}

	// Return created task
	responseBody, _ := json.Marshal(task)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       string(responseBody),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
