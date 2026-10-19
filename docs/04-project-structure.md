# Project Structure

> "A well-organized architecture is the foundation of a maintainable software system."

---

# Document Information

| Field | Value |
|-------|-------|
| Project | DocsAI |
| Document | Project Structure |
| Version | 1.0 |
| Audience | Developers, Contributors, Architects |

---

# 1. Introduction

DocsAI follows a modular monorepo architecture where each major component is separated according to its responsibility.

The repository contains three primary applications:

- Frontend (User Interface)
- Backend (Business Logic)
- AI Service (Artificial Intelligence)

These applications share common contracts through Protocol Buffers while remaining independently deployable.

---

# 2. Repository Structure

```
DocsAI/
│
├── backend/
├── frontend/
├── ai-service/
├── proto/
├── deployments/
├── docker/
├── scripts/
├── docs/
├── .env.example
├── docker-compose.yml
├── Makefile
└── README.md
```

Each directory has a specific responsibility.

---

# 3. Root Directory

The root of the repository contains shared configuration and orchestration files.

## Responsibilities

- Project configuration
- Docker orchestration
- Shared Protocol Buffers
- Documentation
- Deployment assets
- Development tooling

The root directory intentionally contains no application business logic.

---

# 4. backend/

The backend is responsible for all business logic and serves as the primary API consumed by the frontend.

Technology:

- Golang
- Gin
- GORM

Responsibilities:

- REST APIs
- Authentication
- Authorization
- Validation
- Database access
- File upload
- Metadata management
- AI orchestration
- gRPC client
- Storage management

---

## backend/cmd/

Application entry points.

Example:

```
cmd/server
```

Responsibilities:

- Application startup
- Dependency initialization
- Configuration loading
- HTTP server bootstrap

---

## backend/internal/

Contains all private application code.

Everything inside this directory is intended only for the backend application.

Major modules include:

- document
- ai
- database
- storage
- middleware
- config
- router
- dto
- schema

---

## backend/internal/document/

Document management domain.

Responsibilities:

- Upload documents
- Rename documents
- Delete documents
- List documents
- Save chunks
- Save embeddings
- Document repository
- Document service
- Handlers

This module owns the lifecycle of uploaded documents.

---

## backend/internal/ai/

AI client abstraction.

Responsibilities:

- gRPC client
- AI service communication
- Protocol Buffer integration

The backend never performs AI work directly.

Instead, it delegates processing to the AI Service.

---

## backend/internal/database/

Database layer.

Responsibilities:

- PostgreSQL connection
- Migrations
- Schema definitions
- Repository support

---

## backend/internal/storage/

Object storage abstraction.

Responsibilities:

- Upload files
- Download files
- Delete files
- Generate object paths

Current implementation uses MinIO.

Future implementations may include AWS S3 or Google Cloud Storage.

---

## backend/pkg/

Reusable packages that may be shared across modules.

Unlike `internal`, packages here can be safely imported by other projects if needed.

---

## backend/proto/

Generated Go code from Protocol Buffers.

Contains:

- gRPC client interfaces
- Message definitions

These files are auto-generated and should never be manually edited.

---

# 5. frontend/

The frontend provides the user interface.

Technology:

- Next.js
- React
- TypeScript
- Tailwind CSS

Responsibilities:

- Authentication
- Dashboard
- Upload pages
- Chat interface
- Search interface
- Workspace management
- API integration

The frontend contains no business logic.

---

# 6. ai-service/

The AI Service is a dedicated Python application responsible for computationally intensive AI operations.

Technology:

- Python
- gRPC
- OCR libraries
- Embedding providers
- LLM integrations

Responsibilities:

- Document parsing
- OCR
- Text extraction
- Chunk generation
- Embedding generation
- Prompt construction
- Future LLM orchestration

---

## ai-service/providers/

Provider abstraction layer.

Current providers include:

- Ollama
- Gemini
- OpenAI (planned)

This abstraction allows switching AI providers without changing business logic.

---

## ai-service/services/

Core AI processing logic.

Examples:

- Document processor
- Chunking
- Embedding pipeline

---

## ai-service/models/

Internal data models used by the AI service.

These are separate from backend database schemas.

---

## ai-service/proto/

Generated Python Protocol Buffer files.

Responsibilities:

- gRPC server definitions
- Message types
- Callback contracts

These files are generated automatically.

---

# 7. proto/

The shared Protocol Buffer definitions.

This directory acts as the contract between:

- Backend (Go)
- AI Service (Python)

No implementation logic belongs here.

Only shared service definitions and message schemas.

---

# 8. deployments/

Deployment manifests.

Future responsibilities include:

- Kubernetes
- Helm Charts
- Production configuration
- Infrastructure manifests

---

# 9. docker/

Containerization assets.

Responsibilities:

- Dockerfiles
- Container configuration
- Development environment

---

# 10. scripts/

Automation utilities.

Examples:

- Build scripts
- Development helpers
- Maintenance tasks

---

# 11. docs/

Project documentation.

Contains:

- Product documentation
- Architecture
- Database
- APIs
- Deployment guides
- Developer onboarding

This directory serves as the project's single source of truth for documentation.

---

# 12. Development Workflow

A simplified development flow:

```
Developer

↓

Frontend

↓

Backend

↓

Database

↓

Storage

↓

AI Service

↓

LLMs
```

Each layer communicates through well-defined interfaces.

---

# 13. Dependency Direction

Dependencies flow in one direction:

```
Frontend
    ↓
Backend
    ↓
Repository
    ↓
Database

Backend
    ↓
AI Client
    ↓
AI Service

AI Service
    ↓
LLM Providers
```

This avoids circular dependencies and keeps responsibilities clear.

---

# 14. Architectural Principles

The project structure is designed around these principles:

- High cohesion within modules
- Low coupling between modules
- Clear ownership of responsibilities
- Independent deployability
- Technology specialization
- Extensibility

---

# 15. Summary

The DocsAI repository is intentionally organized to separate presentation, business logic, artificial intelligence, and infrastructure concerns.

This modular structure allows the platform to evolve over time while keeping the codebase maintainable, scalable, and approachable for new contributors.

A developer joining the project should be able to identify where functionality belongs based on responsibility rather than implementation details, making the system easier to understand and extend.