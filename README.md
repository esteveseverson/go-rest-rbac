# RBAC with Golang
The goal is to make a Go REST API with role-based access control. The objective is to create an email sender, manipulate images, JWT authentication, password hashing. Technologies will be used: _Golang_, _SQLC_ for _Postgres_, _dbmate_ for migrations and _Redis_ for caching.

Keywords:
- Go, Golang
- SQLC
- dbmate
- Postgres
- Hashing
- API REST
- JWT

## .env example (necessary to run)
Create .env file in root folder with:
```bash
# Server
SERVER_PORT=
DATABASE_URL=
ENVIRONMENT=
LOG_LEVEL=

# Database
POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_DB=

```

## Run Database
```bash
docker compose up -d
```

## Run project
```bash
go mod tidy
go run main.go
```