## Macmini m4 virtual machine

CPU cores: 10

Memory: 7.80745 GB

Linux ubuntu 6.17.8-orbstack-00308-g8f9c941121b1 #1 SMP PREEMPT Thu Nov 20 09:34:02 UTC 2025 aarch64 aarch64 aarch64 GNU/Linux


command hyperfine --warmup 3 ./pi_n ./pi_go ./pi_rs "node main.js"

## nature v0.7.2

```
Benchmark 1: ./pi_n
  Time (mean ± σ):     512.9 ms ±   1.3 ms    [User: 510.5 ms, System: 1.2 ms]
  Range (min … max):   511.8 ms … 516.4 ms    10 runs

```

## go1.25.5

```
Benchmark 2: ./pi_go
  Time (mean ± σ):     512.3 ms ±   1.4 ms    [User: 511.6 ms, System: 1.2 ms]
  Range (min … max):   510.9 ms … 515.7 ms    10 runs

```


## rust 1.92.0

```
  Time (mean ± σ):     552.6 ms ±  12.9 ms    [User: 551.6 ms, System: 0.5 ms]
  Range (min … max):   544.0 ms … 583.5 ms    10 runs
```

## node.js v25.2.0

```
Benchmark 4: node main.js
  Time (mean ± σ):     850.8 ms ±   5.6 ms    [User: 846.8 ms, System: 4.9 ms]
  Range (min … max):   846.0 ms … 864.8 ms    10 runs
```