# Retrieval-Augmented Generation (RAG) Architecture

> "The Retrieval-Augmented Generation (RAG) Engine is the intelligence layer of DocsAI. It enables Large Language Models to answer questions using verified knowledge extracted from user documents rather than relying solely on their pre-trained knowledge."

---

# Document Information

| Field | Value |
|-------|-------|
| Project | DocsAI |
| Document | Retrieval-Augmented Generation (RAG) Architecture |
| Version | 1.0 |
| Status | Under Development |
| AI Models | Gemini / Future LLMs |
| Embedding Provider | Ollama |

---

# 1. Introduction

Traditional Large Language Models generate answers from information learned during training.

While powerful, they suffer from several limitations:

- Limited knowledge of user documents
- Hallucinations
- No awareness of newly uploaded files
- Cannot answer organization-specific questions

DocsAI solves these limitations using Retrieval-Augmented Generation (RAG).

Instead of asking the LLM to answer from memory, DocsAI first retrieves the most relevant information from the user's documents and then uses that information to generate grounded responses.

---

# 2. What is RAG?

RAG combines two independent systems:

1. Information Retrieval
2. Large Language Models

Instead of:

```
Question
        ↓
LLM
        ↓
Answer
```

DocsAI performs:

```
Question

↓

Embedding

↓

Semantic Search

↓

Relevant Chunks

↓

Prompt Builder

↓

LLM

↓

Grounded Answer
```

The retrieved document context becomes the foundation for every AI response.

---

# 3. Why RAG?

RAG enables DocsAI to:

- Reduce hallucinations
- Answer from user documents
- Support private knowledge bases
- Improve answer accuracy
- Provide citations
- Support continuously changing knowledge

Without RAG, uploaded documents would never influence AI responses.

---

# 4. RAG Architecture

```
                        User

                         │

                         ▼

                 Ask Question

                         │

                         ▼

                  Backend API

                         │

                         ▼

                EmbedQuery (gRPC)

                         │

                         ▼

                 AI Service

                         │

                         ▼

              Generate Embedding

                         │

                         ▼

                Backend Repository

                         │

                         ▼

          PostgreSQL + pgvector Search

                         │

                         ▼

             Top Matching Chunks

                         │

                         ▼

               Context Builder

                         │

                         ▼

               Prompt Builder

                         │

                         ▼

               Gemini / LLM

                         │

                         ▼

              Generated Answer

                         │

                         ▼

                  HTTP Response
```

---

# 5. Current Implementation Status

### Completed

✓ Document Upload

✓ Chunk Generation

✓ Embedding Generation

✓ Embedding Storage

✓ pgvector Integration

✓ EmbedQuery gRPC

✓ Provider Abstraction

---

### In Progress

- Semantic Search Repository
- Context Builder
- Prompt Builder
- Chat Endpoint
- Citation Generation

---

### Planned

- Hybrid Search
- Conversation Memory
- Query Rewriting
- Re-ranking
- Streaming Responses

---

# 6. Step 1 – User Question

Example:

```
How does JWT authentication work?
```

The backend receives the question through a REST API.

At this stage, the backend does not attempt to answer the question.

Instead, it begins the retrieval process.

---

# 7. Step 2 – Query Embedding

The question is sent to the AI Service.

```
Question

↓

EmbedQuery

↓

Embedding Vector
```

This embedding represents the semantic meaning of the question.

Current provider:

- Ollama

Future providers:

- Gemini Embeddings
- OpenAI Embeddings
- BGE
- E5

---

# 8. Step 3 – Vector Search

The generated embedding is compared against all stored document embeddings.

```
Question Vector

↓

pgvector

↓

Similarity Search

↓

Nearest Chunks
```

Instead of keyword matching, pgvector compares mathematical distance.

---

# 9. Distance Metrics

pgvector supports multiple similarity calculations.

Examples:

- Cosine Similarity
- Inner Product
- Euclidean Distance

DocsAI currently uses semantic similarity to identify the closest document chunks.

---

# 10. Top-K Retrieval

Instead of returning every matching chunk, the repository returns only the most relevant results.

Example:

```
Question

↓

5000 Chunks

↓

Rank

↓

Top 5 Chunks
```

This dramatically reduces LLM context size.

---

# 11. Context Builder

Retrieved chunks are combined into one context.

