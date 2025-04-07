# Assignment 1 – Clean Architecture-Based Microservices
**Advanced Programming II**  
**Name:** Margulan Suyundykov
**Deadline:** 07.04.2025  
---

## 📦 Project Overview

This project implements a basic e-commerce platform using Clean Architecture principles and RESTful microservices. The system is composed of the following three main components:

1. **API Gateway** – Handles request routing, logging, telemetry, and authentication.
2. **Inventory Service** – Manages products, categories, stock levels, and prices.
3. **Order Service** – Manages order creation, status updates, and payment handling.

Each service is a standalone Go (Golang) application using the Gin framework and PostgreSQL as the database.

---

## 🏗️ Microservices & Architecture

All services are structured following Clean Architecture:
- `cmd/` – Application entrypoint
- `internal/` – Domain logic (usecases, entities, repositories, routes, handlers)
- `config/` – Configuration and environment loading
- `db/` – DB migrations and init scripts
- `Dockerfile` – For building and running containers

---

## 🚀 Endpoints

### Inventory Service
| Method | Endpoint           | Description                      |
|--------|--------------------|----------------------------------|
| POST   | /products          | Create a new product             |
| GET    | /products/:id      | Get a product by ID              |
| PATCH  | /products/:id      | Update product by ID             |
| DELETE | /products/:id      | Delete product by ID             |
| GET    | /products          | List all products (with filters) |

### Order Service
| Method | Endpoint           | Description                      |
|--------|--------------------|----------------------------------|
| POST   | /orders            | Create a new order               |
| GET    | /orders/:id        | Get order details by ID          |
| PATCH  | /orders/:id        | Update order status              |
| GET    | /orders            | List all orders                  |

### API Gateway
- Handles authentication (mock)
- Routes requests to services
- Exposes:
  - `/inventory/...`
  - `/orders/...`

---

## 🐳 Running with Docker

### 1. Clone the repository
```bash
git clone https://github.com/yourusername/ecommerce-microservices.git
cd ecommerce-microservices
