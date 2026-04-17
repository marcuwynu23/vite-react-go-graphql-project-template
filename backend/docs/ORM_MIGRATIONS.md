# ORM Migration Guide (GORM AutoMigrate)

This project uses **pure ORM migrations** with GORM `AutoMigrate` (no SQL migration files).

## Where migration runs

Migration is executed in two places:

- `cmd/server/main.go` on API startup
- `cmd/dbtool/main.go` for CLI commands (`make migrate`, `make bootstrap`)

Current migration call:

- `gormDB.AutoMigrate(&models.User{})`

## Commands

Run from `backend/`:

- `make migrate` - run schema sync only
- `make seed` - run migration + seed initial user
- `make bootstrap` - run migration + seed (same as seed for now)
- `make run` - start API (also runs migration)

## How to add a new model migration

### 1) Create or update model struct

Add a new model file in `internal/models`, for example:

```go
package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserID uint   `gorm:"not null;index"`
	Bio    string `gorm:"type:text"`
}
```

### 2) Register model in AutoMigrate

Update both files:

- `cmd/server/main.go`
- `cmd/dbtool/main.go`

Example:

```go
if err := gormDB.AutoMigrate(&models.User{}, &models.Profile{}); err != nil {
	log.Fatalf("migrate: %v", err)
}
```

### 3) Run migration locally

```bash
cd backend
make migrate
```

### 4) Verify

- Check DB tables (SQLite file `app.db` by default)
- Run tests:

```bash
go test ./...
```

## Notes and limits

- `AutoMigrate` handles table/column/index creation and safe alterations.
- It **does not** safely cover every destructive/complex schema change.
- For complex production changes (data backfills, column type rewrites), use a planned manual migration process.

## SQLite vs PostgreSQL

- `DB_URL` empty -> SQLite (`app.db`)
- `DB_URL=postgres://...` or `postgresql://...` -> PostgreSQL

Migration command stays the same in both cases (`make migrate`), because GORM selects the dialect through `DB_URL`.
