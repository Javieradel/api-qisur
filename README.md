# Go API Qisur

This is a sample API built with Go and Fiber that manages products and categories. It uses a PostgreSQL database and provides a Swagger interface for API documentation.

## Features

-   CRUD operations for products and categories.
-   Product change history tracking.
-   Event-driven architecture for decoupling components.
-   Swagger documentation for the API.
-   Support for running with Docker or locally.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

-   [Go](https://golang.org/doc/install) (version 1.25 or higher)
-   [Docker](https://docs.docker.com/get-docker/)
-   [Docker Compose](https://docs.docker.com/compose/install/)
-   [Air](https://github.com/cosmtrek/air) for live reloading
-   [Swag](https://github.com/swaggo/swag) for Swagger documentation generation

You can install Air and Swag with the following commands:

```sh
go install github.com/cosmtrek/air@latest
go install github.com/swaggo/swag/cmd/swag@latest
```

### Installation

1.  **Clone the repository**

    ```sh
    git clone <repository-url>
    cd <repository-name>
    ```

2.  **Install dependencies**

    ```sh
    go mod tidy
    ```

## Configuration

The application requires a PostgreSQL database. Configuration is managed through environment variables.

1.  **Create a `.env` file**

    Create a `.env` file in the root of the project by copying the `.env.example` file:

    ```sh
    cp .env.example .env
    ```

2.  **Set environment variables**

    The `.env` file should contain the following variables:

    ```env
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=qisur_dev
    DB_PASSWORD=qisur_dev
    DB_NAME=qisur-products
    DB_SSLMODE=disable
    ```

## Usage

### Running the application

You can run the application using Docker or locally.

#### With Docker

The easiest way to get started is by using Docker Compose. This will start the PostgreSQL database.

```sh
docker-compose up -d
```

Then, you can run the application locally, and it will connect to the database running in Docker.

#### Locally with Air (Live Reloading)

For development, you can use `air` for live reloading. This will automatically rebuild and restart the application when you change a file.

1.  Make sure the database is running (e.g., with `docker-compose up -d`).

2.  Run the application with `air`:

    ```sh
    air
    ```

This will start the API server on `http://localhost:3000`.

#### Locally with `go run`

If you don't want to use `air`, you can run the application with `go run`:

1.  Make sure the database is running.

2.  Generate swagger documentation:

    ```sh
    swag init -g ./src/cmd/api/app.go -o ./docs
    ```

3.  Run the application:

    ```sh
    go run ./src/cmd/api/
    ```


## API Documentation

The API documentation is generated with Swagger and is available at:

[http://localhost:3000/api/docs/index.html](http://localhost:3000/api/docs/index.html)

## Future Enhancements (TODOs)

*   Implement JWT authentication.
*   Create basic roles (admin, client, etc.).
*   Protect routes based on roles.
*   Implement a WebSocket system.
*   Implement a WebSocket system to emit real-time events for:
    *   Creation of new products/categories.
    *   Updates to products/categories.
    *   Deletion of products/categories.
