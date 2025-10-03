# AWS Task Management API - Serverless REST API

A production-ready serverless REST API built with **Go**, **AWS Lambda**, **API Gateway**, and **DynamoDB** for learning AWS serverless architecture.

## 🏗️ Architecture

```
┌─────────────┐      ┌──────────────┐      ┌─────────────┐      ┌──────────────┐
│   Client    │─────▶│ API Gateway  │─────▶│   Lambda    │─────▶│  DynamoDB    │
└─────────────┘      └──────────────┘      └─────────────┘      └──────────────┘
```

**Components:**
- **API Gateway**: RESTful API endpoint management
- **Lambda Functions**: 5 serverless functions handling CRUD operations
- **DynamoDB**: NoSQL database for persistent storage
- **AWS SAM**: Infrastructure as Code (IaC)

## ✨ Features

- ✅ **Create** tasks with title, description, and status
- ✅ **List** all tasks with count
- ✅ **Get** individual task by ID
- ✅ **Update** task properties
- ✅ **Delete** tasks
- ✅ **Serverless** - auto-scales, pay-per-use
- ✅ **Free tier eligible** - learn without costs!

## 📋 API Endpoints

| Method | Endpoint        | Description          | Request Body                              |
|--------|----------------|----------------------|-------------------------------------------|
| POST   | `/tasks`       | Create a new task    | `{"title": "...", "description": "..."}` |
| GET    | `/tasks`       | List all tasks       | -                                         |
| GET    | `/tasks/{id}`  | Get task by ID       | -                                         |
| PUT    | `/tasks/{id}`  | Update task          | `{"title": "...", "status": "..."}` (partial) |
| DELETE | `/tasks/{id}`  | Delete task          | -                                         |

## 📂 Project Structure

```
aws-task-api/
├── functions/                  # Lambda function handlers
│   ├── create-task/
│   │   └── main.go            # POST /tasks
│   ├── list-tasks/
│   │   └── main.go            # GET /tasks
│   ├── get-task/
│   │   └── main.go            # GET /tasks/{id}
│   ├── update-task/
│   │   └── main.go            # PUT /tasks/{id}
│   └── delete-task/
│       └── main.go            # DELETE /tasks/{id}
├── pkg/                       # Shared packages
│   ├── models/
│   │   └── task.go           # Task data model
│   └── dynamodb/
│       └── client.go         # DynamoDB client & operations
├── template.yaml             # AWS SAM CloudFormation template
├── Makefile                  # Build & deployment commands
├── go.mod                    # Go module dependencies
├── .gitignore               # Git ignore rules
└── README.md                # This file
```

## 🛠️ Prerequisites

Before you begin, ensure you have the following installed:

