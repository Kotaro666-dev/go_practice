package main

import "fmt"
import "sync"

func routine(i int, wg *sync.WaitGroup) {
	// Decreasing the counter.
	defer wg.Done()
	fmt.Printf("routine %v finished\n", i)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		// Increasing the counter
		wg.Add(1)
		go routine(i, &wg)
	}
	// Waiting for the counter to reach 0
	wg.Wait()
	fmt.Println("main finished")
}