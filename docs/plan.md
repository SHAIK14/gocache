# gocache — Build Plan

Mini Redis from scratch in Go. TCP server, not HTTP. Standard library only.

## Layers

| # | Layer | Status | Key concept |
|---|-------|--------|-------------|
| 1 | TCP server | 🔲 | net.Listen + net.Conn + goroutine per client |
| 2 | Protocol parser | 🔲 | split raw text → command / key / value |
| 3 | Storage engine | 🔲 | map[string]string |
| 4 | Concurrency safety | 🔲 | sync.RWMutex — read lock GET, write lock SET/DEL |
| 5 | TTL expiration | 🔲 | expiry time per key + background sweeper goroutine |

## Commands to support
SET key value | GET key | DEL key | EXPIRE key 30

## Resume target
"Built an in-memory cache server from scratch in Go. TCP protocol, concurrent connections,
TTL expiration, mutex-based thread safety. Benchmarked at X ops/sec."

## Status legend
🔲 not started | 🔄 in progress | ✅ done
