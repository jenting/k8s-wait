package main

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

func main() {
	backoff := wait.Backoff{
		Duration: 3 * time.Second,
		Factor:   0.0,
		Jitter:   0.0,
		Steps:    20,
		Cap:      0,
	}

	// Poll every 3 seconds with 1 minute timeout.
	count := 0
	wait.ExponentialBackoff(backoff, func() (bool, error) {
		count++
		fmt.Printf("[%d] %s\n", count, time.Now())
		return false, nil
	})
	fmt.Println("Time out!")
}
