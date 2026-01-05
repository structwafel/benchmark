## Macmini m4 virtual machine

CPU cores: 10

Memory: 7.80745 GB

Linux ubuntu 6.17.8-orbstack-00308-g8f9c941121b1 #1 SMP PREEMPT Thu Nov 20 09:34:02 UTC 2025 aarch64 aarch64 aarch64 GNU/Linux

## nature v0.7.1

```
hyperfine ./ffi_n
Benchmark 1: ./ffi_n
  Time (mean ± σ):      73.9 ms ±   3.9 ms    [User: 75.0 ms, System: 0.6 ms]
  Range (min … max):    71.1 ms …  94.3 ms    31 runs
```


## go1.25.5

```
hyperfine ./ffi_go
Benchmark 1: ./ffi_go
  Time (mean ± σ):      2.178 s ±  0.031 s    [User: 2.176 s, System: 0.003 s]
  Range (min … max):    2.116 s …  2.210 s    10 runs
```