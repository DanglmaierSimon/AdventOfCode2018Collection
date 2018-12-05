package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var diffs uint = 1

	var resPartTwo []string

	//file, err := os.Open("input_test_1.txt")
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Error: Could not open file!")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	inputOriginal := scanner.Text()

	input := inputOriginal

	//fmt.Printf("Input: %s\n", input)

	for diffs > 0 {
		diffs = 0

		for pos := 1; pos < len(input); pos++ {
			if input[pos-1] == (input[pos]+32) || input[pos-1] == (input[pos]-32) {

				//fmt.Printf("Deleted %c and %c\n", input[pos-1], input[pos])
				input = input[:pos-1] + input[pos+1:]
				diffs++
			}
		}
	}

	fmt.Printf("Result: %d\n", len(input))

	for ch := 'A'; ch < 'Z'+1; ch++ {
		diffs = 1
		input = inputOriginal

		input = strings.Replace(input, string(ch), "", -1)
		input = strings.Replace(input, string(ch+32), "", -1)

		for diffs > 0 {
			diffs = 0

			for pos := 1; pos < len(input); pos++ {
				if input[pos-1] == (input[pos]+32) || input[pos-1] == (input[pos]-32) {

					//fmt.Printf("Deleted %c and %c\n", input[pos-1], input[pos])
					input = input[:pos-1] + input[pos+1:]
					diffs++
				}
			}
		}

		resPartTwo = append(resPartTwo, input)
	}

	min := len(resPartTwo[0])

	for pos := range resPartTwo {
		if len(resPartTwo[pos]) < min {
			min = len(resPartTwo[pos])
		}
	}

	fmt.Printf("Shortest length is %d\n", min)
}
