# Real‑Time Go Chat Server — Event‑Driven WebSocket Rooms

High‑performance, multi‑room WebSocket chat server in Go. Shows off clean concurrency with goroutines/channels, a repository‑backed storage layer (Postgres or in‑memory), and a tiny query builder for safe SQL.

- Endpoints: WebSocket for chat + HTTP for message history and static client
- Messaging: Broadcast per room via a central Hub
- Persistence: PostgreSQL (via pgx) with in‑memory fallback
- Concurrency: Fan‑out/fan‑in workers for async persistence

Quick links:
- Client: `client.html`
- WebSocket handler: `server/websocket.go`
- Hub and rooms: `chat/hub.go`, `chat/room.go`, `chat/client.go`
- Storage: `storage/pgsql_repository.go`, `storage/memory_repo.go`


## What this project showcases

- Lightweight, idiomatic concurrency with goroutines and channels
- Repository pattern to abstract storage (Postgres or memory)
- Parameterized SQL with a minimal Query Builder (`utils/query_builder.go`)
- Clean WebSocket handling and room fan‑out
- Simple, explicit error handling and resource cleanup


## High‑level architecture

Components:
- Client (HTML/JS): Connects via WebSocket to join/send messages
- Router + WebSocket handler: Upgrades connection and attaches to a room
- Hub: Central coordinator for rooms and broadcasts
- Workers: Async persistence of messages to storage
- Storage: PostgreSQL (pgx) or in‑memory repository

Diagram:

```text
Client Browser  ↔  WebSocket Handler  ↔  Hub & Workers  ↔  Storage Layer
                                    +-------- In‑Memory --------+
                                    |                           |
                               PostgreSQL                 (optional MQ)
```


## How it works (chat flow)

1) Client connects to `ws://localhost:8080/chat?r=<room>&u=<user>`
- Server upgrades to WebSocket and registers the client in a room

2) Hub broadcasts messages per room
- Incoming messages are fanned out to all clients in the same room

3) Workers persist asynchronously
- Messages are queued to a worker pool to write to the configured repository

4) History API serves past messages
- `GET /history?r=<room>` returns recent messages for the room


## Design patterns and tactics

- Repository abstraction for data access (`storage/*`)
- Singleton‑like Hub instance coordinating rooms and clients
- Observer‑style subscriptions for room broadcasts
- Builder pattern for composing SQL queries (`utils/query_builder.go`)
- Defensive error handling with `defer` cleanup


## Technologies

- Go (net/http, Gorilla Mux), pgx (Postgres)
- WebSockets for real‑time messaging
- Docker‑friendly layout (single binary)


## Run locally

Prereqs: Go 1.23+; Postgres optional (in‑memory if `DATABASE_URL` is unset)

1) Build and run

```powershell
# From this folder
go mod tidy
go build -o chat-server.exe .
./chat-server.exe
```

2) Configure (optional)

If you want Postgres persistence, set the following env vars (or a `.env` if you use `godotenv`):

```powershell
$env:DATABASE_URL = "postgres://user:pass@localhost:5432/chatdb?sslmode=disable"
$env:PORT = "8080"
```

3) Verify

- Open http://localhost:8080/client.html
- Join a room, send messages, and open a second tab to see broadcasts


## Endpoints

- WebSocket: `ws://localhost:8080/chat?r=<room>&u=<user>`
- History API: `GET /history?r=<room>` (JSON array of messages)
- Static Client: `GET /client.html`


## Configuration

| Setting        | Default                 | Description                          |
|----------------|-------------------------|--------------------------------------|
| `DATABASE_URL` | (unset)                 | PostgreSQL connection string         |
| `PORT`         | `8080`                  | HTTP/WebSocket listen port           |
| In‑Memory      | if `DATABASE_URL` unset | Messages stored only in memory       |


## Folder map

- `client.html`: Minimal browser client
- `main.go`: Program entrypoint and wiring
- `server/`: HTTP router and WebSocket handler
- `chat/`: Hub, room, and client types
- `model/`: Message model(s)
- `storage/`: Repositories (Postgres + in‑memory)
- `utils/`: Query builder


## Notable implementation details

- Hub coordinates room membership and broadcasts; workers persist messages without blocking writes
- Postgres repository uses parameterized queries via the Query Builder
- If `DATABASE_URL` is not provided, storage falls back to an in‑memory repository for ultra‑low latency


## Next steps (ideas)

- Add auth and per‑room ACLs
- Add message pagination and retention policies
- Add metrics (connections per room, broadcast latency) and tracing
- Add Dockerfile and compose for easy Postgres spins


## Contributors & Credits

Author: IAmSotiris.
````


