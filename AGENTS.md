# Gymtrack - Documentation and Standards for AI Agents (AGENTS.md)

This document describes the architecture, technologies, and code standards of the Gymtrack project. It must be used as a strict rules guide for any AI or development agent modifying the codebase.

## 1. Overview and Architecture

- **Monorepo (`go.work`):** The project is organized as a monorepo using Go Workspaces.
  - Microservices and applications are located in the `apps/` directory (e.g. `apps/api/identity`, `apps/api/training-plan`).
  - Shared libraries and modules are located in the `libs/` directory (e.g. `libs/auth`, `libs/db`).

- **Clean Architecture / DDD:** Each context or entity inside a microservice (located in the `internal/` folder) follows a clear separation of responsibilities:
  - `handler.go`: HTTP layer. Receives requests, validates payloads, and returns HTTP responses. It must not contain complex business rules.
  - `service.go`: Application and business rules layer. Validates domain logic and orchestrates repository calls.
  - `repository.go`: Infrastructure and persistence layer. Exclusively responsible for direct communication with the database.
  - `domain/`: Contains entity definitions (models), constants, enums, and domain error declarations.

- **Dependency Injection (DI):** The project heavily uses the `go.uber.org/fx` library to orchestrate application startup, dependency injection, and lifecycle management in `cmd/main.go`. New components must be registered in the `fx` container.

## 2. Technology Stack and Tools

- **Language:** Go
- **Web Framework:** Gin (`github.com/gin-gonic/gin`)
- **Database:** PostgreSQL
  - **VERY IMPORTANT:** Data access is implemented using the native `database/sql` package together with the `github.com/lib/pq` driver. **Do not use ORMs** (such as GORM, Ent, etc). All queries must be written in raw SQL.
- **Authentication:** JWT-based, encapsulated and provided by the internal `libs/auth` library.
- **Logging:** Uses the standard library `log/slog`, configured through the internal `libs/log` package.

## 3. Code Standards and Strict Conventions

### Routing and HTTP Handlers

- Handlers must implement the `RegisterRoutes(r *gin.Engine)` method or a similar signature (e.g. receiving a `*gin.RouterGroup`) to register their own routes in an isolated manner.

### Context Management (`context.Context`)

- The web request context MUST be extracted in the handler (`ctx.Request.Context()`) and passed as the first parameter to all `Service` and `Repository` layer functions. This guarantees support for propagated timeouts and cancellations.

### Payloads and Data Validation

- Do not use database/domain structs to receive request payloads.
- Declare anonymous structs or dedicated DTO (Data Transfer Object) structs directly within the handler scope.
- Validate input data using Gin binding tags (e.g. `` binding:"required,min=1" ``) together with `ctx.ShouldBindJSON()`.

### Authentication and Current User

- In routes protected by the `auth.AuthMiddleware()` middleware, logged-in user information must be extracted exclusively using the helper below:

```go
userVal, ok := auth.GetAuthUser(ctx)
if !ok {
    return // The helper already handles the error response internally
}
```

- When calling internal microservices that require authentication, the bearer token should be extracted directly from the context in the Service layer using:
```go
token, _ := ctx.Value(string(auth.TokenContextKey)).(string)
```
Do not extract the token in the Handler layer to pass it as a parameter to the Service.

### Error Handling and Responses
- Domain Errors: Expected business rule errors (e.g. "user not found", "invalid password") must be declared as global variables or specific types inside the domain package (e.g. domain.ErrUserNotFound).
- HTTP Mapping: In the Handler layer, inspect errors returned by the Service using errors.Is(err, domain.Err...) to map them to the correct HTTP Status Code (400 Bad Request, 401 Unauthorized, 403 Forbidden, 404 Not Found, etc).
- Unexpected Error Logging: If the error is an internal database error or an unmapped failure, use slog.ErrorContext to log the original error before returning an HTTP 500 response to the client.

### Database Migrations
- Any database schema change (tables, columns, indexes) REQUIRES the creation of .up.sql files inside the migrations/ folder of the respective microservice. NEVER modify table schemas through code or direct commands during runtime.

### UUID Generation
- NEVER generate UUIDs (IDs) at the database level (e.g., using `DEFAULT uuid_generate_v4()`).
- All IDs must be generated within the application's Service layer using the utility `libs/utils/uuid.go` (specifically `utils.GenerateUUIDV7String`).

## 4. Expected Service Structure

Example:

apps/api/identity/
├── cmd/
│   └── main.go
├── internal/
│   └── user/
│       ├── handler.go
│       ├── service.go
│       ├── repository.go
│       └── domain/
│           ├── model.go
│           └── errors.go
├── migrations/
├── go.sum
└── go.mod

## 5. Naming Conventions

- Interfaces should end with `Repository`, `Service`, or `Provider`.
- Tests files must end with `_test`.
- Mock files must begin with `mock_`.
- Database tables must use `snake_case`.
- Go package names must be lowercase and singular when possible.

## 6. Testing

- Every new business rule must include unit tests.
- Prefer table-driven tests in Go.
- Use `gomock`(`go.uber.org/mock/gomock`) to create repository mocks
- Add generate to directive to mocks repositories/clients, example:
```go
//go:generate go run go.uber.org/mock/mockgen -source=repository.go -destination=mock_repository.go -package=post
type Repository interface {
	Create(ctx context.Context, post *domain.Post) error
}
```
- Always run `go build` and `go test`

## 7. Observability

- All errors must include contextual logging.
- Logs should include request IDs when available.
- Avoid logging sensitive information such as passwords, JWTs, or tokens.
- External API calls should include latency metrics when possible.

## 8. Performance Guidelines

- Avoid unnecessary allocations in hot paths.
- Prefer streaming over loading large datasets into memory.
- Goroutines must always have bounded lifecycles.
- Context cancellation must always be respected.

## 9. Forbidden Patterns

- Do not introduce global mutable state.
- Do not access the database directly from handlers.
- Do not place business logic inside HTTP handlers.
- Do not use panic for expected errors.
- Do not introduce new external dependencies without necessity.
- Do not use reflection-based magic libraries.

## 10. Concurrency

- Shared state must be protected properly.
- Never start goroutines without cancellation strategy.
- Use `errgroup` when coordinating concurrent operations.
- Channels must not leak goroutines.

## 11. API Standards

- All APIs must return JSON.
- Error responses must follow the format:

```json
{
  "error": "user not found"
}
```
- Use cursor pagination for list endpoints.

### API Response Guidelines

- `POST` and `PUT` endpoints should avoid returning full resource objects unless strictly necessary.
- Prefer minimal success responses such as:
  - `201 Created`
  - `200 OK`
- Full entity responses should only be returned when the client immediately requires the updated/generated data.
- Avoid unnecessary database re-queries only to build response payloads.