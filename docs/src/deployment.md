# Deployment

This guide covers how to build and run the **SD-SVC-GATEWAY**.

## Prerequisites

- **Go**: Version 1.21 or higher.
- **Make**: (Optional) For running convenience scripts.

## Building the Service

To build the binary, run:

```bash
go build -o bin/gateway cmd/gateway/main.go
```

Or using `make`:

```bash
make build
```

## Running Locally

1. **Set Environment Variables**: Ensure you have the necessary config set.
   ```bash
   export JWT_SECRET=mysecret
   ```

2. **Run the Binary**:
   ```bash
   ./bin/gateway
   ```

   Or using `make`:
   ```bash
   make run
   ```

## Docker

To build a Docker image:

```bash
docker build -t sd-svc-gateway:latest .
```

To run the container:

```bash
docker run -p 8080:8080 \
  -e AUTH_SVC_URL=http://host.docker.internal:8081 \
  -e JWT_SECRET=mysecret \
  sd-svc-gateway:latest
```
