## nature

```
hyperfine ./pi_n
Benchmark 1: ./pi_n
  Time (mean ± σ):     762.4 ms ±   4.7 ms    [User: 771.5 ms, System: 2.4 ms]
  Range (min … max):   756.4 ms … 771.7 ms    10 runs

```

## golang

```
hyperfine ./pi_go
Benchmark 1: ./pi_go
  Time (mean ± σ):     991.2 ms ±   8.5 ms    [User: 990.8 ms, System: 1.1 ms]
  Range (min … max):   986.5 ms … 1014.3 ms    10 runs

```


## rust

```
hyperfine ./pi_rs
Benchmark 1: ./pi_rs
  Time (mean ± σ):     547.6 ms ±   6.5 ms    [User: 547.2 ms, System: 0.3 ms]
  Range (min … max):   544.0 ms … 565.9 ms    10 runs
```

## node.js

```
hyperfine "node main.js"
Benchmark 1: node main.js
  Time (mean ± σ):     837.8 ms ±  31.6 ms    [User: 834.2 ms, System: 4.3 ms]
  Range (min … max):   819.3 ms … 925.7 ms    10 runs
```