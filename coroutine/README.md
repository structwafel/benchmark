## Macmini m4 virtual machine

CPU cores: 10

Memory: 7.80745 GB

Linux ubuntu 6.17.8-orbstack-00308-g8f9c941121b1 #1 SMP PREEMPT Thu Nov 20 09:34:02 UTC 2025 aarch64 aarch64 aarch64 GNU/Linux

## nature v0.7.1

./coroutine_n
559
559 968016
588 974332
589

    PID+USER       PRI  NI  VIRT   RES   SHR S  CPU% MEM%   TIME+  Command (merged)
274640 root        20   0 1722M  969M   572 S   0.0 12.1  0:00.29 coroutine_n|./coroutine_n

## go1.25.5

./coroutine_go // have sleep
1035
1035 948113
1046 949043
1047

    PID+USER       PRI  NI  VIRT   RES   SHR S  CPU% MEM%   TIME+  Command (merged)
274033 root        20   0 3721M 2580M  1136 S   0.0 32.3  0:00.12 coroutine_go|./coroutine_go


