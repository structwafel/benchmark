# Nature Language Benchmark

This repository contains performance benchmarks comparing the Nature programming language with other mainstream languages.

## Test Environment

| Item | Configuration |
|------|---------------|
| **Device** | Mac mini M4 (VM) |
| **CPU** | Apple M4, 10 cores |
| **Memory** | 7.80 GB |
| **System** | Linux ubuntu 6.17.8-orbstack (aarch64) |

---

## Benchmark Results

### 1. HTTP Server

HTTP server performance using ApacheBench (`ab -n 100000 -c 1000`).

| Language | Version | QPS (req/sec) | Mean Response Time |
|----------|---------|---------------|---------------------|
| Nature | v0.7.2 | ~104,000 | 9.60 ms |
| Go | 1.25.5 | ~90,000 | 11.02 ms |
| Node.js | v25.2.0 | ~66,000 | 14.99 ms |

---

### 2. Fibonacci Recursion

Calculate `fib(45)`, testing recursion performance and function call overhead.

| Language | Version | Mean Time |
|----------|---------|-----------|
| Rust | 1.92.0 | ~1.758 s |
| Nature | v0.7.2 | ~2.453 s |
| Go | 1.25.5 | ~2.421 s |
| Node.js | v25.2.0 | ~6.198 s |

---

### 3. Pi Calculation

Calculate π using Leibniz series, testing floating-point arithmetic performance.

| Language | Version | Mean Time |
|----------|---------|-----------|
| Nature | v0.7.2 | ~512.9 ms |
| Go | 1.25.5 | ~512.3 ms |
| Rust | 1.92.0 | ~552.6 ms |
| Node.js | v25.2.0 | ~850.8 ms |

---

### 4. FFI Calls

Call C library function `sqrt()` 100 million times, testing FFI call overhead.

| Language | Version | Mean Time |
|----------|---------|-----------|
| Nature | v0.7.2 | ~73.9 ms |
| Go | 1.25.5 | ~2178 ms |

---

### 5. Coroutine

Create 1 million coroutines for concurrent execution, testing coroutine creation efficiency and memory usage.

| Language | Version | Creation Time | Completion Time | Peak Memory |
|----------|---------|---------------|-----------------|-------------|
| Nature | v0.7.2 | ~559 ms | ~589 ms | ~969 MB |
| Go | 1.25.5 | ~1035 ms | ~1047 ms | ~2580 MB |

---

## Test Directories

Each test directory contains source code for each language and detailed test results:

- [`http_server/`](./http_server/) - HTTP server test
- [`fib/`](./fib/) - Fibonacci recursion test
- [`pi/`](./pi/) - Pi calculation test  
- [`ffi/`](./ffi/) - FFI call test
- [`coroutine/`](./coroutine/) - Coroutine test
