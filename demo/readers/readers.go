package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter your name: ")
	// name, _ := r.ReadString('\n')
	// fmt.Printf("Hello, %s", name)

	sum := 0
	for {
		input, inputErr := r.ReadString(' ')
		n := strings.TrimSpace(input)
		if n == "" {
			continue
		}
		num, convErr := strconv.Atoi(n)
		if convErr != nil {
			fmt.Println("Error: ", convErr)
		} else {
			sum += num
		}

		if inputErr == io.EOF {
			break
		}

		if inputErr != nil {
			fmt.Println("Error: ", inputErr)
			break
		}
	}

	fmt.Println("Sum: ", sum)
}
