package main

import (
	"fmt"
)

// import "time"

func main() {
	ch := make(chan int)
	count := 0
	for i := 0; i < 3; i++ {
		go func() {
			count = i
			ch <- count
		}()
	}

	first := <-ch
	second := <-ch
	third := <-ch
	fmt.Println(first, second, third)

}
