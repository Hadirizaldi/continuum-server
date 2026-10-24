# Continuum Server

Modular resumable upload system — Go microservice with chunked upload, pause/resume, and pluggable storage & database adapters.

## Tech Stack

- **Language:** Go
- **HTTP Router:** Chi
- **Database:** PostgreSQL (pgx + pgxpool)
- **Migration:** golang-migrate
- **Storage:** Local disk (MVP)
- **Config:** Go structs + godotenv (.env override)
- **UUID:** google/uuid

## Project Structure

```
cmd/server/main.go          # entrypoint
internal/
  config/                   # configuration structs & loader
  adapter/                  # StorageAdapter & DatabaseAdapter interfaces
  postgres/                 # PostgreSQL adapter implementation
  localstorage/             # local disk storage implementation
  service/                  # core business logic
  handler/                  # HTTP handlers
migrations/                 # SQL migration files (golang-migrate format)
```

## Commands

```bash
# Run server
go run ./cmd/server

# Run migrations
migrate -path migrations -database "postgres://postgres:password@localhost:5432/uploads?sslmode=disable" up

# Run tests
go test ./...
```

## Conventions

- All code in English (variables, functions, comments, commit messages)
- Follow standard Go conventions (effective Go, Go code review comments)
- Always pass `context.Context` as the first parameter
- Return `error` as the last return value
- Wrap errors with `fmt.Errorf("operation: %w", err)` for context
- Use table-driven tests
- No global mutable state; inject dependencies via constructor
