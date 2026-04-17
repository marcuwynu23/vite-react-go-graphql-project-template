# PLAN.md — vite-react-go-graphql-project-template

## Overview

Create a production-ready fullstack GraphQL template:

- Frontend: Vite + React + TypeScript
- Backend: Go (modular architecture)
- Styling: TailwindCSS + shadcn/ui
- State: Zustand
- Data Fetching: TanStack React Query + Axios
- Config: Centralized theme + environment variables
- Structure: Code-split, scalable, maintainable

---

## Goals

- Fast project bootstrap
- Scalable architecture (frontend + backend)
- Clean separation of concerns
- Reusable UI components
- Environment-driven configuration
- Easy deployment
- **Test-driven habits:** new features land with tests first (see [Testing Strategy (Test-Driven)](#testing-strategy-test-driven))

---

## Tech Stack

### Frontend

- Vite
- React + TypeScript
- TailwindCSS
- shadcn/ui
- Zustand
- TanStack React Query
- Axios

### Backend

- Go (Golang)
- Fiber + gqlgen
- ORM (GORM or Ent)
- Makefile
- dotenv (.env)

---

## Folder Structure

### Root

```
project-root/
  frontend/
  backend/
  README.md
  .env.example
```

---

## Frontend Architecture

```
frontend/
  src/
    app/
      providers/
      router/
    components/
      ui/            # shadcn components
      common/
    features/
      auth/
      dashboard/
    hooks/
    lib/
      axios.ts
      queryClient.ts
    store/
      useStore.ts
    styles/
      globals.css
      theme.ts
    types/
    utils/
    main.tsx
    App.tsx
```

---

## Frontend Implementation Steps

### 1. Initialize Project

```
npm create vite@latest frontend -- --template react-ts
cd frontend
npm install
```

### 2. Install Dependencies

```
npm install axios zustand @tanstack/react-query
npm install tailwindcss postcss autoprefixer
npx tailwindcss init -p
```

### 3. Setup Tailwind

- Configure `tailwind.config.ts`
- Add base styles

### 4. Setup shadcn

```
npx shadcn-ui@latest init
```

### 5. Setup Axios

Create `lib/axios.ts`

- Base URL from env
- Interceptors for auth/error

### 6. Setup React Query

Create `lib/queryClient.ts`

- Global config
- Retry + caching strategy

### 7. Setup Zustand

Create `store/useStore.ts`

- Global state slices

### 8. Centralized Theme

Create `styles/theme.ts`

- Colors
- Dark/light mode

### 9. Code Splitting

- Lazy load routes
- Feature-based structure

### 10. Routing

- Use React Router
- Protected routes

---

## Backend Architecture

```
backend/
  cmd/
    server/
      main.go
  internal/
    config/
    db/
    models/
    repository/
    service/
    handler/
    middleware/
  pkg/
  migrations/
  Makefile
  go.mod
  .env
```

---

## Backend Implementation Steps

### 1. Initialize Go Module

```
go mod init project
```

### 2. Install Dependencies

- ORM: GORM or Ent
- API layer: Fiber + gqlgen
- dotenv loader

### 3. Config Loader

- Load `.env`
- Central config struct

### 4. Database Setup

- Initialize ORM
- Connection pooling

### 5. Models

- Define entities
- Use ORM tags

### 6. Repository Layer

- Abstract DB operations

### 7. Service Layer

- Business logic

### 8. Handler Layer

- HTTP handlers

### 9. Middleware

- Logging
- Auth
- CORS

### 10. Makefile

Example:

```
run:
	go run ./cmd/server

build:
	go build -o app ./cmd/server

migrate:
	go run ./migrations
```

---

## Environment Variables

### Root `.env.example`

```
# Frontend
VITE_API_URL=http://localhost:8080

# Backend
PORT=8080
DB_URL=postgres://user:pass@localhost:5432/db
```

---

## Code Quality

- ESLint + Prettier (frontend)
- golangci-lint (backend)
- Consistent naming conventions

---

## Testing Strategy (Test-Driven)

Development follows **test-driven development (TDD)** where practical: write a **failing test** first (red), implement the **minimum code** to pass (green), then **refactor** without changing behavior. For UI and HTTP integration, prefer **behavior-focused tests** (user-visible outcomes, status codes, JSON shapes) over implementation details.

### Tooling

| Area | Stack |
|------|--------|
| Frontend | Vitest, React Testing Library, MSW (optional) for API mocking |
| Backend | `testing`, `httptest`, table-driven tests; optional `testify` for assertions |

### TDD workflow (feature work)

1. **Specify** — Add or update a test case in this plan (or a short `*_test` spec comment) describing input → expected behavior.
2. **Red** — Write the test; run it; confirm it fails for the right reason.
3. **Green** — Implement production code until the test passes.
4. **Refactor** — Clean up; keep tests green.
5. **Regression** — Run full suites: `npm run test` and `go test ./...`.

### Frontend test cases

Co-locate tests as `*.test.ts` / `*.test.tsx` next to source or under `src/__tests__/` consistently.

| ID | Area | Behavior to assert |
|----|------|---------------------|
| FE-01 | `lib/axios` | Base URL respects `VITE_API_URL` (including empty → relative origin for proxy). |
| FE-02 | `lib/axios` | Request interceptor attaches `Authorization: Bearer <token>` when `auth_token` exists in `localStorage`. |
| FE-03 | `lib/axios` | On 401 response, `auth_token` is cleared from `localStorage`. |
| FE-04 | `lib/queryClient` | Default query options match policy (e.g. stale time, retry rules for 4xx vs network). |
| FE-05 | `store/useStore` | `setToken` persists token and clears it consistently with axios expectations. |
| FE-06 | Router | Unauthenticated user hitting a protected route is redirected to login (with return URL preserved where applicable). |
| FE-07 | `features/auth/LoginPage` | Successful login calls API, stores token, navigates to dashboard (mock server or MSW). |
| FE-08 | `features/auth/LoginPage` | API error surfaces an error state; demo / offline path still works if specified. |
| FE-09 | `features/dashboard/DashboardPage` | Health query renders loading, then success JSON, or error when API unavailable (mocked). |
| FE-10 | UI (shadcn) | Critical interactive components (e.g. `Button`) meet a11y expectations: disabled state, `type="submit"` in forms. |

**Frontend test commands**

```bash
cd frontend
npm run test          # CI: vitest run
npm run test -- --watch   # local TDD loop
```

### Backend test cases

Use **table-driven** tests in `*_test.go` files beside packages under `internal/`.

| ID | Package / symbol | Behavior to assert |
|----|------------------|---------------------|
| BE-01 | `handler.Health` | `GET /health` returns `200` and JSON body includes `status` and `service` keys. |
| BE-02 | GraphQL `login` mutation | Invalid input returns GraphQL validation/runtime error shape. |
| BE-03 | GraphQL `login` mutation | Valid input returns `data.login.token` (non-empty) when DB is disabled / demo mode. |
| BE-04 | GraphQL `login` mutation | Invalid credentials surface an auth error when user store is enabled. |
| BE-05 | `service.AuthService` | Login with repository disabled issues a token without error. |
| BE-06 | `service.AuthService` | Login with repository enabled enforces validation rules (e.g. empty password → invalid credentials). |
| BE-07 | `repository.UserRepository` | `FindByEmail` returns `ErrRecordNotFound` when DB is nil or user missing. |
| BE-08 | `middleware` / CORS | Preflight `OPTIONS` from allowed `CORS_ORIGINS` succeeds when tested via `app.Test`. |
| BE-09 | `config` | Required defaults (`PORT`, `CORS_ORIGINS`) apply when env vars are unset. |

**Backend test commands**

```bash
cd backend
go test ./... -count=1
go test -race ./...    # optional CI
```

### Order of implementation (TDD-friendly)

1. **Backend:** BE-01 → BE-02 → BE-03 (establish contract the frontend will call).
2. **Frontend:** FE-01 → FE-07 → FE-09 (client matches API contract).
3. Fill in BE-04–BE-07 and FE-02–FE-06, FE-08, FE-10 as auth and UX harden.

---

## README.md (Root)

```
# vite-react-go-graphql-project-template

## Features
- Vite + React + TypeScript
- TailwindCSS + shadcn/ui
- Zustand (state management)
- React Query + Axios
- Go backend with clean architecture + GraphQL (`gqlgen`)
- ORM-based database access
- Environment-based config

## Project Structure
- frontend/ — React app
- backend/ — Go GraphQL API

## Getting Started

### Prerequisites
- Node.js
- Go
- PostgreSQL (optional)

### Setup

1. Clone repo
2. Copy env file
```

cp .env.example .env

```

### Run Frontend
```

cd frontend
npm install
npm run dev

```

### Run Backend
```

cd backend
make run

```

## Build

### Frontend
```

npm run build

```

### Backend
```

make build

```

## Testing

### Frontend
```

npm run test

```

### Backend
```

go test ./...

```

## Future Improvements
- Docker support
- CI/CD pipeline
- Auth system
- Role-based access control

## License
MIT
```

---

## Acceptance Criteria

- Frontend runs with dev server
- Backend runs via Makefile
- API communication works
- Env variables properly loaded
- Clean modular structure
- **Frontend:** `npm run test` passes; core cases FE-01–FE-09 covered or explicitly deferred with a tracked TODO
- **Backend:** `go test ./...` passes; core cases BE-01–BE-07 covered for handlers and auth/service paths in use

---

## Deliverables

- Fully working template
- Clean folder structure
- Ready-to-use boilerplate
- Documentation (README)
- **Test catalog** implemented or scheduled: frontend IDs `FE-*`, backend IDs `BE-*` in this document reflect runnable tests or the next sprint backlog