1. **AWS Account** - [Sign up for free](https://aws.amazon.com/free/)
2. **AWS CLI** - [Installation guide](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
3. **AWS SAM CLI** - [Installation guide](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html)
4. **Go 1.24+** - [Download Go](https://golang.org/dl/)
5. **Make** - Pre-installed on macOS/Linux, [Windows setup](https://gnuwin32.sourceforge.net/packages/make.htm)

## 🚀 Getting Started

### 1. Configure AWS Credentials

```bash
aws configure
```

You'll be prompted for:
- **AWS Access Key ID**: From AWS IAM console
- **AWS Secret Access Key**: From AWS IAM console
- **Default region**: e.g., `us-east-1`
- **Output format**: `json`

### 2. Clone & Setup Project

```bash
cd /Users/amitsingh/GolandProjects/aws-task-api

# Install Go dependencies
go mod download
```

### 3. Build Lambda Functions

```bash
make build
```

This compiles all 5 Lambda functions for Linux x86_64 architecture (required by AWS Lambda).

### 4. Deploy to AWS

**First deployment (interactive):**
```bash
make deploy
```

Answer the prompts:
- **Stack Name**: `aws-task-api` (or your choice)
- **AWS Region**: `us-east-1` (or your preferred region)
- **Confirm changes**: `Y`
- **Allow SAM CLI IAM role creation**: `Y`
- **Allow functions without authorization**: `Y` (for all 5 functions)
- **Save arguments to configuration**: `Y`

**Subsequent deployments:**
```bash
make deploy-fast
```

### 5. Get Your API URL

After successful deployment, look for:

```
Outputs
---------------------------------------------------------------
Key                 TaskAPIUrl
Description         API Gateway endpoint URL for Prod stage
Value               https://xxxxxxxxxx.execute-api.us-east-1.amazonaws.com/prod/
```

**Copy this URL** - you'll use it to test the API!

## 🧪 Testing the API

Replace `YOUR-API-URL` with your actual API Gateway URL from the deployment output.

### 1. Create a Task

```bash
curl -X POST https://YOUR-API-URL/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn AWS Lambda",
    "description": "Build a serverless API with Go",
    "status": "in-progress"
  }'
```

**Response:**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "Learn AWS Lambda",
  "description": "Build a serverless API with Go",
  "status": "in-progress",
  "created_at": "2025-10-03T12:00:00Z",
  "updated_at": "2025-10-03T12:00:00Z"
}
```

### 2. List All Tasks

```bash
curl https://YOUR-API-URL/tasks
```

**Response:**
```json
{
  "tasks": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "title": "Learn AWS Lambda",
      "description": "Build a serverless API with Go",
      "status": "in-progress",
      "created_at": "2025-10-03T12:00:00Z",
      "updated_at": "2025-10-03T12:00:00Z"
    }
  ],
  "count": 1
}
```

### 3. Get Task by ID

```bash
curl https://YOUR-API-URL/tasks/550e8400-e29b-41d4-a716-446655440000
```

### 4. Update a Task

```bash
curl -X PUT https://YOUR-API-URL/tasks/550e8400-e29b-41d4-a716-446655440000 \
  -H "Content-Type: application/json" \
  -d '{
    "status": "completed"
  }'
```

### 5. Delete a Task

```bash
curl -X DELETE https://YOUR-API-URL/tasks/550e8400-e29b-41d4-a716-446655440000
```

## 💰 Cost Analysis

### AWS Free Tier (Monthly)

| Service        | Free Tier                                  | Cost After Free Tier |
|----------------|-------------------------------------------|----------------------|
| **Lambda**     | 1M requests + 400,000 GB-seconds          | $0.20 per 1M requests |
| **API Gateway**| 1M calls (first 12 months)                | $3.50 per 1M calls   |
| **DynamoDB**   | 25 GB storage + 25 read/write units       | Pay-per-request      |

**For learning projects**: You'll stay **100% within the free tier!** ✅

### Example Usage Calculation

**Scenario**: 10,000 requests/month
- **Lambda**: 10K requests = FREE (under 1M limit)
- **API Gateway**: 10K calls = FREE (first year)
- **DynamoDB**: Minimal storage = FREE (under 25 GB)

**Total Cost**: **$0.00** 🎉

## 📊 Monitoring & Logs

### View Lambda Logs

```bash
# Tail logs for create-task function
sam logs -n CreateTaskFunction --stack-name aws-task-api --tail

# View logs for specific time range
sam logs -n CreateTaskFunction --stack-name aws-task-api --start-time '10min ago'
```

### Check DynamoDB Data

```bash
# Scan all items in Tasks table
aws dynamodb scan --table-name Tasks

# Get specific item
aws dynamodb get-item --table-name Tasks --key '{"id":{"S":"YOUR-TASK-ID"}}'
```

### View API Gateway Metrics

```bash
# Get API ID
aws apigateway get-rest-apis --query 'items[?name==`aws-task-api`].id' --output text
```

## 🧪 Local Development

Test your API locally before deploying:

```bash
make local
```

This starts a local API Gateway emulator at `http://localhost:3000`

**Test locally:**
```bash
curl -X POST http://localhost:3000/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Local Test", "status": "pending"}'
```

## 🔧 Development Commands

```bash
# Build all Lambda functions
make build

# Clean build artifacts
make clean

# Deploy with prompts (first time)
make deploy

# Fast deploy (uses saved config)
make deploy-fast

# Run tests
make test

# Start local API
make local
```

## 🧹 Cleanup (Delete Everything)

**⚠️ WARNING**: This permanently deletes all resources!

