# gocache

An in-memory cache server built from scratch in Go. No frameworks, no external dependencies — raw TCP socket, standard library only.

## What it is

A mini Redis. Clients connect over TCP and send commands as plain text. Data lives in memory with optional TTL expiration.

**Supported commands:**
```
SET key value       — store a key
GET key             — retrieve a key
DEL key             — delete a key
EXPIRE key seconds  — delete key after N seconds
```

## How to run

```bash
go run .
```

Connect with netcat:
```bash
nc localhost 6379
```

## How it works

- Raw TCP server using `net.Listen` — no HTTP layer
- Goroutine per client connection
- Protocol parser reads raw bytes, splits into command/key/value
- Key-value store backed by `map[string]string`
- `sync.RWMutex` — concurrent reads, exclusive writes
- Background goroutine sweeps expired keys every second

## Benchmark

Tested on Apple M3 (arm64):

```
BenchmarkSet-8    89319027    13.08 ns/op    (~89M ops/sec)
BenchmarkGet-8    149885625    7.908 ns/op   (~149M ops/sec)
```

```bash
go test -bench=. ./store/
```
