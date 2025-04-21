# Golang API Example

## Run application

```
docker build -t go-api .
docker run -p 8080:8080 go-api

```


Module 1: Course Introduction
Duration: ~3 hours

🎯 Learning Objectives

By the end of this module, you will:

✅ Understand the course structure and expected outcomes.

✅ Set up a Golang development environment with Devcontainers in VS Code.

✅ Create and run a basic "Hello World" REST API in Go.

✅ Containerize the API with Docker.

✅ Understand key differences between Go and Python for API development.


1️⃣ Course Roadmap Overview

We’ll be building a high-availability, scalable API using:

✅ Golang for API development.

✅ MongoDB for persistence.

✅ RabbitMQ for async processing.

✅ Redis for caching.

✅ OpenTelemetry for observability.

✅ Docker & Kubernetes for deployment.

📌 Strategy: We’ll start simple (local "Hello World" API) and iterate towards a production-grade solution with authentication, logging, monitoring, and high availability.

2️⃣ Setting Up the Development Environment

✅ Install & Configure the Tools

Make sure you have the following installed:

- Golang (latest version)

- VS Code + Devcontainers Extension

- Docker