```bash
sam delete --stack-name aws-task-api
```

This removes:
- ✅ All 5 Lambda functions
- ✅ API Gateway
- ✅ DynamoDB table (and all data)
- ✅ IAM roles and policies
- ✅ CloudWatch log groups

## 🎓 What You'll Learn

### AWS Services
1. **AWS Lambda** - Serverless compute
2. **API Gateway** - RESTful API management
3. **DynamoDB** - NoSQL database
4. **CloudFormation** - Infrastructure as Code
5. **IAM** - Security & permissions
6. **CloudWatch** - Logging & monitoring

### Concepts
- Serverless architecture patterns
- Event-driven programming
- Infrastructure as Code (IaC)
- Pay-per-use pricing model
- Auto-scaling in the cloud
- RESTful API design

### Go Skills
- Lambda function handlers
- AWS SDK for Go
- JSON marshaling/unmarshaling
- Error handling in production code
- Environment variable configuration

## 🛡️ Security Best Practices

This project follows AWS security best practices:

✅ **Least Privilege IAM**: Lambda functions have minimal permissions
✅ **Resource Isolation**: Each Lambda can only access specific DynamoDB table
✅ **No Hardcoded Credentials**: Uses IAM roles for authentication
✅ **CORS Enabled**: Secure cross-origin requests
✅ **Input Validation**: Request validation in Lambda handlers

## 🚀 Next Steps & Enhancements

Ready to take it further? Try adding:

1. **Authentication**
   - Add AWS Cognito for user authentication
   - Implement JWT token validation

2. **Advanced Features**
   - Pagination for list endpoint
   - Filtering & sorting tasks
   - Search functionality
   - Task categories/tags

3. **Observability**
   - AWS X-Ray for distributed tracing
   - Custom CloudWatch metrics
   - Alarms for error rates

4. **CI/CD Pipeline**
   - GitHub Actions for automated deployment
   - Unit & integration tests
   - Automated rollbacks

5. **Performance**
   - DynamoDB indexes for faster queries
   - Lambda reserved concurrency
   - API Gateway caching

6. **Additional Services**
   - S3 for file attachments
   - SNS for notifications
   - SQS for async processing
   - EventBridge for scheduled tasks

## 📚 Additional Resources

### Documentation
- [AWS Lambda Go Documentation](https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html)
- [AWS SAM Documentation](https://docs.aws.amazon.com/serverless-application-model/)
- [DynamoDB Developer Guide](https://docs.aws.amazon.com/dynamodb/)
- [API Gateway Documentation](https://docs.aws.amazon.com/apigateway/)

### Tutorials
- [AWS Serverless Workshop](https://serverless-workshop.aws/)
- [Build a Serverless Web Application](https://aws.amazon.com/getting-started/hands-on/build-serverless-web-app-lambda-apigateway-s3-dynamodb-cognito/)

## 🐛 Troubleshooting

### Common Issues

**1. "sam: command not found"**
- Install AWS SAM CLI: [Guide](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html)

**2. "Error: No AWS credentials found"**
```bash
aws configure
# Enter your credentials
```

**3. "Build failed" errors**
```bash
# Make sure you're using Go 1.24+
go version

# Clear and rebuild
make clean
make build
```

**4. "Stack aws-task-api already exists"**
```bash
# Update existing stack
make deploy-fast
```

**5. Lambda function errors**
```bash
# Check logs
sam logs -n CreateTaskFunction --stack-name aws-task-api --tail
```

## 🤝 Contributing

This is a learning project! Feel free to:
- Fork and experiment
- Add new features
- Improve documentation
- Share your enhancements

## 📝 License

MIT License - Free to use for learning and commercial projects!

---

## 🎉 Congratulations!

You've successfully built and deployed a production-ready serverless API!

**You now understand:**
- ✅ How AWS Lambda works
- ✅ Building REST APIs with API Gateway
- ✅ Using DynamoDB for data persistence
- ✅ Infrastructure as Code with AWS SAM
- ✅ Cloud-native development with Go

**Keep building! 🚀**

---

**Questions or issues?** Check CloudWatch logs or review AWS documentation.

**Happy Learning! 💙**
