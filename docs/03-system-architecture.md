# System Architecture

> "A scalable, modular, AI-first architecture designed for intelligent document understanding and conversational knowledge retrieval."

---

# Document Information

| Field | Value |

|-------|-------|

| Project | DocsAI |

| Document | System Architecture |

| Version | 1.0 |

| Status | Active Development |

| Audience | Developers, Architects, DevOps, AI Engineers |

---

# 1. Introduction

DocsAI follows a modular service-oriented architecture where each component is responsible for a single domain.

Rather than placing all logic inside one application, responsibilities are divided between specialized services that communicate using well-defined interfaces.

This architecture provides several advantages:

- Separation of concerns

- Independent scalability

- Easier maintenance

- Technology flexibility

- Better fault isolation

- Future extensibility

Each service is responsible for solving one problem well.

---

# 2. High-Level Architecture

```

                    +-----------------------------------+

                    |          User / Browser           |

                    +----------------+------------------+

                                     |

                                     |

                              HTTPS / REST

                                     |

                                     ▼

                    +-----------------------------------+

                    |         Next.js Frontend          |

                    +----------------+------------------+

                                     |

                                     |

                              REST APIs

                                     |

                                     ▼

              +-----------------------------------------------+

              |               Go Backend                       |

              |            Gin Framework                       |

              +---------+--------------+-----------------------+

                        |              |

                        |              |

                        |              |

                        ▼              ▼

                 PostgreSQL         MinIO

                  pgvector      Object Storage

                        |

                        |

                        |

                 gRPC Communication

                        |

                        ▼

          +--------------------------------------+

          |         Python AI Service            |

          +----------------+---------------------+

                           |

        +------------------+----------------------+

        |                  |                      |

        ▼                  ▼                      ▼

    OCR Engine        Chunking Engine      Embedding Engine

                                                  |

                                    +-------------+-------------+

                                    |                           |

                                    ▼                           ▼

                               Ollama                     Gemini/OpenAI

```

---

# 3. Architectural Philosophy

DocsAI follows four primary architectural principles.

## Separation of Responsibilities

Each service owns a single responsibility.

The frontend manages user interaction.

The backend manages business logic.

The AI service manages artificial intelligence.

The database stores structured information.

Object storage stores binary files.

---

## Service Independence

Each service can evolve independently.

For example:

The AI service may later move to GPU infrastructure without affecting the backend.

The frontend can evolve independently of AI model changes.

The storage provider can change without rewriting AI logic.

---

## Clear Communication Contracts

Communication occurs using:

REST

for frontend ↔ backend

and

gRPC

for backend ↔ AI service.

This keeps APIs strongly typed and versionable.

---

## AI Provider Independence

DocsAI is intentionally designed so that AI providers can change without affecting the rest of the platform.

Today:

- Ollama

- Gemini

Future:

- OpenAI

- Anthropic

- Mistral

- Azure OpenAI

- Local LLMs

---

# 4. Component Overview

## Frontend

Technology

- Next.js

- React

- TypeScript

- Tailwind CSS

Responsibilities

- User authentication

- Dashboard

- Upload interface

- Chat interface

- Search interface

- Workspace management

- API communication

The frontend contains no document processing logic.

---

## Backend

Technology

- Golang

- Gin

- GORM

Responsibilities

- Authentication

- Authorization

- Business rules

- API layer

- Validation

- File upload

- Metadata management

- Database persistence

- AI orchestration

The backend acts as the central coordinator of the system.

---

## PostgreSQL

Stores

- Users

- Documents

- Chunks

- Embeddings

- Metadata

- Workspaces

- Future chat history

PostgreSQL is selected because it provides:

- ACID transactions

- Mature ecosystem

- JSON support

- pgvector integration

---

## pgvector

The pgvector extension enables semantic similarity search directly inside PostgreSQL.

Instead of keyword matching, vectors are compared using mathematical distance metrics.

This allows DocsAI to retrieve information based on meaning rather than exact wording.

---

## MinIO

MinIO stores the original uploaded files.

Examples:

PDFs

DOCX

Images

Presentations

Large binary objects are intentionally separated from relational data.

---

## Redis

Redis is used for high-speed temporary storage.

Current and future use cases include:

- Caching

- Session storage

- Rate limiting

- Background jobs

- AI response caching

---

## AI Service

The Python AI Service performs computationally intensive operations.

Responsibilities include:

- OCR

- Text extraction

- Parsing

- Chunking

- Embedding generation

- Prompt preparation

- Future LLM inference

Separating AI from the backend keeps the Go service lightweight.

---

# 5. End-to-End Document Flow

```

User

↓

Upload Document

↓

Backend

↓

Save File → MinIO

↓

Save Metadata → PostgreSQL

↓

Send gRPC Request

↓

Python AI Service

↓

Download File

↓

Extract Text

↓

Chunk Text

↓

Generate Embeddings

↓

Callback

↓

Save Chunks

↓

Save Embeddings

↓

Document Ready

```

---

# 6. End-to-End Question Answering Flow (Target Architecture)

```

User

↓

Ask Question

↓

Backend

↓

Generate Query Embedding

↓

Vector Search

↓

Retrieve Relevant Chunks

↓

Construct Prompt

↓

Gemini

↓

Generate Response

↓

Return Answer

```

This architecture is currently under active development.

---

# 7. Why Two Services?

Some might ask:

"Why not implement everything in Go?"

The answer is specialization.

Go excels at:

- APIs

- Concurrency

- Performance

- Networking

Python excels at:

- AI libraries

- Machine Learning

- NLP

- OCR

- Model integration

Using each language where it performs best results in a cleaner, more maintainable system.

---

# 8. Why gRPC?

The backend and AI service communicate using gRPC because it offers:

- Strong typing

- High performance

- Efficient serialization

- Code generation

- Language interoperability

This allows Go and Python to communicate through a shared contract defined in Protocol Buffers.

---

# 9. Why Object Storage?

Storing large files directly inside PostgreSQL would reduce database performance.

Instead:

Metadata lives in PostgreSQL.

Files live in MinIO.

This separation improves scalability and backup strategies.

---

# 10. Why Vector Search?

Traditional keyword search answers:

"Does this word exist?"

Vector search answers:

"What does this text mean?"

This enables semantic understanding.

Example:

Question:

"How do users sign in?"

Relevant documentation may contain:

"Authentication"

"Login"

"User verification"

even if the phrase "sign in" never appears.

---

# 11. Future Architecture

The current architecture intentionally leaves room for future expansion.

Examples include:

Knowledge Graph

Conversation Memory

Hybrid Search

AI Agents

Workflow Automation

Enterprise Authentication

Plugin System

External Connectors

Analytics

Monitoring

Each future capability can be added without redesigning the core platform.

---

# 12. Architectural Strengths

Current strengths include:

✓ Modular design

✓ Independent AI service

✓ Provider abstraction

✓ Scalable storage

✓ Strong API contracts

✓ AI-ready architecture

✓ Service isolation

✓ Database extensibility

✓ Future-proof foundation

---

# 13. Current Limitations

The architecture is still evolving.

Current engineering focus includes:

- Retrieval layer

- Semantic search

- Prompt orchestration

- Chat APIs

- Source citations

- Conversation history

- Access control

These components build upon the existing ingestion pipeline.

---

# 14. Architecture Summary

DocsAI is designed around a simple principle:

Each service should focus on one responsibility while communicating through well-defined interfaces.

This separation enables the platform to scale from a personal document assistant into an enterprise-grade AI Knowledge Platform without fundamental architectural changes.

The architecture intentionally prioritizes modularity, maintainability, and future extensibility over short-term simplicity.