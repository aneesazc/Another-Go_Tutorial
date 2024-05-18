package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func incrementCounter() {
	for i := 0; i < 3; i++ {
		mutex.Lock()
		{
			counter++
			fmt.Println("Counter incremented to:", counter)
		}
		mutex.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incrementCounter()
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
