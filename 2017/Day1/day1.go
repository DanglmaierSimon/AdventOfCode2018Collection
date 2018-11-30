package main

import (
	"fmt"
	"os"
)

func main() {
	var sum uint64
	number := os.Args[1]
	length := len(number)

	for i := 0; i < len(number); i++ {
		if number[i] == number[(i+1)%length] {
			sum += uint64((number[i] - 48))
		}
	}

	fmt.Printf("The sum for part 1 is: %d \n", sum)

	sum = 0

	for i := 0; i < len(number); i++ {
		if number[i] == number[(i+length/2)%length] {
			sum += uint64((number[i] - 48))
		}
	}

	fmt.Printf("The sum for part 2 is: %d \n", sum)
}
