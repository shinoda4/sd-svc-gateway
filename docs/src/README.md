# Introduction

Welcome to the **SD-SVC-GATEWAY** documentation.

The **SD-SVC-GATEWAY** is a lightweight, high-performance API gateway written in Go. It serves as the single entry point for the SD System, handling request routing, authentication, and proxying to downstream microservices.

## Key Features

- **Unified Entry Point**: All client requests go through the gateway, simplifying the client-side logic.
- **Authentication**: Validates JWT tokens using `AuthMiddleware` before forwarding requests to protected services.
- **Reverse Proxy**: Efficiently proxies requests to backend services like the Authentication Service.
- **Health Checks**: Provides a simple health check endpoint for monitoring.
- **Extensible**: Built with [Gin](https://github.com/gin-gonic/gin), making it easy to add custom middleware and routes.

## Getting Started

To get started with the gateway, check out the [Deployment](./deployment.md) guide.
