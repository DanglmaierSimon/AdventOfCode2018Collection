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
	var totalsum, sumpart2 int
	//var ch, me, nodeIdx int
	//var remHeaders int
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

	input = strings.Split(input[0], " ")

	totalsum, _ = procInput(input, 0)

	fmt.Printf("Sum: %d\n", totalsum)

}

func procInputpart2(input []string, sum int) (int, []string) {

	ch, me, input := getHeader(input)

	for i := 0; i < ch; i++ {
		s, c := procInput(input, sum)

		sum = s
		input = c
	}

	res, input := getMetaData(me, input)

	for _, i := range res {
		sum += i
	}

	return sum, input
}

func procInput(input []string, sum int) (int, []string) {

	ch, me, input := getHeader(input)

	for i := 0; i < ch; i++ {
		s, c := procInput(input, sum)

		sum = s
		input = c
	}

	res, input := getMetaData(me, input)

	for _, i := range res {
		sum += i
	}

	return sum, input
}

func getHeader(input []string) (int, int, []string) {
	numOfChildren, err := strconv.Atoi(input[0])

	if err != nil {
		fmt.Println("Error converting first number!")
		return 0, 0, nil
	}

	numOfMetaD, err := strconv.Atoi(input[1])

	if err != nil {
		fmt.Println("Error converting second number!")
		return 0, 0, nil
	}

	return numOfChildren, numOfMetaD, input[2:]
}

func getMetaData(count int, input []string) ([]int, []string) {
	if len(input) < count {
		fmt.Println("Error! Trying to extract more Metadata than elements in vector!")
		return nil, nil
	}

	var retVal []int

	for i := 0; i < count; i++ {
		num, err := strconv.Atoi(input[i])

		if err != nil {
			fmt.Println("Error converting metadata to int!")
			return nil, nil
		}

		retVal = append(retVal, num)
	}

	return retVal, input[count:]

}
