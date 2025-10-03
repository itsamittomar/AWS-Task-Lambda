package models

import "time"

// Task represents a task in our system
type Task struct {
	ID          string    `json:"id" dynamodbav:"id"`
	Title       string    `json:"title" dynamodbav:"title"`
	Description string    `json:"description" dynamodbav:"description"`
	Status      string    `json:"status" dynamodbav:"status"` // pending, in-progress, completed
	CreatedAt   time.Time `json:"created_at" dynamodbav:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" dynamodbav:"updated_at"`
}

// CreateTaskRequest represents the request body for creating a task
type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// UpdateTaskRequest represents the request body for updating a task
type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
