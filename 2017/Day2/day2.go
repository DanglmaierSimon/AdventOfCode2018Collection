package main

import "bufio"
import "os"
import "fmt"
import "strings"
import "strconv"

func main() {
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
		diff := 0

		text := scanner.Text()

		fmt.Printf("Line: %s \n", text)

		line := strings.Split(text, "\t")

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
}
