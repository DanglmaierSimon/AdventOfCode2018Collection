package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Stack struct {
	idx      int
	children int
	metad    int
	metav    []int
}

func main() {
	var input []string
	//nodeStack := stack.New()

	file, err := os.Open("input_test_1.txt")

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

	ch := input[0]
	me := input[1]

	input = input[2:]

	//for len(input) > 0 {
	fmt.Println(ch)
	fmt.Println(me)
	//}

}
