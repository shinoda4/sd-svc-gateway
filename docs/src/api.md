# API Reference

The gateway exposes several endpoints, categorized into public and protected routes.

## Public Endpoints

These endpoints do not require authentication.

### Health Check
Checks if the gateway service is running.

- **URL**: `/health`
- **Method**: `GET`
- **Response**:
  ```json
  {
    "status": "ok"
  }
  ```

### Authentication Service Proxy
All requests starting with `/auth` are forwarded to the configured Authentication Service.

- **URL**: `/auth/*`
- **Method**: `ANY`
- **Description**: Proxies requests to the Auth Service. For example, a request to `/auth/login` on the gateway is forwarded to `/login` (or `/auth/login` depending on upstream config) on the Auth Service.

## Protected Endpoints

These endpoints require a valid JWT in the `Authorization` header.
**Header Format**: `Authorization: Bearer <token>`

### Get Current User
Returns the information of the currently authenticated user.

- **URL**: `/api/me`
- **Method**: `GET`
- **Response**:
  ```json
  {
    "message": "You are authenticated",
    "userID": "user_123",
    "email": "user@example.com"
  }
  ```
- **Errors**:
  - `401 Unauthorized`: Missing or invalid token.
