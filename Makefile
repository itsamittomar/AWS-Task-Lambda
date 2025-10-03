.PHONY: build clean deploy test local

build:
	GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o functions/create-task/bootstrap functions/create-task/main.go
	GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o functions/list-tasks/bootstrap functions/list-tasks/main.go
	GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o functions/get-task/bootstrap functions/get-task/main.go
	GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o functions/update-task/bootstrap functions/update-task/main.go
	GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o functions/delete-task/bootstrap functions/delete-task/main.go

clean:
	rm -f functions/*/bootstrap

deploy: build
	sam deploy --guided

deploy-fast: build
	sam deploy

local:
	sam local start-api

test:
	go test ./...
