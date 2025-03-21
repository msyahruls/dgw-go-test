# KreditPlus Go Backend System

A Kredit Management System built with **Golang**, **PostgreSQL**, **Gin**, and **GORM**, supporting secure, concurrent processing of credit transactions and robust security practices.

---

## 🚀 Features

- User Management
- Credit Limit Management per tenor
- Transaction Recording with Limit Deduction
- Pay Installment per month
- JWT Authentication
- Swagger API Documentation
- Secure, concurrency-safe transaction handling
- Input Validation, Rate Limiting & Security Middleware

---

## 📝 ERD and Architecture
![alt text](https://raw.githubusercontent.com/msyahruls/kreditplus-go-test/refs/heads/main/assets/kreditplus-test.drawio.png)

## 🐘 Database SQL
https://raw.githubusercontent.com/msyahruls/kreditplus-go-test/refs/heads/main/assets/Localhost-2025-03-21-21_47_35.sql

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

Or access the postman documentation: 
https://documenter.getpostman.com/view/1475503/2sAYkGKeKE

---

## 🔐 Security Implementations (OWASP Top 10)

| OWASP Category                         | Mitigation Implemented                                               |
|----------------------------------------|----------------------------------------------------------------------|
| **Injection (SQL Injection)**          | Using **GORM ORM**, which automatically parameterizes queries.       |
| **Authentication & Session Management**| JWT-based authentication with expiry, `Bearer` token required.       |
| **Broken Access Control**              | All sensitive endpoints are protected with JWT middleware.           |
| **Security Misconfiguration**          | Sensitive configs in `.env`, debug routes disabled.                  |
| **Cryptographic Failures**             | JWT uses **HS256**, with secret stored securely.                     |
| **Input Validation (Insecure Design)** | Strict input validation via Gin binding & validator tags.            |
| **Security Headers (Misconfiguration)**| Middleware added for headers like `X-Frame-Options`, `CSP`, etc.     |
| **Rate Limiting (Auth Failures)**      | Middleware limits request rate to prevent brute force attacks.       |

---

## ⚙️ Concurrency Handling

The system ensures safe concurrent transaction processing to prevent double spending:

- **Database Transactions:**  
  Each transaction (limit check, deduction, transaction record) is wrapped inside a DB transaction to ensure atomicity and rollback safety.

- **Row-Level Locking:**  
  Row-level lock implemented with `SELECT FOR UPDATE` in GORM:

  ```go
  tx.Clauses(clause.Locking{Strength: "UPDATE"}).
      Where("user_id = ? AND tenor_months = ?", userID, tenor).
      First(&limit)
  ```

  Ensures that only one transaction can modify a user’s limit at a time.

<!-- - **Concurrent Batch Support:**  
  The service supports batch processing of transactions using Goroutines & WaitGroup, with DB transaction & locking mechanisms ensuring safe concurrent execution. -->

---

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

## 🧪 Tests

```bash
# Unit tests
go test ./... -v -cover

# Run concurrent transaction tests
go test ./internal/service -run TestConcurrentTransaction
```
