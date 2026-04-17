.PHONY: help run build test tidy gqlgen migrate seed bootstrap orm-status reset-sqlite

help:
	@echo "Targets:"
	@echo "  make run          - run GraphQL API server"
	@echo "  make build        - build backend binary (bin/server)"
	@echo "  make test         - run all Go tests"
	@echo "  make tidy         - tidy go modules"
	@echo "  make gqlgen       - regenerate GraphQL generated files"
	@echo "  make migrate      - run GORM AutoMigrate (ORM schema sync)"
	@echo "  make seed         - seed initial user from SEED_EMAIL/SEED_PASSWORD"
	@echo "  make bootstrap    - migrate + seed in one command"
	@echo "  make orm-status   - print active DB settings"
	@echo "  make reset-sqlite - delete local sqlite db file (app.db)"

run:
	go run ./cmd/server

build:
	go build -o bin/server ./cmd/server

test:
	go test ./...

tidy:
	go mod tidy

gqlgen:
	go get github.com/99designs/gqlgen@v0.17.89
	go run github.com/99designs/gqlgen generate

migrate:
	go run ./cmd/dbtool migrate

seed:
	go run ./cmd/dbtool seed

bootstrap:
	go run ./cmd/dbtool bootstrap

orm-status:
	@echo "DB_URL=$${DB_URL:-<empty -> sqlite app.db>}"
	@echo "SEED_EMAIL=$${SEED_EMAIL:-admin@example.com}"
	@echo "SEED_PASSWORD=$${SEED_PASSWORD:-admin123}"

reset-sqlite:
	powershell -NoProfile -Command "if (Test-Path 'app.db') { Remove-Item 'app.db'; Write-Host 'removed app.db' } else { Write-Host 'app.db not found' }"
