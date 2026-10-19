# Document Processing Pipeline

> "The Document Processing Pipeline is the heart of DocsAI. It transforms an uploaded document into structured, searchable knowledge through a series of coordinated processing stages."

---

# Document Information

| Field | Value |
|-------|-------|
| Project | DocsAI |
| Document | Document Processing Pipeline |
| Version | 1.0 |
| Status | Active Development |
| Pipeline | Document Ingestion |

---

# 1. Overview

The Document Processing Pipeline is responsible for converting raw user documents into AI-searchable knowledge.

The pipeline begins when a user uploads a document and ends when that document becomes available for semantic retrieval and conversational AI.

The design intentionally separates:

- Upload
- Storage
- Processing
- Embedding Generation
- Persistence

This allows heavy AI processing to occur asynchronously without blocking the user.

---

# 2. End-to-End Pipeline

```
                    User
                      │
                      ▼
             Upload Document
                      │
                      ▼
              Go Backend API
                      │
                      ▼
              Request Validation
                      │
          ┌───────────┴───────────┐
          │                       │
          ▼                       ▼
 Store File in MinIO      Save Metadata
                               │
                               ▼
                     Document Status = PROCESSING
                               │
                               ▼
                    gRPC ProcessDocument()
                               │
                               ▼
                     Python AI Service
                               │
                               ▼
                    Download Document
                               │
                               ▼
                      Extract Text
                               │
                               ▼
                      Chunk Document
                               │
                               ▼
                  Generate Embeddings
                               │
                               ▼
             SaveProcessingResult Callback
                               │
                               ▼
                      Go Backend
                ┌────────┴────────┐
                ▼                 ▼
         Save Chunks      Save Embeddings
                │                 │
                └────────┬────────┘
                         ▼
              Document Status = READY
                         ▼
             Available for AI Search
```

---

# 3. Step 1 – User Upload

The process begins when a user uploads a document from the frontend.

Supported document types may include:

- PDF
- DOCX
- TXT
- Markdown
- Images (future)
- Excel (future)
- PowerPoint (future)

The frontend submits the document through a REST API.

---

# 4. Step 2 – Request Validation

The backend validates:

- File existence
- File size
- Supported MIME type
- Request metadata

If validation fails, processing stops immediately and the user receives an appropriate error response.

---

# 5. Step 3 – Object Storage

After validation, the original document is uploaded to MinIO.

MinIO stores:

- Original PDF
- Original DOCX
- Original binary file

The backend stores only:

- Object key
- Bucket name
- Metadata

This keeps PostgreSQL lightweight.

---

# 6. Step 4 – Database Entry

A new document record is created.

Typical information includes:

- Document ID
- Filename
- Object Key
- MIME Type
- Status
- Created Timestamp

Initial status:

```
PROCESSING
```

The backend immediately returns control to the user while processing continues asynchronously.

---

# 7. Step 5 – gRPC Processing Request

The backend sends a gRPC request to the AI Service.

Current request includes:

- Document ID
- Bucket Name
- Object Key
- MIME Type

The backend does not send the entire file.

Instead, the AI Service downloads it from object storage.

---

# 8. Step 6 – AI Service Receives Request

The Python AI Service receives:

```
ProcessDocumentRequest
```

The request is converted into an internal model.

Example:

```
DocumentInfo

↓

Processor
```

The AI pipeline now begins.

---

# 9. Step 7 – Download Document

The AI Service downloads the file from MinIO.

Advantages:

- Backend remains lightweight.
- No large binary transfer over gRPC.
- Easier retry logic.
- Better scalability.

---

# 10. Step 8 – Text Extraction

The AI Service extracts readable text.

Current capabilities:

- PDF text extraction

Future capabilities:

- OCR
- Image extraction
- Tables
- Charts
- Handwritten text
- Layout understanding

The objective is to obtain clean textual content.

---

# 11. Step 9 – Text Cleaning

Before chunking, extracted text is normalized.

Possible operations include:

- Remove extra whitespace
- Normalize line breaks
- Remove invisible characters
- Clean encoding issues

Clean text produces better embeddings.

---

# 12. Step 10 – Chunk Generation

The cleaned document is divided into smaller sections.

```
Document

↓

Pages

↓

Paragraphs

↓

Chunks
```

Each chunk contains:

- Text
- Page number
- Section
- Metadata

Chunking improves retrieval quality while respecting LLM context windows.

---

# 13. Step 11 – Embedding Generation

Each chunk is passed to the embedding provider.

Current provider:

- Ollama

Future providers:

- Gemini
- OpenAI
- BGE
- E5
- Nomic
- Mixedbread

Each chunk produces one embedding vector.

```
Chunk

↓

Embedding Model

↓

1024-Dimensional Vector
```

---

# 14. Step 12 – Processing Result

The AI Service combines:

- Metadata
- Chunks
- Embeddings

into one processing result.

This object is transmitted back using the callback API.

The AI Service never writes directly to PostgreSQL.

---

# 15. Step 13 – Callback

Instead of database access:

```
AI Service

↓

Callback

↓

Go Backend
```

Current callback contains:

- Document ID
- Chunks
- Embeddings
- Metadata

The backend remains the single owner of business data.

---

# 16. Step 14 – Save Chunks

The backend repository persists every generated chunk.

Relationship:

```
Document

↓

Many Chunks
```

Chunks become searchable context.

---

# 17. Step 15 – Save Embeddings

Every chunk embedding is stored inside PostgreSQL using pgvector.

Relationship:

```
Chunk

↓

Embedding
```

Embeddings provide semantic meaning.

---

# 18. Step 16 – Update Status

After successful persistence:

```
PROCESSING

↓

READY
```

The document is now available for semantic retrieval.

---

# 19. Current Pipeline Characteristics

The current implementation provides:

✓ Asynchronous processing

✓ Independent AI service

✓ Object storage separation

✓ Vector generation

✓ Chunk persistence

✓ Callback architecture

✓ Status tracking

✓ Modular provider abstraction

---

# 20. Failure Handling

Failures may occur during:

Upload

↓

Storage

↓

Text Extraction

↓

Chunking

↓

Embedding

↓

Callback

↓

Database

Each stage should return meaningful errors.

Future improvements include:

- Retry queues
- Dead Letter Queue (DLQ)
- Processing audit logs
- Automatic retries
- Failure notifications

---

# 21. Future Pipeline

The pipeline will continue evolving.

Future stages include:

```
Upload

↓

Virus Scan

↓

OCR

↓

Layout Detection

↓

Table Extraction

↓

Image Captioning

↓

Entity Recognition

↓

Knowledge Graph

↓

Chunk Ranking

↓

Embedding

↓

Storage
```

This enables richer document understanding.

---

# 22. Engineering Principles

The pipeline follows several important principles:

### Asynchronous

Users are not blocked by AI processing.

---

### Fault Isolation

Failures inside AI processing do not crash the backend.

---

### Stateless AI

The AI Service maintains no business state.

---

### Single Source of Truth

Only the backend writes to PostgreSQL.

---

### Provider Independence

Embedding providers can change without redesigning the pipeline.

---

# 23. Pipeline Summary

The Document Processing Pipeline converts uploaded files into structured knowledge through a carefully orchestrated sequence of storage, parsing, chunking, embedding generation, and persistence.

By separating responsibilities between the backend and AI Service, DocsAI achieves scalability, maintainability, and flexibility while preparing every document for Retrieval-Augmented Generation (RAG).

This ingestion pipeline forms the foundation upon which all future AI capabilities—including semantic search, conversational AI, knowledge graphs, and intelligent agents—will be built.