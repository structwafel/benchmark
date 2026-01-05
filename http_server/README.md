## Macmini m4 virtual machine

CPU cores: 10

Memory: 7.80745 GB

Linux ubuntu 6.17.8-orbstack-00308-g8f9c941121b1 #1 SMP PREEMPT Thu Nov 20 09:34:02 UTC 2025 aarch64 aarch64 aarch64 GNU/Linux


## nature v0.7.1

ab -n 100000 -c 1000 http://127.0.0.1:8888/
This is ApacheBench, Version 2.3 <$Revision: 1913912 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        
Server Hostname:        127.0.0.1
Server Port:            8888

Document Path:          /
Document Length:        11 bytes

Concurrency Level:      1000
Time taken for tests:   0.960 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      11000000 bytes
HTML transferred:       1100000 bytes
Requests per second:    104185.88 [#/sec] (mean)
Time per request:       9.598 [ms] (mean)
Time per request:       0.010 [ms] (mean, across all concurrent requests)
Transfer rate:          11191.84 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    3   1.1      3      12
Processing:     1    6   1.6      6      19
Waiting:        1    5   1.6      5      17
Total:          3   10   1.5      9      20

Percentage of the requests served within a certain time (ms)
  50%      9
  66%     10
  75%     10
  80%     10
  90%     11
  95%     12
  98%     14
  99%     15
 100%     20 (longest request)

## go1.25.5

```
ab -n 100000 -c 1000 http://127.0.0.1:8888/
This is ApacheBench, Version 2.3 <$Revision: 1913912 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        
Server Hostname:        127.0.0.1
Server Port:            8888

Document Path:          /
Document Length:        11 bytes

Concurrency Level:      1000
Time taken for tests:   1.102 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      12800000 bytes
HTML transferred:       1100000 bytes
Requests per second:    90768.73 [#/sec] (mean)
Time per request:       11.017 [ms] (mean)
Time per request:       0.011 [ms] (mean, across all concurrent requests)
Transfer rate:          11346.09 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    5   1.4      5      16
Processing:     1    6   1.7      6      17
Waiting:        0    4   1.7      4      13
Total:          1   11   2.1     11      30

Percentage of the requests served within a certain time (ms)
  50%     11
  66%     11
  75%     12
  80%     12
  90%     13
  95%     14
  98%     17
  99%     19
 100%     30 (longest request)
 ```

## node.js v25.2.0

```
ab -n 100000 -c 1000 http://127.0.0.1:8888/
This is ApacheBench, Version 2.3 <$Revision: 1913912 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        
Server Hostname:        127.0.0.1
Server Port:            8888

Document Path:          /
Document Length:        12 bytes

Concurrency Level:      1000
Time taken for tests:   1.499 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      11300000 bytes
HTML transferred:       1200000 bytes
Requests per second:    66731.62 [#/sec] (mean)
Time per request:       14.985 [ms] (mean)
Time per request:       0.015 [ms] (mean, across all concurrent requests)
Transfer rate:          7363.94 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    5  72.7      0    1048
Processing:     2    8  19.1      7     421
Waiting:        1    8  19.1      7     421
Total:          3   13  87.4      7    1468

Percentage of the requests served within a certain time (ms)
  50%      7
  66%      7
  75%      7
  80%      7
  90%      9
  95%     10
  98%     12
  99%     22
 100%   1468 (longest request)
 ```