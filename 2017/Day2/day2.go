package main

import "bufio"
import "os"
import "fmt"
import "strings"
import "strconv"

func main() {

	var input []string
	var checksum uint64

	//The Spreadsheet is loaded from a file and read line by line
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

	for _, wholeLine := range input {
		diff := 0

		fmt.Printf("Line: %s \n", wholeLine)

		line := strings.Split(wholeLine, "\t")

		min, errmin := strconv.Atoi(line[0])
		max, errmax := strconv.Atoi(line[0])

		if errmax != nil || errmin != nil {
			fmt.Println("Error in String conversion")
			return
		}

		for _, val := range line {

			if len(val) > 0 {

				value, err := strconv.Atoi(val)

				if err != nil {
					fmt.Println("Error converting value to int!")

					fmt.Printf("String: %s \n", val)

					return
				}

				if value < min {
					fmt.Printf("New min value! Old: %d; New:%d\n", min, value)
					min = value
				}

				if value > max {
					fmt.Printf("New max value! Old: %d; New:%d\n", max, value)
					max = value
				}

				diff = max - min
			}
		}

		checksum += uint64(diff)

	}

	fmt.Printf("The final checksum is: %d\n", checksum)

	var sum int

	for _, wholeLine := range input {

		//fmt.Printf("Line: %s \n", wholeLine)

		line := strings.Split(wholeLine, "\t")

		for i := 0; i < len(line); i++ {
			for j := 0; j < len(line); j++ {

				if !(len(line[i]) == 0 || len(line[j]) == 0) {

					num1, err1 := strconv.Atoi(line[i])
					num2, err2 := strconv.Atoi(line[j])

					if err1 != nil || err2 != nil {
						fmt.Println("Something went wrong!")
						return
					}

					//fmt.Printf("Comparing %d and %d\n", num1, num2)

					if num1%num2 == 0 && num1 != num2 {
						fmt.Printf("1: %d divides %d\n", num1, num2)
						sum += num1 / num2
						fmt.Printf("Sum is: %d\n", sum)
					}

				}
			}
		}
	}

	fmt.Printf("Sum for Part 2 is: %d\n", sum)
}
