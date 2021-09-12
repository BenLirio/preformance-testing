package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 1<<22; i++ {

	}
	duration := time.Since(start)
	fmt.Println(duration)
}
