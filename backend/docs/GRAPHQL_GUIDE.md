# GraphQL Guide (gqlgen)

This backend uses `gqlgen` for GraphQL on top of Fiber.

## Current endpoints

- `POST /graphql` - GraphQL API endpoint
- `GET /playground` - local GraphQL Playground

Run server:

```bash
cd backend
make run
```

Then open:

- [http://localhost:8080/playground](http://localhost:8080/playground)

## GraphQL file layout

- Schema: `internal/graphql/schema.graphqls`
- Generated runtime: `internal/graphql/generated/generated.go`
- Generated models: `internal/graphql/model/models_gen.go`
- Resolver root: `internal/graphql/resolver.go`
- Resolver implementations: `internal/graphql/schema.resolvers.go`
- Generator config: `gqlgen.yml`

## Request flow

1. Client sends GraphQL operation to `/graphql`
2. gqlgen maps operation to resolver function
3. Resolver calls service layer (`internal/service`)
4. Service calls repository (`internal/repository`) and returns business result
5. Resolver maps service result to GraphQL model

## Add a new GraphQL field (query or mutation)

### 1) Update schema

Edit `internal/graphql/schema.graphqls`.

Example query:

```graphql
type Query {
  health: Health!
  me: User!
}
```

### 2) Regenerate gqlgen code

From `backend/`:

```bash
go run github.com/99designs/gqlgen generate
```

After generation, gqlgen updates generated files and may add resolver stubs.

### 3) Implement resolver

Add implementation in `internal/graphql/schema.resolvers.go`.

Resolver should:

- validate input
- call service layer
- map output to GraphQL model
- return typed errors when needed

### 4) Add/extend service + repository

Keep business logic in `internal/service`, DB access in `internal/repository`.

Avoid placing DB logic directly inside resolvers.

### 5) Test in playground

Run:

```bash
make run
```

Open playground and execute operation with variables.

## Example operations

### Health query

```graphql
query Health {
  health {
    status
    service
  }
}
```

### Login mutation

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

## Error handling notes

- Resolver errors are returned in GraphQL `errors` array.
- Authentication/validation errors should come from service layer and be surfaced clearly.

## Common commands

From `backend/`:

- `make run` - start API server
- `make test` - run backend tests
- `make migrate` - ORM schema sync (AutoMigrate)
- `make seed` - ensure initial login user exists

## Tips

- Keep schema changes backward compatible when possible.
- Generate code immediately after schema edits to avoid drift.
- Keep resolver functions thin; push logic down to service/repository.
