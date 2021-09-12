package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 1<<31; i++ {

	}
	duration := time.Since(start)
	fmt.Println(duration)
}
