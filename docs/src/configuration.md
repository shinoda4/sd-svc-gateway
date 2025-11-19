# Configuration

The **SD-SVC-GATEWAY** is configured using environment variables. You can set these variables in your shell or use a `.env` file if you are running with a tool that supports it (like `make` or `docker-compose`).

## Environment Variables

| Variable | Description | Default Value |
|----------|-------------|---------------|
| `PORT` | The port on which the gateway server listens. | `8080` |
| `AUTH_SVC_URL` | The full URL of the upstream Authentication Service. | `http://localhost:8081` |
| `JWT_SECRET` | The secret key used to validate JWT signatures. **Must match the one used by the Auth Service.** | `change_me` |

## Example Configuration

```bash
export PORT=8000
export AUTH_SVC_URL=http://auth-service:8081
export JWT_SECRET=super_secret_key
```
