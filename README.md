# Simple CRUD Go Backend System

A Simple CRUD Management System built with **Golang**, **PostgreSQL**, **Gin**, and **GORM**.

---

## 🚀 Features

- User, Category & Product Management
- JWT Authentication
- Swagger API Documentation
- Input Validation & Security Middleware

---

## 🏗️ Project Structure

```
/cmd/server/main.go     → Main entry
/internal/
    /config             → Config & DB setup
    /domain             → Data models
    /repository         → DB access layer (GORM)
    /service            → Business logic
    /handler            → API Handlers
    /middleware         → JWT, Security headers, etc.
    /dto                → Request/Response structs
/docs/                  → Swagger generated docs and files
/migrations/            → SQL Migration structure
```

---

## 📚 API Documentation (Swagger)

Run the app and access:

```
http://localhost:8080/swagger/index.html
```

Or access the postman documentation: https://documenter.getpostman.com/view/1475503/2sAYkGKeKE

## 🛠️ Setup & Run

```bash
# Install dependencies
go mod tidy

# Generate Swagger docs
swag init --generalInfo cmd/server/main.go --output docs

# Run app
go run cmd/server/main.go
```

---

## 🐳 Quick Start (Docker Compose)

```bash
# Run the entire stack (App + DB)
docker-compose up --build
```
API: http://localhost:8080