Example:

```
Chunk 1

Chunk 2

Chunk 3

Chunk 4

↓

Combined Context
```

The context builder prepares structured information for prompt construction.

---

# 12. Prompt Builder

The prompt builder combines:

System Prompt

+

Retrieved Context

+

User Question

↓

Final Prompt

Example:

```
You are DocsAI.

Answer ONLY using the provided context.

If the answer cannot be found,
say you don't know.

Context:

...

Question:

...
```

This reduces hallucinations.

---

# 13. LLM Response Generation

The completed prompt is sent to the selected LLM.

Current plan:

Gemini

Future:

OpenAI

Claude

Mistral

Local Models

The LLM never receives the entire database.

It receives only the retrieved context.

---

# 14. Source Citations

Every retrieved chunk retains metadata.

Examples:

- Document Name
- Page Number
- Section
- Chunk ID

Future responses will include citations.

Example:

```
Authentication uses JWT.

Source:

Authentication.pdf

Page 17
```

This improves trust and explainability.

---

# 15. Why Chunks?

Sending entire documents to an LLM is inefficient.

Instead:

```
500 Pages

↓

Top 5 Chunks

↓

Prompt
```

Benefits:

- Lower cost
- Faster responses
- Better accuracy
- Lower hallucination rate

---

# 16. Why Store Embeddings?

Embeddings are generated once during document ingestion.

During search:

```
Question

↓

Embedding

↓

Compare

↓

Answer
```

This avoids regenerating embeddings for every document during each query.

---

# 17. Hybrid Search (Future)

Future retrieval combines:

```
Keyword Search

+

Vector Search

↓

Merge

↓

Rank

↓

LLM
```

Benefits:

- Better precision
- Better recall
- Improved handling of technical terms

---

# 18. Re-ranking (Future)

Instead of trusting pgvector alone:

```
Top 20 Results

↓

Cross Encoder

↓

Re-rank

↓

Top 5
```

This improves answer quality.

---

# 19. Conversation Memory (Future)

Future conversations will maintain context.

```
Question 1

↓

Answer

↓

Question 2

↓

Conversation Memory

↓

Improved Retrieval
```

Users won't need to repeat themselves.

---

# 20. Multi-document Retrieval

Future architecture:

```
Workspace

↓

Multiple Documents

↓

Semantic Search

↓

Relevant Chunks

↓

Combined Context

↓

Answer
```

The AI will reason across many uploaded documents.

---

# 21. Confidence Scoring

Every retrieval result can include a similarity score.

Example:

```
Chunk A

96%

Chunk B

92%

Chunk C

88%
```

Low-confidence retrievals may prompt the AI to respond more cautiously.

---

# 22. Failure Handling

Possible failures include:

- No relevant chunks
- Embedding provider unavailable
- LLM timeout
- Empty document
- Unsupported language

The backend should return meaningful responses while preserving system stability.

---

# 23. Planned Knowledge Module

As DocsAI evolves, retrieval responsibilities should be separated from document management.

Proposed architecture:

```
Knowledge Module

↓

Embed Query

↓

Retrieve Chunks

↓

Build Context

↓

Prompt Builder

↓

LLM

↓

Response
```

The Document Module remains responsible for uploads and metadata.

---

# 24. Engineering Principles

The RAG engine follows these principles:

✓ Grounded Responses

✓ Explainable AI

✓ Source Attribution

✓ Provider Independence

✓ Scalable Retrieval

✓ Low Hallucination Risk

✓ Modular Architecture

✓ Extensible Design

---

# 25. Long-Term Vision

The RAG engine is only the first step toward a complete Knowledge Intelligence Platform.

Future capabilities include:

- Knowledge Graphs
- AI Agents
- Workflow Automation
- Enterprise Search
- Cross-document Reasoning
- Semantic Recommendations
- Personalized AI Assistants
- Organization-wide Knowledge Memory

---

# 26. Summary

The Retrieval-Augmented Generation architecture transforms DocsAI from a document storage platform into an intelligent knowledge system.

By combining semantic retrieval, vector search, prompt engineering, and modern LLMs, DocsAI enables users to interact with their own information through natural conversation while maintaining accuracy, explainability, and trust.

The RAG engine is the foundation upon which all future AI capabilities of DocsAI will be built.