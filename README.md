# golang-ddd-clear-architecture File Structure

## Overview
This repository implements a clear architecture based on Domain-Driven Design (DDD) principles using Go. It's structured to separate concerns and promote a clean and maintainable codebase.

## File Structure

- `README.md`: The main documentation file for the project.
- `docker-compose.yaml`: Docker Compose configuration file for setting up the project environment.
- `go.mod` and `go.sum`: Go module files for managing project dependencies.

### src Directory
The `src` directory contains the main application code, organized into various layers as per DDD and clean architecture.

#### Adapter Layer
- `adapter/controllers/task_controller.go`: Controller for task-related operations.
- `adapter/messages/success.go`: Success message definitions.
- `adapter/requests/task_request.go`: Request structures for task operations.
- `adapter/router.go`: Router setup for HTTP endpoints.

#### Config
- `config/env.go`: Configuration related to environment variables.

#### Domain Layer
- `domain/entities/task.go`: Entity definition for a task.
- `domain/entities/user.go`: Entity definition for a user.
- `domain/fields/task_status.go`: Definitions of task status fields.
- `domain/repositories/task_repository.go`: Interface for task repository.
- `domain/repositories/user_repository.go`: Interface for user repository.
- `domain/values/task_priority.go`: Definitions of task priority values.

#### Infrastructure Layer
- `infrastructure/mysql/clinet.go`: MySQL client setup.
- `infrastructure/task_repository.go`: Implementation of the task repository.

#### Main Application
- `main.go`: The entry point of the application.

#### Usecase Layer
- `usecase/dto/task_dto.go`: Data transfer objects for tasks.
- `usecase/params/task_params.go`: Parameter structures for task use cases.
- `usecase/task_usecase.go`: Business logic for task operations.

## Running the Project
To run the project, use Docker Compose:

