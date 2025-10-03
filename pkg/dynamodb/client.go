package dynamodb

import (
	"aws-task-api/pkg/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type Client struct {
	DynamoDB  dynamodbiface.DynamoDBAPI
	TableName string
}

// NewClient creates a new DynamoDB client
func NewClient(tableName string) (*Client, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Change to your preferred region
	})
	if err != nil {
		return nil, err
	}

	return &Client{
		DynamoDB:  dynamodb.New(sess),
		TableName: tableName,
	}, nil
}

// PutTask creates or updates a task in DynamoDB
func (c *Client) PutTask(task *models.Task) error {
	av, err := dynamodbattribute.MarshalMap(task)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(c.TableName),
	}

	_, err = c.DynamoDB.PutItem(input)
	return err
}

// GetTask retrieves a task by ID from DynamoDB
func (c *Client) GetTask(id string) (*models.Task, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(c.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}

	result, err := c.DynamoDB.GetItem(input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	task := &models.Task{}
	err = dynamodbattribute.UnmarshalMap(result.Item, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// ListTasks retrieves all tasks from DynamoDB
func (c *Client) ListTasks() ([]models.Task, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(c.TableName),
	}

	result, err := c.DynamoDB.Scan(input)
	if err != nil {
		return nil, err
	}

	tasks := []models.Task{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// DeleteTask deletes a task by ID from DynamoDB
func (c *Client) DeleteTask(id string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(c.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}

	_, err := c.DynamoDB.DeleteItem(input)
	return err
}
