# Backend Guide

This backend provides a Go + Fiber + GraphQL (`gqlgen`) API with GORM-based persistence.

## Stack

- Go
- Fiber
- GraphQL (`gqlgen`)
- GORM
- SQLite by default (`app.db`)
- PostgreSQL optional via `DB_URL`

## Prerequisites

- Go 1.25+
- `make`

## Configuration

Backend reads environment values from root `.env` (via `godotenv`).

Common variables:

- `PORT` (default: `8080`)
- `ENV` (default: `development`)
- `DB_URL` (default empty; empty means SQLite file `app.db`)
- `CORS_ORIGINS` (default: `http://localhost:5173`)
- `SEED_EMAIL` (default: `admin@example.com`)
- `SEED_PASSWORD` (default: `admin123`)

### DB behavior

- If `DB_URL` is empty: uses SQLite (`app.db`) in `backend/`
- If `DB_URL` starts with `postgres://` or `postgresql://`: uses PostgreSQL

Example PostgreSQL:

```env
DB_URL=postgres://user:pass@localhost:5432/mydb?sslmode=disable
```

## Makefile Commands

Run these from `backend/`.

- `make help`  
  Show all available targets.

- `make run`  
  Start API server on `PORT` (default `8080`).

- `make build`  
  Build binary at `bin/server`.

- `make test`  
  Run `go test ./...`.

- `make tidy`  
  Run `go mod tidy`.

- `make orm-status`  
  Print effective DB + seed values.

- `make migrate`  
  Run GORM `AutoMigrate` for current models.

- `make seed`  
  Seed initial login user from `SEED_EMAIL` / `SEED_PASSWORD` (safe/idempotent).

- `make bootstrap`  
  Run migrate + seed together.

- `make reset-sqlite`  
  Delete local SQLite file (`app.db`) for a clean local reset.

## Typical Local Workflow

1. Check config:

   ```bash
   make orm-status
   ```

2. Create/update schema:

   ```bash
   make migrate
   ```

3. Seed initial user:

   ```bash
   make seed
   ```

4. Start API:

   ```bash
   make run
   ```

5. Open playground:

   - [http://localhost:8080/playground](http://localhost:8080/playground)

## ORM Migrations

Schema synchronization is managed with GORM `AutoMigrate` in:

- `cmd/server/main.go` (startup)
- `cmd/dbtool/main.go` via `make migrate` / `make bootstrap`

For new models, add your struct under `internal/models` and include it in the `AutoMigrate(...)` calls.

Detailed guide: see `docs/ORM_MIGRATIONS.md`.

## GraphQL Endpoints

- `POST /graphql` — GraphQL API endpoint
- `GET /playground` — local playground UI

Detailed guide: see `docs/GRAPHQL_GUIDE.md`.

## Auth Quick Test

Use the seeded credentials in the `login` mutation:

```graphql
mutation Login($input: LoginInput!) {
  login(input: $input) {
    token
  }
}
```

Variables:

```json
{
  "input": {
    "email": "admin@example.com",
    "password": "admin123"
  }
}
```
