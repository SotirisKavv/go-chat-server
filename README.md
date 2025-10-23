# 💬 Real-Time Go Chat Server 🚀

Welcome to the **Real-Time Go Chat Server**, a high-performance, full-stack WebSocket chat application built in Go. Whether you're a recruiter, fellow developer, or curious hacker, this README will give you a whirlwind tour of what this project does, how it works, and why it’s a stellar example of modern Go engineering.

---

## 🤔 What Is This?
A lightning-fast, multi-room chat server that handles thousands of concurrent users with ease. Powered by:

- **WebSockets** for instant bidirectional messaging
- **PostgreSQL** persistence (with in-memory fallback)
- A **custom Query Builder** for safe, parameterized SQL
- **Goroutines & Channels** for supercharged concurrency

Think Slack meets Go—only lighter, faster, and 100% open source.

---

## 🏗️ Architecture Overview

```text
Client Browser  ↔  WebSocket Handler  ↔  Hub & Workers  ↔  Storage Layer
                                             m
                                    +-------- In-Memory --------+
                                    |                           |
                               PostgreSQL                 Message Queue
```

1. **Client**: HTML/JS front-end connects via `ws://localhost:8080/chat`.
2. **Router & Handler**: `gorilla/mux` routes HTTP & WebSocket endpoints.
3. **Hub**: Central goroutine manages rooms and broadcasts messages.
4. **Workers**: Dedicated goroutines save messages to storage asynchronously.
5. **Storage**:
   - **PostgreSQL** via `pgx` with auto-schema creation
   - **In-Memory** fallback store for ultra-low-latency reads
6. **Repository Pattern** abstracts storage details and injects the proper layer.

---

## ✨ Go Features & Patterns Used

- **Concurrency**: Lightweight goroutines & channel fan-out/fan-in patterns
- **Design Patterns**:
  - **Repository** for data access abstractions
  - **Singleton** for a single Hub instance
  - **Observer** for subscription/event handling in rooms
  - **Builder** for SQL query construction (`utils/QueryBuilder`)
- **Error Handling**: Idiomatic Go error checks, `defer` cleanup
- **Environment Config**: `godotenv` for local `.env` support
- **Module Management**: Go Modules (`go.mod`, `go.sum`)

---

## 🚀 Getting Started

### Prerequisites

- Go 1.23+
- PostgreSQL (if you want persistence) or none (in-memory only)
- `make` or your preferred task runner (optional)

### Quick Start

```powershell
# Clone this repo
git clone https://github.com/yourusername/chat-server.git
cd chat-server

# Copy the sample .env (optional)
cp .env.example .env
# Edit .env for your DB credentials:
# DATABASE_URL=postgres://user:pass@localhost:5432/chatdb?sslmode=disable

# Install dependencies & build
go mod tidy
go build -o chat-server.exe .

# Run the server
.\\chat-server.exe
```

Then open your browser at `http://localhost:8080/client.html` and join a room:

- Enter **username** & **room**
- Chat away! 🎉

---

## 🔧 Configuration Options

| Setting           | Default                             | Description                            |
|-------------------|-------------------------------------|----------------------------------------|
| `DATABASE_URL`    | (unset)                             | PostgreSQL connection string           |
| In-Memory Mode    | if `DATABASE_URL` unset             | Messages stored only in memory         |
| `PORT`            | `8080`                              | HTTP/WebSocket listen port             |

---

## 🚩 Endpoints

- **WebSocket**: `ws://localhost:8080/chat?r=<room>&u=<user>`
- **History API**: `GET /history?r=<room>` (JSON array of messages)
- **Static Client**: `GET /client.html`

---

## 🦄 Contributors & Credits

Author: IAmSotiris.


