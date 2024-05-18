package main

import (
	"fmt"
)

func sendInts(ch chan int) {
	ch <- 42
	close(ch)
}

func sendStrings(ch chan string) {
	ch <- "Hello"
	ch <- "World"
	close(ch)
}

func main() {
	intCh := make(chan int)
	stringCh := make(chan string)

	go sendInts(intCh)
	go sendStrings(stringCh)

	for i := range intCh {
		fmt.Println("Received int:", i)
	}

	for s := range stringCh {
		fmt.Println("Received string:", s)
	}
}

// package main

// import (
// 	"fmt"
// )

// func sendInts(ch chan int) {
// 	for i := 0; i < 5; i++ {
// 		ch <- i
// 	}
// 	close(ch)
// }

// func main() {
// 	intCh := make(chan int)

// 	go sendInts(intCh)

// 	// Receiving values from the channel using range
// 	for val := range intCh {
// 		fmt.Println("Received int:", val)
// 	}
// }
