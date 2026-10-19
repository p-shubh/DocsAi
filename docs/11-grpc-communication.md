# gRPC Communication

> "gRPC provides the high-performance communication layer between the Go Backend and the Python AI Service, enabling strongly typed, language-independent, and efficient service-to-service communication."

---

# Document Information

| Field | Value |
|-------|-------|
| Project | DocsAI |
| Document | gRPC Communication |
| Version | 1.0 |
| Protocol | gRPC |
| Serialization | Protocol Buffers |
| Languages | Go & Python |

---

# 1. Introduction

DocsAI is composed of two primary backend services:

• Go Backend
• Python AI Service

Rather than communicating over REST, these services communicate using gRPC.

This provides:

- Strongly typed contracts
- Binary serialization
- Better performance
- Automatic code generation
- Cross-language compatibility

---

# 2. Why gRPC?

Traditional REST APIs are ideal for client-server communication.

However, internal service-to-service communication benefits from:

✓ Lower latency

✓ Smaller payloads

✓ Type safety

✓ Contract-first development

✓ Language independence

---

# 3. Communication Architecture

```
                Go Backend
                     │
              DocumentAIClient
                     │
                gRPC Channel
                     │
        document.proto Contract
                     │
             Python AI Server
```

Both services are generated from the same Protocol Buffer definition.

---

# 4. Shared Protocol

```
proto/

└── document.proto
```

This file is the single source of truth.

Nothing should be implemented without first updating the protobuf contract.

---

# 5. Current Services

## DocumentAI

Responsibilities:

- Process uploaded documents
- Generate query embeddings

Current RPCs

```
ProcessDocument()

EmbedQuery()
```

---

## DocumentCallback

Responsibilities:

Receive processed results from the AI Service.

Current RPC

```
SaveProcessingResult()
```

---

# 6. ProcessDocument()

Purpose

Start asynchronous AI processing.

Input

```
Document ID

Bucket

Object Key

Mime Type
```

Output

```
Accepted
```

The backend only checks whether the AI Service accepted the request.

---

# 7. EmbedQuery()

Purpose

Generate an embedding for a user's question.

Flow

```
Question

↓

AI Service

↓

Embedding Vector
```

This vector is used for semantic search.

---

# 8. SaveProcessingResult()

Purpose

Return processed document information to the backend.

Contains:

- Chunks
- Embeddings
- Metadata

The backend remains the only component that writes to PostgreSQL.

---

# 9. Generated Code

Go

```
backend/proto/documentpb/
```

Python

```
ai-service/proto/
```

Generated files must never be manually edited.

---

# 10. Protobuf Workflow

Whenever document.proto changes:

```
Modify document.proto

↓

Generate Go Files

↓

Generate Python Files

↓

Compile Backend

↓

Compile AI Service

↓

Verify Both
```

---

# 11. Current RPC Flow

```
Upload

↓

ProcessDocument

↓

Python Processing

↓

Callback

↓

Database

↓

READY
```

Question flow

```
Question

↓

EmbedQuery

↓

Vector Search

↓

Gemini

↓

Answer
```

---

# 12. Future RPCs

Potential additions:

```
GenerateAnswer()

Summarize()

Translate()

GenerateQuiz()

ExtractEntities()

ExtractTables()

GenerateFlashcards()

DetectLanguage()

ConversationMemory()
```

---

# 13. Versioning Strategy

Never break existing RPCs.

Instead:

- Add fields.
- Add new RPCs.
- Deprecate old ones.
- Generate new clients.

---

# 14. Error Handling

All RPC errors should return:

- Invalid Argument
- Not Found
- Internal
- Deadline Exceeded
- Unavailable

Backend converts these into HTTP responses.

---

# 15. Summary

gRPC forms the internal backbone of DocsAI by providing fast, strongly typed, and maintainable communication between the Go backend and the Python AI service. It enables both services to evolve independently while sharing a stable contract through Protocol Buffers.