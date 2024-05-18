package main

import (
	"fmt"
	"sync"
	"unicode"
)

func main() {
	var wg sync.WaitGroup
	data := []rune{'a', 'b', 'c', 'd'}
	capitalized := make([]rune, len(data))
	var mutex sync.Mutex

	capt := func(r rune, index int) {
		defer wg.Done()
		capitalizedRune := unicode.ToUpper(r)
		mutex.Lock()
		capitalized[index] = capitalizedRune
		mutex.Unlock()
		fmt.Printf("%c done!\n", r)
	}

	for i, r := range data {
		wg.Add(1)
		go capt(r, i)
	}

	wg.Wait()
	fmt.Printf("Capitalized: %s\n", string(capitalized))
}
