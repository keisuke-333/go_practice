package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch1 <- 10
		time.Sleep(500 * time.Millisecond)
	}()
	fmt.Println(<-ch1)
	wg.Wait()

	ch2 := make(chan int, 1)
	ch2 <- 1
	// ch2 <- 2 // deadlock!
	fmt.Println(<-ch2)
}
