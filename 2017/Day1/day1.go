package main

import (
	"fmt"
	"os"
)

func main() {

	//Sum storing the final result
	var sum uint64

	//Number to be analysed is passed to the program via command line
	number := os.Args[1]
	length := len(number)

	//-----------
	//Part 1
	//-----------
	for i := 0; i < len(number); i++ {
		if number[i] == number[(i+1)%length] {
			sum += uint64((number[i] - 48))
		}
	}

	fmt.Printf("The sum for part 1 is: %d \n", sum)

	//-----------
	//Part 2
	//-----------
	sum = 0

	for i := 0; i < len(number); i++ {
		if number[i] == number[(i+length/2)%length] {
			sum += uint64((number[i] - 48))
		}
	}

	fmt.Printf("The sum for part 2 is: %d \n", sum)
}
