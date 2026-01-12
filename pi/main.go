package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("rounds.txt")
	rounds, _ := strconv.Atoi(strings.TrimSpace(string(file)))

	x := 1.0
	pi := 1.0
	stop := rounds + 2

	for i := 2; i <= stop; i++ {
		x = -x
		pi += x / float64(2*i-1)
	}

	pi *= 4.0
	fmt.Println(pi)
}
