Go Microservices Architecture
A production-ready microservices template built with Go, the Chi router, and gRPC for internal service-to-service communication. Designed with a multi-database PostgreSQL backend, containerized via Docker Compose, and fronted by an API Gateway.
🏗️ Architecture & Communication Flow
+-----------------------+
                     |    External Client    |
                     +-----------+-----------+
                                 | REST / JSON (Port 8080)
                                 v
                     +-----------------------+
                     |      API Gateway      |
                     +---+---+---+---+-------+
         +---------------+   |   |   +----------------+
         | (HTTP)            |   |            | (HTTP)
         v                   |   |            v
   +-----------+             |   |      +------------+
   |   Auth    |             |   |      |  Payment   |
   |  Service  |             |   |      |  Service   |
   | (Port 8081|             |   |      | (Port 8084)|
   +-----------+             |   |      +------------+
                             v   v
                +-------------------------+      gRPC (Port 50051)      +-------------------+
                |      Order Service      | --------------------------> |  Product Service  |
                |      (Port 8083)        |                             |   (Port 8082)     |
                +-------------------------+                             +-------------------+
Service Breakdown
API Gateway (8080): Single entry point handling reverse-proxy routing and enforcing JWT verification middleware.

Auth Service (8081): Manages user registration, password hashing (bcrypt), and token generation (JWT).

Product Service (8082 / gRPC 50051): Manages the product inventory catalog via REST and provides a fast internal gRPC server for stock and pricing checks.

Order Service (8083): Handles order creation, acting as a gRPC client to validate and fetch product details before persisting orders.

Payment Service (8084): Processes payments and tracks transaction lifecycle states, secured behind middleware.

🛠️ Technology Stack
Language: Go (Golang)

HTTP Router: [github.com/go-chi/chi/v5](https://github.com/go-chi/chi/v5)

ORM & Database: GORM (gorm.io/gorm, gorm.io/driver/postgres) with PostgreSQL

RPC Framework: gRPC & Protocol Buffers (google.golang.org/grpc)

Authentication: JWT ([github.com/golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt/v5)) & bcrypt

Containerization: Docker & Docker Compose

📁 Repository Structure
Code snippet
GOLANG_MICROSERVICE_PROJECT/
├── api-gateway/         # Reverse proxy & Auth middleware
├── auth-service/        # User accounts, authentication & JWT
├── product-service/     # Product catalog REST API & gRPC server
├── order-service/       # Order management & gRPC client
├── payment-service/     # Payment transaction processing
├── proto/               # Protobuf schemas & generated stubs
├── docker-compose.yml   # Multi-container orchestration config
├── Dockerfile           # Multi-stage generic build setup
├── init.sql             # PostgreSQL initial multi-database setup
└── go.mod / go.sum      # Root Go module tracking
🚀 Quick Start Guide (Docker)
Clone the repository:

Bash
git clone https://github.com/SurajSmidms/Go-Microservice-Project.git
cd Go-Microservice-Project
Build and spin up all containers:

Bash
docker-compose up --build -d
Verify running containers:

Bash
docker-compose ps
Tear down containers:

Bash
docker-compose down
🧪 API Endpoints Reference (cURL)
All requests route through the API Gateway on port 8080.

1. Register User
Bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "mypassword123"}'
2. Login (Get JWT Token)
Bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "mypassword123"}'
(Copy the resulting token string and use it below as YOUR_JWT_TOKEN)

3. Create Product
Bash
curl -X POST http://localhost:8080/api/products \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name": "Mechanical Keyboard", "description": "Wireless RGB", "price": 120.50, "stock": 15}'
4. List Products
Bash
curl -X GET http://localhost:8080/api/products \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
5. Create Order (Triggers gRPC check)
Bash
curl -X POST http://localhost:8080/api/orders \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"user_id": 1, "product_id": 1, "quantity": 2}'
6. Process Payment
Bash
curl -X POST http://localhost:8080/api/payments \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"order_id": 1, "amount": 241.00}'
7. Get Payment Status
Bash
curl -X GET http://localhost:8080/api/payments/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
