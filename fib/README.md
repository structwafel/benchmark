
## Macmini m4 virtual machine

CPU cores: 10

Memory: 7.80745 GB

Linux ubuntu 6.17.8-orbstack-00308-g8f9c941121b1 #1 SMP PREEMPT Thu Nov 20 09:34:02 UTC 2025 aarch64 aarch64 aarch64 GNU/Linux

## nature v0.7.1

```
hyperfine ./fib_nature 
Benchmark 1: ./fib_nature
  Time (mean ± σ):      2.340 s ±  0.025 s    [User: 2.370 s, System: 0.011 s]
  Range (min … max):    2.309 s …  2.375 s    10 runs
```

## go1.25.5

```
hyperfine ./fib_go
Benchmark 1: ./fib_go
  Time (mean ± σ):      2.341 s ±  0.013 s    [User: 2.340 s, System: 0.002 s]
  Range (min … max):    2.324 s …  2.367 s    10 runs
```

## rust 1.92.0

```
hyperfine ./fib_rust
Benchmark 1: ./fib_rust
  Time (mean ± σ):      1.712 s ±  0.015 s    [User: 1.711 s, System: 0.001 s]
  Range (min … max):    1.697 s …  1.742 s    10 runs
 
```

## node.js v25.2.0

```
hyperfine "node fib.js"
Benchmark 1: node fib.js
  Time (mean ± σ):      6.036 s ±  0.128 s    [User: 6.022 s, System: 0.014 s]
  Range (min … max):    5.927 s …  6.347 s    10 runs

```