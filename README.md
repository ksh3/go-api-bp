# 📐 Architecture Design Guidelines

This document defines the architectural layering, naming conventions, and dependency rules used in this project to maximize maintainability, scalability, and readability.

---

## 🔁 Layered Architecture and Responsibilities

The project is primarily divided into the following 3 layers:

```
UI (event, state) ↓
presentation (gateway) ↓
domain (entity, usecase, adapter, repository) ↓
infrastructure (service, repository impl, resource)
```

| Layer            | Responsibilities                                                         |
|------------------|--------------------------------------------------------------------------|
| `presentation`   | Handles UI connection, input events, output states, UI-specific gateways |
| `domain`         | Business rules, use cases, entities, adapter and repository interfaces   |
| `infrastructure` | Implements external services, adapters, repositories, and data access    |

---

## 🧱 Directory Structure and Role

```
├── domain
│   ├── adapter.go             # Interface definitions for external dependencies
│   ├── repository.go          # Interface definitions for data repositories
│   ├── entity/                # Entities
│   └── usecase/               # Use case implementations
│
├── infrastructure
│   ├── service.go             # Implementations of Adapters (e.g. API clients)
│   ├── dto.go                 # DTOs for internal/external data mapping
│   ├── repository/            # Concrete repository implementations (e.g. user_repo_impl.go)
│   └── data_resource/              # Data sources (API, files, local etc.)
│
├── presentation
│   ├── gateway.go             # UI-oriented external integration (e.g. calendar display)
│   ├── event.go               # Input events
│   ├── state.go               # Output states
│   └── api_response.go        # API response format mapping
│   └── web_response.go        # API response format mapping

├── core
│   ├── contract/              # Generic reusable interfaces and base contracts
│   │   ├── comparable.go      # Comparable[T] - defines equality check
│   │   ├── copyble.go         # Copyble[T] - defines immutable copy method
│   │   ├── dto.go             # DTO[T] - combines Copyble + Serializable
│   │   ├── error.go           # AppError - application-wide error type
│   │   ├── result.go          # Result[T] - either-style result wrapper
│   │   └── serializable.go    # Serializable[T] - for JSON conversion
```

---

## 🔌 Naming Conventions for Adapter / Gateway / Service / Repository

| Name           | Description                                     | Location                        |
|----------------|-------------------------------------------------|---------------------------------|
| `Adapter`      | Interface used by domain for external access    | `domain/adapter.go`             |
| `Service`      | Adapter implementation (technical / API calls)  | `infrastructure/`               |
| `Gateway`      | Presentation-specific API interface             | `presentation/`                 |
| `Repository`   | Repository interface for domain layer           | `domain/repository.go`          |
| `xxRepoImpl`   | Concrete repository implementation              | `infrastructure/repository/`    |

> Note: Even for the same API, implementations are separated by purpose (display vs business logic).

---

## 💡 Dependency Injection Policy

- `usecase` depends only on `Adapter` and `Repository` interfaces
- Concrete `Service`, `Gateway`, and `RepoImpl` are injected via DI (e.g. dig)

```go
container.Provide(func() domain.HolidayAdapter {
    return &infra.GoogleHolidayService{...}
})

container.Provide(func() domain.UserRepository {
    return &infra.UserRepoImpl{...}
})
```

---

## 📌 Naming Examples

| Interface (Abstract) | Concrete Implementation       | Purpose                      |
|----------------------|-------------------------------|------------------------------|
| `HolidayAdapter`     | `GoogleHolidayService`        | Holiday fetching (business)  |
|                      | `CalendarGateway`             | Holiday display (UI)         |
| `MailerAdapter`      | `SendGridService`             | Email sending                |
| `UserRepository`     | `UserRepoImpl`                | User data persistence        |

---

## 📦 `contract` Package Design Policy

### Purpose
The `contract` package defines generic reusable interfaces and types used across the app to enforce consistency, safety, and clarity in design.

### Interfaces

#### `Comparable`
```go
type Comparable[T any] interface {
	Equals(other T) bool
}
```
Used for equality checks on value objects or entities.

#### `Copyble`
```go
type Copyble[T any] interface {
	CopyWith(opts T) T
}
```
Supports immutable updates by copying with overrides.

#### `Serializable`
```go
type Serializable[T any] interface {
	ToJSON() (string, error)
	FromJSON(jsonStr string) (T, error)
}
```
Used for converting objects to/from JSON strings.

#### `DTO`
```go
type DTO[T any] interface {
	Copyble[T]
	Serializable[T]
}
```
Composite interface representing a copyable and serializable object, often used for transport or persistence.

### Errors

#### `AppError`
```go
type AppError struct {
	Code       int
	Message    string
	Cause      error
	IsCritical bool
}
```
General-purpose error structure.

```go
func ValidationError(code int, message string, cause error) *AppError
func InternalError(message string, cause error) *AppError
```
Predefined constructors for soft (validation) and critical (system) errors.

### Result Wrapper

#### `Result`
```go
type Result[T any] struct {
	Value T
	Err   error
}
```
A generic either-like result wrapper for representing success or failure.

```go
func Success[T any](value T) Result[T]
func Failure[T any](err error) Result[T]
```

---

## ✅ Summary of Design Principles

- Keep all abstractions (Adapters, Repositories) in `domain`
- Place implementations (Service, Gateway, RepoImpl) in `infrastructure` or `presentation`
- Separate implementation by purpose even if the same API is used
- Extract common logic into modules like `core/provider.go` to eliminate duplication
- Maintain strict top-down dependency direction (no reversed dependencies)
- Use singular directory and package names (e.g., `route`, `util`, `middleware`, not `routes`, `utils`, etc.)
- Only use plural directory names for static assets (e.g., `assets`, `images`, `docs`)

---

## 📚 Future Extension Guidelines

- If Adapters or Repositories increase, split into `domain/adapter/` or `domain/repository/`
- If presentation logic becomes complex, consider `presentation/ui/` structure
- Domain Services into `domain` as needed
- Consider introducing middleware under `server/middleware/` for cross-cutting concerns

---

This document should be maintained and updated as architecture evolves to ensure consistency across the team and codebase.

