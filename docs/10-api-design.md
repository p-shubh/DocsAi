# API Design

> "The REST API is the public interface of DocsAI. Every interaction between the frontend and backend occurs through carefully designed HTTP endpoints that expose business capabilities while hiding internal implementation details."

---

# Document Information

| Field | Value |
|-------|-------|
| Project | DocsAI |
| Document | API Design |
| Version | 1.0 |
| Protocol | REST |
| Format | JSON |
| Authentication | JWT (Planned) |

---

# 1. Introduction

The DocsAI backend exposes a RESTful API consumed by the frontend and future third-party clients.

The API is designed around resource-oriented principles, consistent response structures, and clear separation between business logic and transport.

Every endpoint follows the same request lifecycle:

```
Client

↓

HTTP Request

↓

Router

↓

Handler

↓

Service

↓

Repository

↓

Database / AI Service

↓

HTTP Response
```

---

# 2. API Principles

DocsAI APIs are designed around the following principles.

### RESTful

Resources are represented using nouns instead of verbs.

Example:

```
/documents
```

instead of

```
/uploadDocument
```

---

### Stateless

Each request contains everything needed to process it.

The server does not depend on previous requests.

---

### Consistent

Every API follows the same response structure.

Errors are returned in a predictable format.

---

### Versioned

All APIs are versioned.

Example:

```
/api/v1/...
```

Future versions can coexist without breaking existing clients.

---

# 3. Authentication

## Current

Authentication is under development.

---

## Planned

JWT Authentication

```
Login

↓

JWT Token

↓

Authorization Header

↓

Protected APIs
```

Header

```
Authorization: Bearer <token>
```

---

# 4. API Categories

The backend currently contains or plans the following API groups.

```
Authentication

Documents

Callbacks

Health

Knowledge (Future)

Chat (Future)

Workspace (Future)

Admin (Future)
```

---

# 5. Health API

Purpose

Verify backend availability.

### Endpoint

```
GET /api/v1/health
```

Response

```json
{
    "status": "ok"
}
```

Status Code

```
200 OK
```

---

# 6. Upload Document API

Purpose

Upload a document for AI processing.

### Endpoint

```
POST /api/v1/documents
```

Content Type

```
multipart/form-data
```

Request

```
File

↓

Validation

↓

Storage

↓

Database

↓

AI Processing
```

Success Response

```json
{
    "id": "...",
    "status": "processing"
}
```

Status

```
202 Accepted
```

Reason:

The document continues processing asynchronously.

---

# 7. Get Document

Purpose

Retrieve document information.

Endpoint

```
GET /api/v1/documents/{id}
```

Response

```json
{
    "id": "...",
    "filename": "...",
    "status": "ready"
}
```

Status

```
200 OK
```

---

# 8. List Documents

Purpose

Return uploaded documents.

Endpoint

```
GET /api/v1/documents
```

Future Query Parameters

```
page

limit

status

search

sort
```

Response

```json
{
    "documents": [],
    "total": 0
}
```

---

# 9. Rename Document

Purpose

Update document name.

Endpoint

```
PUT /api/v1/documents/{id}
```

Request

```json
{
    "name":"Project Documentation"
}
```

Response

```
200 OK
```

---

# 10. Delete Document

Purpose

Delete a document.

Endpoint

```
DELETE /api/v1/documents/{id}
```

Flow

```
Delete Metadata

↓

Delete Object

↓

Soft Delete
```

Status

```
204 No Content
```

---

# 11. Callback API

Purpose

Receive AI processing results.

Caller

Python AI Service

Endpoint

```
POST /api/v1/callback/document
```

Payload

```
Document ID

↓

Chunks

↓

Embeddings

↓

Metadata
```

Responsibilities

- Save Chunks
- Save Embeddings
- Update Status

This endpoint is intended for internal service communication.

---

# 12. Current Processing Flow

```
POST Document

↓

Upload

↓

AI Service

↓

Callback

↓

Ready
```

The upload request and callback request together complete the ingestion pipeline.

---

# 13. Future Knowledge APIs

Future retrieval endpoints.

### Search

```
POST /api/v1/knowledge/search
```

Purpose

Retrieve relevant document chunks.

---

### Embed Query

Internal Flow

```
Question

↓

EmbedQuery

↓

Vector Search

↓

Results
```

The embedding generation is performed by the AI Service through gRPC.

---

# 14. Future Chat APIs

```
POST /api/v1/chat
```

Request

```json
{
    "question":"Explain JWT authentication."
}
```

Response

```json
{
    "answer":"...",
    "sources":[]
}
```

The chat endpoint will orchestrate retrieval, prompt construction, and LLM response generation.

---

# 15. Standard Response Format

Success

```json
{
    "success": true,
    "data": {}
}
```

Failure

```json
{
    "success": false,
    "message":"Validation failed"
}
```

A consistent response structure simplifies frontend integration.

---

# 16. HTTP Status Codes

DocsAI follows standard HTTP semantics.

| Status | Meaning |
|----------|----------|
|200|Request Successful|
|201|Resource Created|
|202|Accepted for Processing|
|204|Deleted Successfully|
|400|Bad Request|
|401|Unauthorized|
|403|Forbidden|
|404|Not Found|
|409|Conflict|
|422|Validation Error|
|500|Internal Server Error|

---

# 17. Validation

Every endpoint validates:

- Required fields
- UUID format
- File size
- MIME type
- JSON payload
- Business constraints

Validation occurs before business logic executes.

---

# 18. Error Handling

Errors propagate through the backend.

```
Repository

↓

Service

↓

Handler

↓

HTTP Response
```

Internal errors are logged.

Clients receive user-friendly error messages.

---

# 19. API Versioning Strategy

Current version

```
v1
```

Future

```
v2
```

Breaking changes will introduce new versions without disrupting existing clients.

---

# 20. Security Considerations

Future security includes:

- JWT Authentication
- Refresh Tokens
- RBAC
- Rate Limiting
- Request Logging
- Audit Logs
- API Keys
- CSRF Protection (where applicable)

---

# 21. Future API Groups

As DocsAI evolves, additional endpoints will be introduced.

```
Authentication

Users

Organizations

Teams

Collections

Folders

Knowledge

Chat

Memory

Search

Analytics

Administration
```

Each group will remain independent and follow the same architectural conventions.

---

# 22. API Lifecycle

A typical request flows through the backend as follows.

```
Client

↓

REST Endpoint

↓

Router

↓

Handler

↓

Service

↓

Repository

↓

Database

↓

Response
```

AI-related requests extend this flow by invoking the Python AI Service through gRPC.

---

# 23. Design Philosophy

The DocsAI API is designed to be:

✓ Consistent

✓ Predictable

✓ Versioned

✓ Secure

✓ Extensible

✓ Frontend Friendly

✓ AI Ready

The API serves as the stable contract between the frontend, backend, and future external integrations.

---

# 24. Summary

The REST API is the primary interface through which users interact with DocsAI.

Its role extends beyond CRUD operations—it orchestrates document ingestion, AI processing, retrieval, and future conversational experiences while maintaining a clean, consistent, and scalable interface for all clients.

As the platform evolves, the API will continue to expand while preserving backward compatibility and maintaining clear architectural boundaries.