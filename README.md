# Fullstack Vite + React + Go Template

Production-oriented starter with a React SPA and a Go API.

## Features

- **Frontend:** Vite, React 19, TypeScript, Tailwind CSS v4, shadcn/ui (new-york), Zustand, TanStack Query, Axios, React Router (lazy routes, protected routes).
- **Backend:** Fiber, layered layout (`handler` → `service` → `repository`), GORM + PostgreSQL (optional at runtime), request logging, CORS, recovery middleware.
- **Config:** Root `.env.example` for documented variables; `frontend/.env.example` for `VITE_API_URL`.

## Project structure

- `frontend/` — React app (`src/app`, `src/features`, `src/components/ui`, `src/lib`, `src/store`).
- `backend/` — Go API (`cmd/server`, `internal/*`, `pkg/`, `migrations/`).

## Prerequisites

- Node.js 20+
- Go 1.23+
- PostgreSQL (optional; omit `DB_URL` to run the API without a database)

## Setup

1. Copy environment files:

   ```bash
   cp .env.example .env
   cp frontend/.env.example frontend/.env
   ```

2. Install frontend dependencies:

   ```bash
   cd frontend
   npm install
   ```

3. Backend dependencies resolve automatically with `go mod tidy` (run from `backend/` if needed).

## Run locally

**Terminal 1 — API**

```bash
cd backend
make run
```

**Terminal 2 — SPA**

```bash
cd frontend
npm run dev
```

With `VITE_API_URL` empty (default), the Vite dev server proxies `/api` and `/health` to `http://localhost:8080`, so the browser uses same-origin requests and CORS is not required for development.

## Build

**Frontend**

```bash
cd frontend
npm run build
```

**Backend**

```bash
cd backend
make build
```

## Testing

**Frontend**

```bash
cd frontend
npm run test
```

**Backend**

```bash
cd backend
go test ./...
```

## API

| Method | Path | Description |
|--------|------|-------------|
| GET | `/health` | Liveness JSON |
| POST | `/api/auth/login` | JSON body `{ "email", "password" }` → `{ "token" }` (demo token when `DB_URL` is unset) |

## Go module path

The module is `github.com/example/fullstack-template`. Replace it with your own module path and update imports under `backend/`.

## License

MIT
