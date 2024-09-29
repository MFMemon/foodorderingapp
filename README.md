# FoodOrderingApp Using MicroServices

## Local

### Docker Compose

Spin up the docker containers for your services, DBs', Consul service discovery and API Gateway
```bash
docker compose up
```

### Test Login Endpoint
```bash
curl http://localhost:6500/user/login/?email=<>&password=<>
```