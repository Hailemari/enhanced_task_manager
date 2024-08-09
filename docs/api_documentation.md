```markdown
# Task Management API Documentation

This document provides information about the Task Management API, which uses MongoDB for persistent data storage.

## Setup

1. Ensure you have MongoDB installed and running.
2. Set the `MONGODB_URI` environment variable in a `.env` file:

   ```
   MONGODB_URI=mongodb://localhost:27017/taskDB
   ```

3. Run the application:

   ```
   go run main.go
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

## Notes

- Ensure that the MongoDB connection string in the `.env` file is correct and points to your MongoDB instance.
- The API uses the Gin framework for routing and request handling.
- All date-time values are in ISO 8601 format.
```
