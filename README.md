🚀 Go Microservices Architecture

A production-ready microservices template built with Go, the Chi router, and gRPC for internal service communication. It features a multi-database PostgreSQL backend, is fully containerized using Docker Compose, and uses an API Gateway as the single entry point for all external traffic.

Services Overview

API Gateway (Port 8080): Serves as the single entry point, handling reverse-proxy routing and enforcing JWT validation middleware.

Auth Service (Port 8081): Manages user registration, secure password hashing with bcrypt, and JWT token generation.

Product Service (Ports 8082 / 50051): Manages the product inventory catalog via REST and provides a high-performance gRPC server for internal stock and pricing verification.

Order Service (Port 8083): Handles order placement, acting as a gRPC client to securely query the Product Service before saving orders.

Payment Service (Port 8084): Processes transaction payments and tracks payment lifecycle statuses, secured behind authentication middleware.

Tech Stack

Language: Go (Golang)

Router: github.com/go-chi/chi/v5

Database & ORM: PostgreSQL with GORM (gorm.io/driver/postgres)

Inter-Service Communication: gRPC & Protocol Buffers (google.golang.org/grpc)

Security & Auth: JWT (golang-jwt/jwt/v5) and bcrypt

Containerization: Docker & Docker Compose

Project Structure

api-gateway/: Reverse proxy routing and authentication middleware.

auth-service/: User accounts, login handling, and token generation.

product-service/: Product catalog REST API and gRPC server.

order-service/: Order processing and gRPC client integration.

payment-service/: Payment transaction management.

proto/: Protocol buffer schemas and generated Go stubs.

docker-compose.yml: Multi-container orchestration configuration.

Dockerfile: Multi-stage generic build setup for services.

init.sql: PostgreSQL initialization script for creating multiple databases.

Quick Start Guide

Clone the repository:

git clone https://github.com/SurajSmidms/Go-Microservice-Project.git
cd Go-Microservice-Project


Build and start all containers:

docker-compose up --build -d


Verify running containers:

docker-compose ps


Stop containers when finished:

docker-compose down


API Endpoints Reference

All external requests must pass through the API Gateway running on port 8080.

1. Register a User

curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "mypassword123"}'


2. Log In (Get JWT Token)

curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "mypassword123"}'


(Copy the returned token string and use it below as YOUR_JWT_TOKEN)

3. Create a Product

curl -X POST http://localhost:8080/api/products \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name": "Mechanical Keyboard", "description": "Wireless RGB", "price": 120.50, "stock": 15}'


4. List Products

curl -X GET http://localhost:8080/api/products \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"


5. Create an Order (Triggers gRPC check)

curl -X POST http://localhost:8080/api/orders \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"user_id": 1, "product_id": 1, "quantity": 2}'


6. Process Payment

curl -X POST http://localhost:8080/api/payments \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"order_id": 1, "amount": 241.00}'


7. Check Payment Status

curl -X GET http://localhost:8080/api/payments/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
