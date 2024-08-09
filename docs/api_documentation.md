```markdown
# Task Management API Documentation

This document provides comprehensive information about the Task Management API, which uses MongoDB for persistent data storage.

## Prerequisites

- Go 1.18 or later
- MongoDB 4.4 or later
- [Gin](https://github.com/gin-gonic/gin) framework for Go
- [MongoDB Go Driver](https://pkg.go.dev/go.mongodb.org/mongo-driver)

## Setup

1. Ensure you have MongoDB installed and running.
2. Set the `MONGODB_URI` environment variable in a `.env` file.

### Example `.env` File

```
MONGODB_URI=mongodb://localhost:27017/taskDB
```

3. Run the application:
  go to the root directory of the app
  ```
    go run main
  ```

The server will start on `http://localhost:8000`.

## API Endpoints

### Get All Tasks

- **URL**: `/tasks`
- **Method**: GET
- **Description**: Retrieves all tasks from the database.
- **Response**: 
  - Status Code: 200 OK
  - Body: JSON array of task objects

### Get a Specific Task

- **URL**: `/tasks/:id`
- **Method**: GET
- **Description**: Retrieves a specific task by its ID.
- **Parameters**: 
  - `id`: The ID of the task (string)
- **Response**: 
  - Status Code: 200 OK (if found), 404 Not Found (if not found)
  - Body: JSON object of the task (if found)

### Create a New Task

- **URL**: `/tasks`
- **Method**: POST
- **Description**: Creates a new task.
- **Request Body**: JSON object with task details
  ```json
  {
    "id": "string",
    "title": "string",
    "description": "string",
    "due_date": "2023-08-09T00:00:00Z",
    "status": "string"
  }
  ```
- **Response**: 
  - Status Code: 201 Created
  - Body: Success message

### Update a Task

- **URL**: `/tasks/:id`
- **Method**: PUT
- **Description**: Updates an existing task.
- **Parameters**: 
  - `id`: The ID of the task to update (string)
- **Request Body**: JSON object with updated task details
  ```json
  {
    "title": "string",
    "description": "string",
    "due_date": "2023-08-09T00:00:00Z",
    "status": "string"
  }
  ```
- **Response**: 
  - Status Code: 200 OK (if updated), 404 Not Found (if not found)
  - Body: Success message

### Delete a Task

- **URL**: `/tasks/:id`
- **Method**: DELETE
- **Description**: Deletes a task by its ID.
- **Parameters**: 
  - `id`: The ID of the task to delete (string)
- **Response**: 
  - Status Code: 200 OK (if deleted), 404 Not Found (if not found)
  - Body: Success message

## Data Model

The Task model has the following structure:

```go
type Task struct {
    ID          string    `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    DueDate     time.Time `json:"due_date"`
    Status      string    `json:"status"`
}
```

## Error Handling

The API returns appropriate HTTP status codes and error messages in the response body when errors occur. Common error responses include:

- 400 Bad Request: For invalid input data
- 404 Not Found: When a requested resource doesn't exist
- 500 Internal Server Error: For server-side errors

## Testing the API

Use [Postman](https://www.postman.com/) or [cURL](https://curl.se/) to test the API endpoints. Here's an example of how to test the `Get All Tasks` endpoint using cURL:

```
curl -X GET http://localhost:8000/tasks
```

The response should be a JSON array of tasks.

## MongoDB Inspection

You can use [MongoDB Compass](https://www.mongodb.com/products/compass) to inspect the data in your MongoDB instance.

## API Versioning

The current API version is `v1`. Future updates and changes may introduce new versions.

## Authorization

Currently, the API is open and does not require authentication. In future versions, authorization may be implemented.
```