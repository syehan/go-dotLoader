package main

import (
	"fmt"
	"time"
)

func main() {
	dot(200 * time.Millisecond)
}

func dot(duration time.Duration) {
	dot := []string{".", "..", "..."}
	for {
		for _, v := range dot {
			fmt.Printf("\r%v ", v)
			time.Sleep(duration)
		}
	}
}
