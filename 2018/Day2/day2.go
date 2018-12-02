package main

import (
	"bufio"
	"fmt"
	"os"
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

	resPartOne := partOne(input)
	resPartTwo := partTwo(input)

	fmt.Println("+++RESULT OF PART ONE+++")
	fmt.Printf("The checksum is %d\n", resPartOne)

	fmt.Println("+++RESULT OF PART TWO+++")
	fmt.Printf("The common letters between the packet IDs are: %s\n", resPartTwo)
}

func partOne(input []string) uint {
	var elemTwos, elemThrees uint

	for _, line := range input {
		chars := make(map[rune]uint)

		fmt.Printf("Evaluating new line: %s\n", line)

		for _, char := range line {
			elem, ok := chars[char]

			if ok {
				chars[char] = (elem + 1)
			} else {
				chars[char] = 1
			}
		}

		foundTwo := false
		foundThree := false

		for _, value := range chars {

			if value == 2 && !foundTwo {
				elemTwos++
				foundTwo = true
			} else if value == 3 && !foundThree {
				elemThrees++
				foundThree = true
			}
		}
	}
	fmt.Printf("Number of chars found twice: %d\n", elemTwos)
	fmt.Printf("Number of chars found thrice: %d\n", elemThrees)

	return elemThrees * elemTwos
}

func partTwo(input []string) string {
	var corrPacketID string

	for i := 0; i < len(input)-1; i++ {
		for j := i; j < len(input)-1; j++ {

			for pos := 0; pos < len(input[i])-1; pos++ {
				if input[i][pos] != input[j][pos] {
					test1 := input[i][:pos] + input[i][pos+1:]
					test2 := input[j][:pos] + input[j][pos+1:]

					if test1 == test2 {
						corrPacketID = test1
						break
					}
				}
			}
		}
	}

	return corrPacketID
}
