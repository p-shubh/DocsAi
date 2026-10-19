# Backend Architecture

> "The backend is the central coordinator of DocsAI. It manages business logic, orchestrates AI workflows, persists data, and exposes APIs to the frontend."

---

# Document Information

| Field | Value |
|-------|-------|
| Project | DocsAI |
| Document | Backend Architecture |
| Version | 1.0 |
| Technology | Golang |
| Framework | Gin |
| ORM | GORM |

---

# 1. Overview

The DocsAI backend is responsible for coordinating every operation that occurs inside the platform.

Unlike the AI Service, which focuses exclusively on artificial intelligence, the backend is responsible for enforcing business rules, managing data, handling user requests, and orchestrating communication between all system components.

The backend never performs expensive AI operations directly. Instead, it delegates those responsibilities to the Python AI Service through gRPC.

---

# 2. Responsibilities

The backend currently manages:

- REST API endpoints
- Request validation
- Business logic
- Authentication (planned)
- Authorization (planned)
- File uploads
- Object storage
- Database persistence
- Metadata management
- gRPC communication
- AI orchestration
- Callback handling
- Error handling

Future responsibilities include:

- Chat orchestration
- Retrieval-Augmented Generation (RAG)
- Conversation history
- Workspaces
- Collections
- Team management
- RBAC
- Analytics

---

# 3. Architectural Pattern

The backend follows a layered architecture.

```
                HTTP Request
                      │
                      ▼
                Gin Router
                      │
                      ▼
                 Handler Layer
                      │
                      ▼
                 Service Layer
                      │
                      ▼
               Repository Layer
                      │
                      ▼
                 PostgreSQL
```

For AI processing:

```
Handler
    │
    ▼
Service
    │
    ├──────── Storage
    │
    ├──────── Database
    │
    └──────── AI Client (gRPC)
                  │
                  ▼
           Python AI Service
```

Each layer has a single responsibility and depends only on the layer below it.

---

# 4. Layer Responsibilities

## Router Layer

The router is responsible for:

- Registering API endpoints
- Applying middleware
- Versioning APIs
- Connecting HTTP routes to handlers

Business logic never belongs in the router.

---

## Handler Layer

Handlers are the entry point for every REST request.

Responsibilities:

- Parse requests
- Validate input
- Call services
- Return HTTP responses
- Handle HTTP status codes

Handlers should remain lightweight.

They should not contain business rules or database queries.

---

## Service Layer

The Service layer contains business logic.

Responsibilities include:

- Coordinating repositories
- Calling storage services
- Calling AI services
- Performing validations
- Managing workflows
- Enforcing business rules

The Service layer is the heart of the backend.

---

## Repository Layer

Repositories interact directly with PostgreSQL.

Responsibilities:

- CRUD operations
- Query optimization
- Transactions
- Persistence
- Vector search (current implementation)

Repositories should never contain business logic.

---

## Storage Layer

The storage abstraction manages binary objects.

Current implementation:

- MinIO

Future implementations may include:

- Amazon S3
- Google Cloud Storage
- Azure Blob Storage

The backend interacts only with the storage interface.

---

## AI Client

The AI Client provides a Go abstraction over gRPC communication.

Responsibilities:

- Build protobuf requests
- Call Python services
- Handle responses
- Abstract transport details

Business logic should not know about gRPC internals.

---

# 5. Dependency Injection

The backend follows dependency injection principles.

A service receives only the dependencies it requires.

Example:

```
Document Service

↓

Repository

↓

Storage

↓

AI Client
```

Benefits:

- Easier testing
- Loose coupling
- Better maintainability
- Easier mocking
- Simpler future extensions

---

# 6. Current Document Upload Flow

The current upload pipeline works as follows:

```
Client

↓

Upload API

↓

Validate Request

↓

Upload File to MinIO

↓

Create Document Record

↓

Call AI Service

↓

Return Accepted Response
```

The AI processing continues asynchronously.

---

# 7. AI Processing Callback

After processing:

```
Python AI Service

↓

Extract Text

↓

Chunk Document

↓

Generate Embeddings

↓

Callback API

↓

Save Chunks

↓

Save Embeddings

↓

Update Document Status
```

This architecture allows document uploads to complete quickly while heavy AI work continues in the background.

---

# 8. Current Backend Modules

The backend currently consists of several logical modules.

## Document

Responsibilities:

- Upload
- Rename
- Delete
- List
- Retrieve
- Chunk persistence
- Embedding persistence

---

## AI

Responsibilities:

- gRPC client
- Protocol Buffers
- AI communication

---

## Database

Responsibilities:

- PostgreSQL connection
- GORM configuration
- Schema management

---

## Storage

Responsibilities:

- File upload
- File download
- File deletion

---

## Common

Shared utilities including:

- Enums
- Helpers
- Constants
- Utilities

---

# 9. Current Data Flow

```
User

↓

REST API

↓

Document Handler

↓

Document Service

↓

Storage

↓

Repository

↓

AI Client

↓

Python AI Service
```

Notice that every request flows downward.

Lower layers never call upper layers.

---

# 10. Error Handling Strategy

Errors propagate upward.

```
Repository

↓

Service

↓

Handler

↓

HTTP Response
```

Each layer adds context while preserving the original error whenever possible.

This approach simplifies debugging.

---

# 11. Current Strengths

The backend already demonstrates several good architectural qualities:

✓ Clear separation of concerns

✓ Layered architecture

✓ AI abstraction

✓ Storage abstraction

✓ Repository pattern

✓ Dependency injection

✓ Modular organization

✓ Asynchronous AI processing

✓ Protocol Buffer contracts

These provide a solid foundation for future growth.

---

# 12. Current Engineering Focus

The backend ingestion pipeline is largely complete.

The current engineering focus is shifting toward Retrieval-Augmented Generation (RAG).

The next major capabilities include:

- Query embedding
- Semantic vector search
- Prompt construction
- LLM orchestration
- Chat endpoints
- Source citations
- Conversation history

These features build on the ingestion pipeline that is already in place.

---

# 13. Planned Architectural Evolution

As DocsAI grows, the backend will likely evolve into additional domains.

A possible future organization:

```
Document Module
    Upload
    Delete
    Rename
    List

Search Module
    Query Embedding
    Vector Search
    Retrieval

Chat Module
    Prompt Construction
    LLM Calls
    Conversation Memory

Workspace Module
    Teams
    Collections
    Permissions

Admin Module
    Users
    Roles
    Monitoring
```

This keeps each module focused on a single responsibility.

---

# 14. Summary

The DocsAI backend is designed as an orchestration layer rather than an AI engine.

Its primary responsibility is to coordinate storage, databases, AI services, and business workflows while exposing a clean and stable API to clients.

By keeping AI processing isolated behind well-defined interfaces, the backend remains maintainable, testable, and flexible enough to support future AI capabilities without major architectural changes.