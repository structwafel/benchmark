package main

/*
#cgo LDFLAGS: -lm
#include <math.h>
*/
import "C"

func main() {
	for i := 0; i < 100000000; i++ {
		_ = C.sqrt(4)
	}
}
