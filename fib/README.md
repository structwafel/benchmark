
## Macmini m4 virtual machine

CPU cores: 10

Memory: 7.80745 GB

Linux ubuntu 6.17.8-orbstack-00308-g8f9c941121b1 #1 SMP PREEMPT Thu Nov 20 09:34:02 UTC 2025 aarch64 aarch64 aarch64 GNU/Linux

hyperfine --warmup 3 --runs 5 ./fib_n ./fib_go ./fib_rs "node fib.js"

nature build -o fib_n main.n

go build -o fib_go main.go

rustc -C opt-level=3 fib.rs -o fib_rust


## nature v0.7.2

```
Benchmark 1: ./fib_n
  Time (mean ± σ):      2.453 s ±  0.010 s    [User: 2.451 s, System: 0.002 s]
  Range (min … max):    2.442 s …  2.466 s    5 runs
```

## go1.25.5

```
Benchmark 2: ./fib_go
  Time (mean ± σ):      2.421 s ±  0.021 s    [User: 2.419 s, System: 0.003 s]
  Range (min … max):    2.401 s …  2.455 s    5 runs
```

## rust 1.92.0

```
Benchmark 3: ./fib_rs
  Time (mean ± σ):      1.758 s ±  0.017 s    [User: 1.758 s, System: 0.000 s]
  Range (min … max):    1.745 s …  1.786 s    5 runs
 
```

## node.js v25.2.0

```
Benchmark 4: node fib.js
  Time (mean ± σ):      6.198 s ±  0.233 s    [User: 6.192 s, System: 0.006 s]
  Range (min … max):    5.913 s …  6.557 s    5 runs

```