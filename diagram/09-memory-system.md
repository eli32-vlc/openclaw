# Memory System

How OpenClaw indexes, embeds, and searches agent memory using a hybrid vector + full-text search engine backed by SQLite.

```mermaid
flowchart TD
    MemoryConfig["Memory config\nconfig.agents.memory"]
    MemoryManager["Memory Manager\nsrc/memory/manager.ts"]
    EmbeddingProvider["Embedding Provider\nsrc/memory/embeddings.ts"]
    SyncOps["Sync Operations\nsrc/memory/manager-sync-ops.ts"]
    EmbeddingOps["Embedding Operations\nsrc/memory/manager-embedding-ops.ts"]
    SearchManager["Search Manager\nsrc/memory/search-manager.ts"]
    QMDProcess["QMD Processor\nsrc/memory/qmd-process.ts"]
    QMDParser["QMD Query Parser\nsrc/memory/qmd-query-parser.ts"]
    SQLiteDB["SQLite DB\n(chunks_vec, chunks_fts, embedding_cache)"]
    FsUtils["FS Utilities\nsrc/memory/fs-utils.ts"]
    Internal["Internal Helpers\nsrc/memory/internal.ts"]

    MemoryConfig --> MemoryManager
    MemoryManager --> EmbeddingProvider
    MemoryManager --> SyncOps
    MemoryManager --> EmbeddingOps
    MemoryManager --> SearchManager
    MemoryManager --> SQLiteDB
    SyncOps --> FsUtils
    SyncOps --> QMDProcess
    QMDProcess --> SQLiteDB
    SearchManager --> QMDParser
    SearchManager --> SQLiteDB
    EmbeddingOps --> EmbeddingProvider
    EmbeddingOps --> SQLiteDB
    Internal --> FsUtils
```

## Embedding Providers

```mermaid
graph TD
    EmbeddingProvider["EmbeddingProvider\nsrc/memory/embeddings.ts"]
    EmbeddingProvider --> OpenAI["OpenAI\nsrc/memory/embeddings-openai.ts"]
    EmbeddingProvider --> Mistral["Mistral\nsrc/memory/embeddings-mistral.ts"]
    EmbeddingProvider --> Ollama["Ollama (local)\nsrc/memory/embeddings-ollama.ts"]
    EmbeddingProvider --> Voyage["Voyage AI\nsrc/memory/batch-voyage.ts"]
    EmbeddingProvider --> RemoteProvider["Remote HTTP\nsrc/memory/embeddings-remote-provider.ts"]
    EmbeddingProvider --> ModelNormalize["Model normalization\nsrc/memory/embeddings-model-normalize.ts"]

    OpenAI --> BatchOpenAI["Batch HTTP\nsrc/memory/batch-openai.ts"]
    Voyage --> BatchVoyage["Batch HTTP\nsrc/memory/batch-voyage.ts"]
    RemoteProvider --> BatchHTTP["Generic Batch\nsrc/memory/batch-http.ts"]
```

## Hybrid Search Pipeline

```mermaid
flowchart LR
    Query["User query string"]
    Expand["Query Expansion\nsrc/memory/query-expansion.ts"]
    VectorSearch["Vector search\n(ANN via sqlite-vec)"]
    FTSSearch["Full-text search\n(FTS5 BM25)"]
    MMR["MMR Re-rank\nsrc/memory/mmr.ts"]
    Results["MemorySearchResult[]"]

    Query --> Expand
    Expand --> VectorSearch
    Expand --> FTSSearch
    VectorSearch --> MMR
    FTSSearch --> MMR
    MMR --> Results
```

## Memory Index Lifecycle

```mermaid
stateDiagram-v2
    [*] --> Initializing : getOrOpenMemoryIndex()
    Initializing --> Syncing : syncFiles()
    Syncing --> Embedding : embedPendingChunks()
    Embedding --> Ready : index warm
    Ready --> Searching : searchMemory()
    Searching --> Ready : results returned
    Ready --> Syncing : file watcher detects change
    Ready --> [*] : closeAllMemoryIndexManagers()
```

## SQLite Schema

```mermaid
erDiagram
    CHUNKS {
        int id PK
        text path
        int chunk_index
        text content
        int updated_at
    }
    CHUNKS_VEC {
        int chunk_id FK
        blob embedding
    }
    CHUNKS_FTS {
        int rowid FK
        text content
    }
    EMBEDDING_CACHE {
        text hash PK
        blob vector
        int created_at
    }

    CHUNKS ||--o{ CHUNKS_VEC : "has vector"
    CHUNKS ||--o{ CHUNKS_FTS : "indexed in FTS"
```

## Key Files

| File | Role |
|------|------|
| `src/memory/manager.ts` | Top-level memory manager (open/close/search) |
| `src/memory/embeddings.ts` | Abstract embedding provider factory |
| `src/memory/embeddings-openai.ts` | OpenAI text-embedding-3-* |
| `src/memory/embeddings-mistral.ts` | Mistral embed models |
| `src/memory/embeddings-ollama.ts` | Local Ollama embedding models |
| `src/memory/batch-voyage.ts` | Voyage AI batch embedder |
| `src/memory/manager-sync-ops.ts` | Sync markdown files into chunk store |
| `src/memory/manager-embedding-ops.ts` | Compute and cache embeddings |
| `src/memory/search-manager.ts` | Hybrid search (vector + FTS) |
| `src/memory/mmr.ts` | Maximal Marginal Relevance re-ranking |
| `src/memory/qmd-process.ts` | QMD (query memory document) processor |
| `src/memory/qmd-query-parser.ts` | Parse QMD query syntax |
