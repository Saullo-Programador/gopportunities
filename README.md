# GoOpportunities 🚀

[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat-square&logo=go)](https://go.dev/)
[![Framework](https://img.shields.io/badge/Framework-Gin-00ADD8?style=flat-square)](https://gin-gonic.com/)
[![Database](https://img.shields.io/badge/Database-SQLite-003B57?style=flat-square&logo=sqlite)](https://www.sqlite.org/)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](LICENSE)

GoOpportunities is a high-performance REST API developed in Go for job opportunity management. It provides a scalable backend for recruitment platforms, allowing full CRUD operations on job openings with a focus on clean code and developer experience.

## ✨ Key Features

- **Full CRUD Lifecycle**: Complete management of job openings (Create, Read, Update, Delete).
- **Interactive Documentation**: Automated API documentation using Swagger/OpenAPI for seamless integration.
- **Efficient Persistence**: Leveraging GORM for optimized database interactions and schema migrations.
- **Structured Logging**: Centralized logging system for better observability and debugging.
- **Type-Safe Schemas**: Strictly defined data models ensuring API consistency.

## 🛠️ Tech Stack

| Technology | Purpose | Justification |
| :--- | :--- | :--- |
| **Go** | Language | Selected for its superior concurrency model and runtime efficiency. |
| **Gin Gonic** | Web Framework | Chosen for its minimalist approach and high-speed routing. |
| **GORM** | ORM | Provides a powerful abstraction layer for database operations. |
| **SQLite** | Database | Lightweight and portable, ideal for rapid development and testing. |
| **Swagger** | Documentation | Standardizes the API contract for frontend and third-party consumers. |

## 🏗️ Architecture

The project follows a modular architecture to ensure separation of concerns and maintainability:

- `config/`: System initialization, environment variables, and database connection.
- `handler/`: Controller layer responsible for request parsing and response orchestration.
- `router/`: Routing definitions and middleware integration.
- `schemas/`: Domain entities and database models (GORM).
- `docs/`: Auto-generated OpenAPI specifications.

## 🚀 Getting Started

### Prerequisites
- Go 1.20+
- Git

### Installation & Setup
1. **Clone the repository**:
   ```bash
   git clone https://github.com/Saullo-Programador/gopportunities.git
   cd gopportunities
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Run the application**:
   ```bash
   go run main.go
   ```

The API will be available at `http://localhost:8080`.

## 📖 API Reference

Access the interactive Swagger UI to explore and test endpoints:
👉 `http://localhost:8080/swagger/index.html`

### Main Endpoints

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/api/v1/openings` | List all available opportunities |
| `GET` | `/api/v1/opening` | Get details of a specific opening |
| `POST` | `/api/v1/opening` | Create a new job opportunity |
| `PUT` | `/api/v1/opening` | Update existing opening data |
| `DELETE` | `/api/v1/opening` | Remove an opening from the system |

## 🗺️ Roadmap

- [ ] **Authentication**: Implement JWT (JSON Web Tokens) for secure access.
- [ ] **Database Migration**: Support for PostgreSQL for production environments.
- [ ] **Unit Testing**: Implement test suites for handlers and services.
- [ ] **Dockerization**: Create a Docker image for easier deployment.
- [ ] **CI/CD**: Setup GitHub Actions for automated linting and testing.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---
Developed by [Saullo Programador](https://github.com/Saullo-Programador)
