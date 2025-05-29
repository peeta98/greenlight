# 🎬 Greenlight API

This project implements **Greenlight**, a JSON-based REST API for managing and retrieving movie data. It's built from the ground up in Go, following the full workflow and best practices outlined in [_Let's Go Further_](https://lets-go-further.alexedwards.net/) by Alex Edwards.

The API is designed for robustness, maintainability, and production-readiness — with a strong focus on layered architecture, performance, security, and clean design.

---

## 🌐 API Endpoints

### 📦 Core Movie Operations

| Method | URL Pattern           | Description                         |
|--------|------------------------|-------------------------------------|
| GET    | `/v1/movies`           | List all movies                     |
| POST   | `/v1/movies`           | Create a new movie                  |
| GET    | `/v1/movies/:id`       | Retrieve a specific movie           |
| PATCH  | `/v1/movies/:id`       | Partially update a movie            |
| DELETE | `/v1/movies/:id`       | Delete a movie                      |

### 👤 User Management

| Method | URL Pattern                  | Description                         |
|--------|-------------------------------|-------------------------------------|
| POST   | `/v1/users`                  | Register a new user                 |
| PUT    | `/v1/users/activated`        | Activate a user via token           |
| PUT    | `/v1/users/password`         | Change user's password              |

### 🔐 Authentication & Token Handling

| Method | URL Pattern                         | Description                              |
|--------|--------------------------------------|------------------------------------------|
| POST   | `/v1/tokens/authentication`         | Generate an authentication token         |
| POST   | `/v1/tokens/password-reset`         | Generate a password reset token          |

### ⚙️ System & Debugging

| Method | URL Pattern         | Description                          |
|--------|----------------------|--------------------------------------|
| GET    | `/v1/healthcheck`   | Check app health/version info        |
| GET    | `/debug/vars`       | View internal application metrics    |

---

## 🧠 What Was Learned

This project applies the complete curriculum from _Let's Go Further_, covering over 20 chapters of professional backend development practices in Go:

### 🧱 Architecture & Design
- Clean layering of handler, service, model, and DB layers
- Dependency injection using structured application state
- RESTful routing with path parameters and query parsing

### 🧾 JSON & Request Handling
- Custom JSON marshaling/unmarshaling
- Structured error responses and enveloped outputs
- Strong input validation and decoding safeguards

### 🛢️ PostgreSQL Integration
- Connection pooling with timeouts
- Optimistic concurrency control
- Full-text search and pagination metadata

### 🛡️ Security & Auth
- Secure password hashing with bcrypt
- Token-based authentication (Bearer scheme)
- Email-based account activation and password reset
- Fine-grained permission-based authorization

### 🚦 Middleware & Lifecycle
- Logging, panic recovery, and metrics middleware
- Context-aware graceful shutdown
- IP-based rate limiting with configurable limits

### 🛠️ DevOps & Production Readiness
- Config management via environment variables
- Structured logs and custom metrics with `expvar`
- SQL migrations and background worker pools
- Systemd services and reverse proxy with Caddy

---

## 🧪 Local Development

```bash
# Clone the project
git clone https://github.com/peeta98/greenlight.git
cd greenlight

# Setup PostgreSQL and apply SQL migrations (see /migrations folder)

# Export required env variables (DB_DSN, SMTP config, etc.)
source .envrc

# Run the API
go run ./cmd/api
```
