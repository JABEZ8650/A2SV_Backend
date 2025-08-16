# Task Management API Documentation

## Endpoints

### GET /tasks
- **Description**: Retrieves all tasks.
- **Response**: A list of all tasks.

### GET /tasks/:id
- **Description**: Retrieves a task by ID.
- **Response**: The task with the specified ID.

### POST /tasks
- **Description**: Creates a new task.
- **Request Body**: Task object.
- **Response**: Confirmation message.

### PATCH /tasks/:id
- **Description**: Updates a task by ID.
- **Request Body**: Fields to update.
- **Response**: Confirmation message.

### DELETE /tasks/:id
- **Description**: Deletes a task by ID.
- **Response**: Confirmation message.

## Example:
- **GET request: http://localhost:8080/tasks (to get all tasks)**
- **GET request: http://localhost:8080/tasks/1 (to get a task with ID 1)**
- **POST request: http://localhost:8080/tasks (to create a new task)**
- **PATCH request: http://localhost:8080/tasks/1 (to update an existing task)**
- **DELETE request: http://localhost:8080/tasks/1 (to delete a task)**