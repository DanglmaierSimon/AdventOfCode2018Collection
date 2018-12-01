package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input []string

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Error: Could not open file!")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	resultPartOne := partOne(input)
	resultPartTwo := partTwo(input)

	fmt.Printf("The final frequency of part one is: %s\n", resultPartOne)

	fmt.Printf("The final frequency of part two is: %s\n", resultPartTwo)
}

func partOne(args []string) string {
	var finalFreq int

	for _, val := range args {

		freq := val

		freqValue, err := strconv.Atoi(freq[1:])

		if err != nil {
			fmt.Println("Error converting string to number!")
			return ""
		}

		if freq[0] == '-' {
			finalFreq -= freqValue
		} else if freq[0] == '+' {
			finalFreq += freqValue
		} else {
			fmt.Printf("Encountered unknown first symbol!\n")
			fmt.Println(freq)
			return ""
		}
	}

	return strconv.Itoa(finalFreq)
}

func partTwo(args []string) string {
	var previousFreqs map[int]int
	var currFreq int
	var index uint

	previousFreqs = make(map[int]int)

	for {

		freq := args[index]

		freqValue, err := strconv.Atoi(freq[1:])

		if err != nil {
			fmt.Println("Error converting string to number!")
			return ""
		}

		if freq[0] == '-' {
			currFreq -= freqValue
		} else if freq[0] == '+' {
			currFreq += freqValue
		} else {
			fmt.Printf("Encountered unknown first symbol!\n")
			fmt.Println(freq)
			return ""
		}

		_, ok := previousFreqs[currFreq]

		// Value is already in the map
		if ok {
			return strconv.Itoa(currFreq)
		}

		previousFreqs[currFreq] = currFreq

		// Loop until back to the end
		index = (index + 1) % uint(len(args))
	}
}
