package main

import (
	"time"
)

var count = 0

func sum_co() {
	count += 1
	time.Sleep(10000 * time.Millisecond)
}

func main() {
	start := time.Now().UnixMilli()
	for i := 0; i < 1000000; i++ {
		go sum_co()
	}

	println(time.Now().UnixMilli() - start) // create time

	prev_count := 0
	for prev_count != count {
		println(time.Now().UnixMilli()-start, count)
		prev_count = count
		time.Sleep(10 * time.Millisecond)
	}
	println(time.Now().UnixMilli() - 10 - start) // calc time
	time.Sleep(3000 * time.Millisecond)          // ms
}
