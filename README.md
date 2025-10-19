# Go Practice Tasks (1–15)

A curated set of practical Go exercises covering HTTP APIs, databases, concurrency (channels, worker pools, pipelines), testing, caching, realtime sockets, streaming, and systems programming. Each task includes a brief goal, suggested tech, and a clear Definition of Done (DoD). Pick a few and iterate!

## Project Structure

```
go.work
/libs/clock/           # โมดูลโค้ดกลาง (เช่น clock, logger, testkit)
/libs/logger/
/tasks/01-crud-api/    # โมดูลโจทย์ 1
/tasks/02-repo-tx/
/tasks/03-worker-pool/
/tasks/04-csv-pipeline/
/tasks/05-cache/
/tasks/06-graceful-server/
/tasks/07-s3-upload/
/tasks/08-grpc-stream/
/tasks/09-rate-limiter/
/tasks/10-testing-suite/
/tasks/11-websocket-chat/
/tasks/12-notification-hub/
/tasks/13-tcp-proxy/
/tasks/14-udp-syslog/
/tasks/15-ws-grpc-gateway/
```

## Tasks

### 1) CRUD REST API + Core Middleware
**Goal:** Build a Todo (or "Securities") REST API with list (paging/filter/sort), get, create, update, delete.

**Tech:** `net/http` + `chi` or `gin`, `go-playground/validator`, JWT.

**DoD:** Request ID, structured logging, panic-recover, validation errors → problem+json, JWT auth, table-driven handler tests.

---

### 2) Database Integration + Repository + Transactions
**Goal:** Implement repository methods with safe dynamic filters, ORDER BY, and keyset pagination; one use case requiring all-or-nothing Tx.

**Tech:** `sqlx` (or `pgx` / MSSQL driver).

**DoD:** Context timeouts for queries, Tx rollback on failure, index-friendly queries, integration test with Testcontainers or docker-compose.

---

### 3) Worker Pool + Rate Limit + Cancellation
**Goal:** Concurrently fetch ~200 URLs with N workers; limit RPS; abort remaining work if errors > K.

**Tech:** channels, `sync.WaitGroup`, context, backoff/retry.

**DoD:** Deterministic unit tests (fake HTTP server), graceful cancel on threshold, metrics: success/error/latency.

---

### 4) CSV Streaming Pipeline (Fan-out/Fan-in)
**Goal:** CSV → validate → transform → persist using goroutines and channels with backpressure.

**Tech:** `encoding/csv`, select, bounded channels.

**DoD:** Separate error channel, clean channel closes, final metrics summary (rows ok/failed), property-based tests for transform.

---

### 5) In-Memory Cache (Generic) with TTL + Bench & Race
**Goal:** Generic cache `Cache[K comparable, V any]` with TTL per key, background eviction, and metrics.

**Tech:** generics, `sync.RWMutex` (or sharded map), time wheel/heap.

**DoD:** Benchmarks (`testing.B`), `-race` clean, 100% deterministic unit tests via injectable clock.

---

### 6) Graceful HTTP Server + Timeouts + Health/Ready + pprof
**Goal:** Production-style server with timeouts, `/healthz` (liveness), `/readyz` (DB ping), pprof, graceful shutdown.

**Tech:** `http.Server`, signals (SIGINT/SIGTERM), `/debug/pprof`.

**DoD:** No dropped requests on shutdown (drain with WG), failing DB toggles readiness, blackbox tests using an ephemeral port.

---

### 7) S3 File Upload Service + Retries + Mock Tests
**Goal:** Stream file uploads to S3 (multipart for large files), retries with exponential backoff + jitter, deadlines.

**Tech:** AWS SDK v2, interfaces for mocking.

**DoD:** Unit tests with mocked S3, retry policy verified, context deadlines enforced, large file path covered.

---

### 8) gRPC Bi-Directional Streaming + Interceptors
**Goal:** "PriceStream" bi-di: client subscribes symbols, server pushes updates.

**Tech:** gRPC, unary/stream interceptors (logging/auth/deadline).

**DoD:** Example CLI client, deadline propagation, reconnection logic, stream cancellation handled cleanly.

---

### 9) Rate Limiter Middleware (Per-IP & Per-User)
**Goal:** Token-bucket limiter layered by IP and by userId (from JWT).

**Tech:** sharded map or `sync.Map`, injectable time source for tests.

**DoD:** Concurrency-safe, starvation-free, unit tests simulate bursts, configurable limits.

---

### 10) Testing Suite: Unit, Table-Driven, Fuzz, Integration
**Goal:**
- Pure functions (e.g., interest/principal calc, date math, text normalize) with table-driven and fuzz tests.
- Integration tests for Task #2 using containers.
- Service layer mocked via gomock.

**Tech:** `testing`, `testing/quick`, `gomock`, Testcontainers.

**DoD:** Meaningful coverage, fuzz finds at least one fixed bug or verified invariants, CI script to run all.

---

### 11) WebSocket Chat: Rooms + DM + Ping/Pong
**Goal:** WS server with rooms, join/leave, broadcast & private DM, heartbeat to prune dead clients.

**Tech:** `nhooyr.io/websocket` or `gorilla/websocket`.

**DoD:** HTML demo client, backpressure handling (bounded outbound queues), auth via JWT header, metrics for online/users/rooms.

---

### 12) Realtime Notifications Hub (Server → Browser)
**Goal:** REST/queue ingests events; per-user ring buffer; push via WS; retry on reconnect.

**Tech:** WebSocket, per-user queues, topic filters.

**DoD:** Simulated 100 clients integration test, p95 latency measured, rate limit per user, graceful shutdown without event loss (within policy).

---

### 13) TCP Proxy (Reverse Tunnel) with Health Check & Drain
**Goal:** TCP proxy from port A to B; failover to B2 if B unhealthy; graceful draining on shutdown.

**Tech:** `net`, `io.Copy` with context, keepalive, health probes.

**DoD:** Load test (wrk/hey for HTTP over proxy or netcat raw), hot-reload config on SIGHUP, zero data races.

---

### 14) UDP Syslog Collector + Batched DB Writes
**Goal:** Receive syslog over UDP, parse, batch-insert into DB on N items or T seconds, with drop policy.

**Tech:** `net.ListenUDP`, bounded channels, `sqlx` batch.

**DoD:** Fuzz tests for parser, throughput benchmark, alert/log when drops exceed threshold, backpressure respected.

---

### 15) WebSocket ↔ gRPC Streaming Gateway
**Goal:** Bridge browser WS clients to backend gRPC bi-di streams, mapping auth/metadata both ways.

**Tech:** WS server + gRPC client streams, interceptors, flow control.

**DoD:** Example web client + integration tests, correct cancel/timeout translation, resilience to transient gRPC failures.

---

## Getting Started

1. Each task is in its own module under `/tasks/`
2. Shared utilities are in `/libs/`
3. Use `go.work` for multi-module workspace management
4. Run tests: `go test ./...` from workspace root
5. Run with race detector: `go test -race ./...`

## Best Practices Covered

- ✅ Structured logging & request tracing
- ✅ Context propagation & cancellation
- ✅ Graceful shutdown patterns
- ✅ Concurrent programming (channels, worker pools, pipelines)
- ✅ Testing strategies (unit, integration, table-driven, fuzz, mocks)
- ✅ Production readiness (health checks, metrics, pprof)
- ✅ Error handling & retry strategies
- ✅ Database transactions & safe query building
- ✅ Real-time communication (WebSocket, gRPC streaming)
- ✅ Systems programming (TCP/UDP, proxies)
