package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input []string

	//For testing part 1 a 8x8 is sufficient
	//var fabric [8][8]int

	var fabric [1000][1000]int

	//file, err := os.Open("input_test_1.txt")
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

	resPartOne, fabric := partOne(input, fabric)
	resPartTwo := partTwo(input, fabric)

	fmt.Println("Result Part One:")
	fmt.Printf("Number of Squares that overlap: %d\n", resPartOne)
	fmt.Println("Result Part Two:")
	fmt.Printf("Claim ID: %d\n", resPartTwo)

}

func partOne(input []string, fabric [1000][1000]int) (uint, [1000][1000]int) {
	var overlap uint

	for _, val := range input {
		firstIdx := strings.Index(val, "@")
		secIdx := strings.Index(val, ":")

		claimID := strings.TrimSpace(val[1:firstIdx])
		coords := strings.TrimSpace(val[firstIdx+1 : secIdx])
		size := strings.TrimSpace(val[secIdx+1:])

		idxComma := strings.Index(coords, ",")
		idxX := strings.Index(size, "x")

		coordX, err := strconv.Atoi(coords[:idxComma])

		if err != nil {
			fmt.Println("Error 1")
			return 0, fabric
		}

		coordY, err := strconv.Atoi(coords[idxComma+1:])

		if err != nil {
			fmt.Println("Error 2")
			return 0, fabric
		}

		sizeX, err := strconv.Atoi(size[:idxX])

		if err != nil {
			fmt.Println("Error 3")
			return 0, fabric
		}

		sizeY, err := strconv.Atoi(size[idxX+1:])

		if err != nil {
			fmt.Println("Error 4")
			return 0, fabric
		}

		for x := 0; x < sizeX; x++ {
			for y := 0; y < sizeY; y++ {
				val := fabric[x+coordX][y+coordY]
				id, err := strconv.Atoi(claimID)

				if err != nil {
					fmt.Println("Error 5")
					return 0, fabric
				}

				if val == 0 {
					fabric[x+coordX][y+coordY] = id
				} else {
					fabric[x+coordX][y+coordY] = -1
				}
			}
		}

	}

	for _, row := range fabric {
		for _, val := range row {
			if val == -1 {
				overlap++
			}
		}
	}

	return overlap, fabric
}

func partTwo(input []string, fabric [1000][1000]int) int {

	for _, val := range input {
		firstIdx := strings.Index(val, "@")
		secIdx := strings.Index(val, ":")

		claimID := strings.TrimSpace(val[1:firstIdx])
		coords := strings.TrimSpace(val[firstIdx+1 : secIdx])
		size := strings.TrimSpace(val[secIdx+1:])

		idxComma := strings.Index(coords, ",")
		idxX := strings.Index(size, "x")

		coordX, err := strconv.Atoi(coords[:idxComma])

		if err != nil {
			fmt.Println("Error 1")
			return 0
		}

		coordY, err := strconv.Atoi(coords[idxComma+1:])

		if err != nil {
			fmt.Println("Error 2")
			return 0
		}

		sizeX, err := strconv.Atoi(size[:idxX])

		if err != nil {
			fmt.Println("Error 3")
			return 0
		}

		sizeY, err := strconv.Atoi(size[idxX+1:])

		if err != nil {
			fmt.Println("Error 4")
			return 0
		}

		id, err := strconv.Atoi(claimID)

		if err != nil {
			fmt.Println("Error 5")
			return 0
		}

		var claimSize int

		for x := 0; x < sizeX; x++ {
			for y := 0; y < sizeY; y++ {
				val := fabric[x+coordX][y+coordY]

				if val == id {
					claimSize++
				}
			}
		}

		if claimSize == sizeX*sizeY {
			return id
		}
	}

	return 0
}
