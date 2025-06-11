# Movie App Backend

This is the backend service for the Movie App application. It uses Go, PostgreSQL, and Docker for the development environment.

## Prerequisites

- Go 1.21 or later
- Docker and Docker Compose
- Git

## Setup Instructions

### 1. Install Go

1. Download and install Go from [golang.org](https://golang.org/dl/)
2. Verify installation:
   ```bash
   go version
   ```

### 2. Clone and Setup Project

1. Clone the repository:
   ```bash
   git clone https://github.com/dhruvvadoliya1/movie-app-backend.git
   cd movie-app-backend
   ```

2. Install Go dependencies:
   ```bash
   go mod download
   ```

### 3. Start Services with Docker Compose

1. Start the required services (PostgreSQL and Adminer):
   ```bash
   docker-compose up -d postgresdb adminer
   ```

2. Verify services are running:
   ```bash
   docker-compose ps
   ```

### 4. Database Migrations

1. Run database migrations:
   ```bash
   go run app.go migrate up
   ```

### 5. Start the API Server

1. Start the API server:
   ```bash
   go run app.go api
   ```

The server should now be running on port 3000.

## Additional Services

The project includes several optional services that can be started using profiles:

### Kratos (Identity Management)
```bash
docker-compose --profile kratos up -d
```

### Flipt (Feature Flags)
```bash
docker-compose --profile flipt up -d
```

## Accessing Services

- API Server: http://localhost:3000
- Adminer (Database Management): http://localhost:8080

## Environment Variables

Make sure to set up your environment variables in a `.env` file. Required variables include:

- DB_HOST=postgresdb
- DB_PORT=5432
- DB_USER=your_db_username
- DB_PASSWORD=your_secure_password
- DB_NAME=your_database_name
- PORT=3000

Example `.env` file:
```env
DB_HOST=postgresdb
DB_PORT=5432
DB_USER=app_user
DB_PASSWORD=your_secure_password_here
DB_NAME=movie_app_db
PORT=3000
```

## Troubleshooting

1. If you encounter database connection issues:
   - Ensure PostgreSQL container is running: `docker-compose ps`
   - Check logs: `docker-compose logs postgresdb`

2. If migrations fail:
   - Ensure database is running
   - Check database credentials in your environment variables

3. If the API server fails to start:
   - Verify all environment variables are set correctly
   - Check if port 3000 is available 
