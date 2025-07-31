# Project Structure Documentation

This document explains the project structure based on Clean Architecture principles with three main layers: **Delivery**, **Usecase**, and **Repository**.

## Overview

The project follows Clean Architecture patterns to ensure separation of concerns, testability, and maintainability. The architecture consists of concentric layers where inner layers are independent of outer layers.

```
lets-go/
├── cmd/                    # CLI commands and application entry points
├── internal/               # Private application code
│   ├── delivery/          # Presentation Layer (HTTP handlers, controllers)
│   ├── usecase/           # Business Logic Layer (use cases, services)
│   ├── repository/        # Data Access Layer (database, external APIs)
│   └── domain/            # Core Domain (entities, interfaces)
│       ├── entity/        # Domain entities and DTOs
│       └── mock/          # Mock implementations for testing
├── samples/               # Example implementations and demos
└── main.go               # Application entry point
```

## Architecture Layers

### 1. Domain Layer (`internal/domain/`)

The **innermost layer** containing the core business logic and rules. This layer has no dependencies on external frameworks or libraries.

#### Components:
- **Entities** (`entity/`): Core business objects and data structures
- **Interfaces**: Define contracts between layers
  - `delivery_interface.go`: Contracts for delivery layer
  - `usecase_interface.go`: Contracts for business logic
  - `repository_interface.go`: Contracts for data access
- **Mocks** (`mock/`): Mock implementations for testing

#### Key Principles:
- Contains pure business logic
- No external dependencies
- Defines interfaces that outer layers implement
- Independent of frameworks, databases, or UI

### 2. Usecase Layer (`internal/usecase/`)

The **business logic layer** that orchestrates the flow of data and implements use cases by coordinating between repositories and applying business rules.

#### Responsibilities:
- Implement business use cases
- Coordinate between different repositories
- Apply business rules and validation
- Transform data between layers
- Handle business logic workflows

#### Dependencies:
- **Depends on**: Domain interfaces
- **Independent of**: Delivery layer, external frameworks

### 3. Repository Layer (`internal/repository/`)

The **data access layer** responsible for data persistence and retrieval from external sources like databases, APIs, or file systems.

#### Responsibilities:
- Implement data access logic
- Handle database operations (CRUD)
- Manage external API calls
- Data transformation from external sources to domain entities
- Caching strategies

#### Dependencies:
- **Depends on**: Domain entities and repository interfaces
- **Implements**: Repository interfaces defined in domain layer

### 4. Delivery Layer (`internal/delivery/`)

The **presentation layer** that handles external communication such as HTTP requests, CLI commands, or message queue consumers.

#### Responsibilities:
- Handle HTTP requests/responses
- Input validation and sanitization
- Request/response transformation
- Route handling
- Authentication and authorization
- Error handling and formatting

#### Dependencies:
- **Depends on**: Usecase interfaces
- **Independent of**: Repository implementations

## Data Flow

The data flows in one direction following the dependency rule:

```
External Request → Delivery → Usecase → Repository → Database/External API
                     ↓         ↓         ↓
                 Response ← Business ← Data
```

1. **Request Flow**: External requests enter through the delivery layer
2. **Business Processing**: Delivery layer calls usecase layer for business logic
3. **Data Access**: Usecase layer calls repository layer for data operations
4. **Response Flow**: Data flows back through the layers to form the response

## Dependency Injection

The architecture uses dependency injection to maintain loose coupling:

- **Interfaces are defined in the domain layer**
- **Implementations are in their respective layers**
- **Dependencies are injected at application startup**

## Example Structure per Feature

Each feature (like `greeting`, `config`) follows the same structure across all layers:

```
Feature: greeting
├── internal/domain/entity/greeting/     # Domain entities
├── internal/repository/greeting/        # Data access implementation
├── internal/usecase/greeting/          # Business logic implementation
└── internal/delivery/greeting/         # HTTP handlers/controllers
```

## Benefits of This Architecture

1. **Testability**: Each layer can be tested independently using mocks
2. **Maintainability**: Clear separation of concerns makes code easier to maintain
3. **Flexibility**: Easy to swap implementations (e.g., change database)
4. **Scalability**: New features can be added following the same pattern
5. **Independence**: Business logic is independent of external frameworks

## Development Guidelines

1. **Domain entities should be pure Go structs** with no external dependencies
2. **Interfaces should be defined in the domain layer** and implemented by outer layers
3. **Use dependency injection** to provide implementations to dependent layers
4. **Keep business logic in the usecase layer**, not in delivery or repository
5. **Repository layer should only handle data access**, no business logic
6. **Delivery layer should only handle presentation concerns**, delegate business logic to usecase layer

## Testing Strategy

- **Unit Tests**: Test each layer independently using mocks
- **Integration Tests**: Test the interaction between layers
- **Mock Generation**: Use the provided mock interfaces for testing
- **Test Coverage**: Aim for high coverage especially in the usecase layer where business logic resides

This architecture ensures that your application is maintainable, testable, and follows clean architecture principles while remaining flexible for future changes and requirements.