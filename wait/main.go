package main

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

func main() {
	// Poll every 3 seconds with 1 minute timeout.
	count := 0
	wait.Poll(3*time.Second, 1*time.Minute, func() (bool, error) {
		count++
		fmt.Printf("[%d] %s\n", count, time.Now())
		return false, nil
	})
	fmt.Println("Time out!")
}
