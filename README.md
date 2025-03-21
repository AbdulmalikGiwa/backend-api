# Backend API

Simple API

## 🚀 Features

- User registration and authentication
- JWT token-based authentication
- PostgreSQL database integration
- Docker containerization
- Environment-based configuration
- RESTful API endpoints

## 🛠️ Prerequisites

Before running this project, make sure you have the following installed:
- [Go](https://golang.org/doc/install) (version 1.23 or higher)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## 🏃‍♂️ Running the Application

### Using Docker (Recommended)

1. **Clone the repository**
   ```bash
   git clone https://github.com/AbdulmalikGiwa/backend-api
   cd backend-api
   ```

2. **Create .env file**
   ```bash
   # Copy the example env file
   cp .env.example .env

   # Edit the .env file with your preferred settings
   # Example .env contents:
   SERVER_PORT=8080
   SERVER_HOST=0.0.0.0
   DB_HOST=postgres
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=auth_api
   DB_SSLMODE=disable
   JWT_SECRET=your-secret-key-here
   JWT_ISSUER=auth-api
   ```

3. **Build and run with Docker Compose**
   ```bash
   # Build and start containers
   docker-compose up --build

   # Or run in detached mode
   docker-compose up -d
   ```

4. **Stop the application**
   ```bash
   docker-compose down
   ```

## 📝 API Documentation

### Authentication Endpoints

#### Register User
```http
POST /api/v1/auth/register
Content-Type: application/json

{
    "username": "user@example.com",
    "password": "securepassword",
    "name": "John Doe"
}
```

#### Login
```http
POST /api/v1/auth/login
Content-Type: application/json

{
    "username": "user@example.com",
    "password": "securepassword"
}
```

## 🔧 Configuration

The application can be configured using environment variables in the `.env` file:

| Variable | Description | Default |
|----------|-------------|---------|
| SERVER_PORT | Port the server listens on | 8080 |
| SERVER_HOST | Host the server binds to | 0.0.0.0 |
| DB_HOST | PostgreSQL host | postgres |
| DB_PORT | PostgreSQL port | 5432 |
| DB_USER | Database username | postgres |
| DB_PASSWORD | Database password | postgres |
| DB_NAME | Database name | auth_api |
| JWT_SECRET | Secret key for JWT tokens | your-secret-key |
| JWT_ISSUER | JWT token issuer | auth-api |

## 🛠️ Development

### Project Structure