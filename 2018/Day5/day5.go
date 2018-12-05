package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var diffs uint = 1

	//file, err := os.Open("input_test_1.txt")
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Error: Could not open file!")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	input := scanner.Text()

	fmt.Printf("Input: %s\n", input)

	for diffs > 0 {
		diffs = 0

		for pos := 1; pos < len(input); pos++ {
			if input[pos-1] == (input[pos]+32) || input[pos-1] == (input[pos]-32) {

				fmt.Printf("Deleted %c and %c\n", input[pos-1], input[pos])
				input = input[:pos-1] + input[pos+1:]
				diffs++
			}
		}
	}

	fmt.Printf("Result: %d\n", len(input))
}
