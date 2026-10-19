# AI Service Architecture

> "The AI Service is the intelligence engine of DocsAI. It transforms raw documents into structured knowledge that can later be searched, retrieved, and understood."

---

# Document Information

| Field | Value |
|-------|-------|
| Project | DocsAI |
| Document | AI Service Architecture |
| Version | 1.0 |
| Technology | Python |
| Communication | gRPC |
| Status | Active Development |

---

# 1. Overview

The AI Service is an independent Python application responsible for every Artificial Intelligence operation performed by DocsAI.

Instead of embedding AI logic inside the Go backend, DocsAI separates AI into its own service.

This decision provides several advantages:

- Better scalability
- Independent deployment
- AI framework flexibility
- Easier GPU migration
- Language specialization
- Cleaner backend architecture

The Go backend acts as the orchestrator.

The Python service acts as the intelligence engine.

---

# 2. Responsibilities

The AI Service is responsible for:

- Receiving document processing requests
- Downloading uploaded files
- Parsing documents
- OCR (future)
- Extracting text
- Cleaning extracted content
- Chunking documents
- Generating embeddings
- Returning structured processing results
- Generating query embeddings
- Future prompt construction
- Future LLM orchestration

The AI Service never stores business data directly.

It sends processed results back to the backend through callbacks.

---

# 3. High-Level Architecture

```
                     Backend (Go)
                           │
                     ProcessDocument
                           │
                        gRPC Call
                           │
                           ▼
               +-------------------------+
               |     AI Service          |
               +-----------+-------------+
                           │
         +-----------------+----------------+
         |                 |                |
         ▼                 ▼                ▼
  File Download      Text Extraction    Metadata
         │
         ▼
     Chunking Engine
         │
         ▼
  Embedding Generator
         │
         ▼
 Callback Client
         │
         ▼
     Go Backend
```

---

# 4. Why Python?

Python is selected because it has the strongest AI ecosystem.

Examples include:

- PyTorch
- TensorFlow
- Hugging Face
- LangChain
- LlamaIndex
- OCR Libraries
- PDF Processing Libraries

Rather than forcing these libraries into Go, DocsAI uses each language for what it does best.

---

# 5. Internal Modules

The AI Service is organized into logical modules.

---

## Config

Responsibilities:

- Environment variables
- Model configuration
- Provider configuration
- Runtime settings

Every configurable value should be centralized here.

---

## Models

Contains internal Python models.

Examples:

- DocumentInfo
- Chunk
- Embedding
- ProcessingResult

These models are internal to the AI service.

They are different from backend database schemas.

---

## Providers

Responsible for communicating with AI providers.

Current provider architecture:

```
Providers

↓

Base Provider

↓

Ollama

↓

Gemini

↓

OpenAI (Future)
```

Each provider implements a common interface.

---

## Services

Contains business logic.

Examples:

- DocumentProcessor
- ChunkGenerator
- EmbeddingGenerator

This layer coordinates AI workflows.

---

## Clients

Responsible for external communication.

Examples:

- Callback Client
- Storage Client (Future)

Clients communicate with external systems.

---

## Proto

Contains generated gRPC files.

Generated from:

document.proto

These files should never be edited manually.

---

# 6. Current Processing Flow

The AI pipeline currently follows these steps:

```
Receive Request

↓

Download File

↓

Extract Text

↓

Chunk Document

↓

Generate Embeddings

↓

Create Processing Result

↓

Callback Backend
```

Each step has a single responsibility.

---

# 7. Provider Abstraction

DocsAI does not directly depend on a specific AI provider.

Instead, it uses a provider abstraction.

```
Embedding Request

↓

Factory

↓

Provider Interface

↓

Selected Provider

↓

Embedding Vector
```

Current providers:

- Ollama
- Gemini

Future providers:

- OpenAI
- Anthropic
- Cohere
- Azure OpenAI
- Local Models

Changing providers should require configuration changes rather than application rewrites.

---

# 8. Current Embedding Pipeline

```
Chunk

↓

Embedding Provider

↓

Vector

↓

Processing Result

↓

Backend Callback

↓

Database
```

Each document chunk produces one embedding vector.

These vectors become searchable through PostgreSQL pgvector.

---

# 9. Query Embedding

The current architecture is evolving to support semantic retrieval.

Flow:

```
Question

↓

EmbedQuery

↓

Embedding Vector

↓

Backend

↓

Vector Search

↓

Relevant Chunks
```

Unlike document embeddings, query embeddings are generated in real time whenever a user asks a question.

---

# 10. Callback Architecture

The AI Service is intentionally stateless.

It does not write directly to PostgreSQL.

Instead:

```
Processing Complete

↓

Callback API

↓

Go Backend

↓

Repository

↓

Database
```

Benefits:

- Business logic remains centralized.
- AI Service has no database dependency.
- Easier testing.
- Better security.
- Independent deployment.

---

# 11. Error Handling

Errors are categorized into:

Input Errors

- Invalid documents
- Unsupported formats

Processing Errors

- OCR failures
- Parsing failures
- Embedding failures

Communication Errors

- gRPC failures
- Callback failures

Provider Errors

- Ollama unavailable
- Gemini unavailable
- Timeout
- Rate limiting

Every layer should add context before propagating errors.

---

# 12. Scalability

The AI Service is designed for horizontal scaling.

Future deployment:

```
Load Balancer

↓

AI Service #1

AI Service #2

AI Service #3

AI Service #4
```

Since the service is stateless, scaling is straightforward.

---

# 13. Future AI Capabilities

The current AI pipeline focuses on document ingestion.

Future capabilities include:

- OCR
- Image Understanding
- Table Extraction
- Diagram Analysis
- Multi-modal AI
- Automatic Summaries
- Keyword Extraction
- Named Entity Recognition
- Citation Generation
- Prompt Optimization
- AI Agents

---

# 14. Engineering Principles

The AI Service follows these principles:

✓ Stateless

✓ Provider Independent

✓ Modular

✓ Scalable

✓ Testable

✓ Replaceable

✓ AI Focused

Business rules remain outside the AI service.

---

# 15. Current Status

### Completed

✓ gRPC Server

✓ Provider Factory

✓ Ollama Integration

✓ Embedding Generation

✓ Chunk Processing

✓ Callback Flow

✓ Document Processing Pipeline

✓ Query Embedding RPC

---

### In Progress

- Retrieval integration
- Prompt construction
- Gemini answer generation
- Citation support

---

### Planned

- Hybrid search
- OCR engine
- Image understanding
- Multi-modal AI
- Streaming responses
- Conversation memory
- AI agents

---

# 16. Summary

The AI Service is the computational intelligence layer of DocsAI.

It converts raw documents into structured knowledge by extracting text, generating embeddings, and preparing data for semantic retrieval.

By isolating AI functionality into its own service, DocsAI remains modular, scalable, and flexible enough to adopt future AI technologies without requiring major backend redesign.