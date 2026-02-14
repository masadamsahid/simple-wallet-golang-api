# wallet-app

A Go-based wallet application built with the Fiber web framework. This project manages users, authentication, and financial transactions (deposits, withdrawals).

## Project Overview

*   **Type:** Go Web Application (REST API)
*   **Framework:** [Fiber v3](https://gofiber.io/)
*   **Database:** PostgreSQL (with [GORM](https://gorm.io/))
*   **Authentication:** JWT (JSON Web Tokens)
*   **Architecture:** Layered Architecture (Controller -> Service -> Repository)

## Key Features

*   **User Management:**
    *   Registration (`POST /api/auth/register`)
    *   Login (`POST /api/auth/login`)
    *   Retrieve Balance (`/api/users/balance`)
*   **Wallet Operations (MUST BE AUTHENTICATED):**
    *   Deposit Funds (`POST /api/transactions/deposit`)
    *   Withdraw Funds (`POST /api/transactions/withdraw`)
    *   Transaction History (`GET /api/transactions`)

## Directory Structure

*   **`main.go`**: Application entry point. Sets up the server, database connection, and routes.
*   **`controller/`**: Handles incoming HTTP requests, validates input, and calls services.
*   **`service/`**: Contains business logic. Orchestrates operations between repositories and other services.
*   **`repository/`**: Handles direct database interactions using GORM.
*   **`model/`**: Defines database schemas and models.
*   **`dto/`**: Data Transfer Objects used for request/response structures.
*   **`routes/`**: Defines API endpoints and maps them to controllers.
*   **`middlewares/`**: Contains middleware functions (e.g., authentication).
*   **`helpers/`**: Utility functions (e.g., JWT token generation, password hashing).
*   **`db/`**: Database configuration and migration scripts.
    *   `migrations/`: SQL migration files.

## Setup & Running

### Prerequisites

*   Go 1.25.4 or higher
*   PostgreSQL
*   [`migrate`](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) CLI tool (for database migrations)

### Configuration

1.  Copy the example environment file:
    ```bash
    cp .env.example .env
    ```
2.  Update `.env` with your database credentials and other settings:
    *   `DB_HOST`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `DB_PORT`
    *   `JWT_SECRET_KEY`

### Database Migrations

Use the [`migrate`](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) tool to set up the database schema:

```bash
# Run migrations
migrate -path="./db/migrations" -database="postgres://{USER}:{PASSWORD}@{HOST}:{PORT}/{DB_NAME}?sslmode=disable" up
```

See `db/migrations/README.md` for more details on creating and running migrations.

### Running the Application

```bash
go run main.go
```

The server will start on the port defined in `.env` (default: 8080).

## Development

### Coding Conventions

*   **Layered Architecture:** Follow the strict separation of concerns: Controller -> Service -> Repository.
*   **Error Handling:** Controllers should handle errors returned by services and return appropriate HTTP status codes.
*   **Validation:** Use `go-playground/validator` for request validation in controllers/DTOs.

### Testing

*   Currently, there are no tests in the project. (TODO: Add unit and integration tests).

## API Endpoints

### Authentication (`/api/auth`)
*   `POST /register`: Register a new user.
*   `POST /login`: Authenticate and receive a JWT token (access_token).

### Transactions (`/api/transactions`)
*   **Protected:** Requires `Authorization: Bearer <token>` header.
*   `POST /deposit`: Add funds to the user's wallet.
*   `POST /withdraw`: Deduct funds from the user's wallet.
*   `GET /`: Get the authenticated user's transaction history.

### Users (`/api/users`)
*   **Protected:** Requires `Authorization: Bearer <token>` header.
*   `GET /balance`: Get the authenticated user's current balance and details.
