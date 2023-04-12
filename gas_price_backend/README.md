# Gas Price

This project is a Go backend application with a PostgreSQL database that can be run using Docker containers.

## Prerequisites

Before running this project, you will need the following:

- Go installed on your system
- Docker installed on your system

## Running the Project

To run the project, use the following commands in your terminal:

1. Run the PostgreSQL database container:
    ```bash
    make db
    ```

2. Initialize the database with the required schema and data:


    ```bash
    make init-db
    ```
3. Run the Go backend application:

    ```bash
    make backend
    ```

## Makefile Targets
    `backend`: Runs the Go backend application.
    `db`: Starts a PostgreSQL database container using Docker.
    `init-db`: Initializes the database with the required schema and data.
