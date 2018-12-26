package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type rule struct {
	in  string
	out string
}

func main() {
	var input []string
	var rules []rule

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

	initialState := input[0]

	initialState = strings.SplitAfter(initialState, ":")[1]

	initialState = strings.TrimSpace(initialState)

	fmt.Println(initialState, len(initialState))

	for _, l := range input[2:] {
		line := strings.Split(l, "=>")
		line[0] = strings.TrimSpace(line[0])
		line[1] = strings.TrimSpace(line[1])

		rules = append(rules, rule{line[0], line[1]})

	}

	prevState := initialState
	var nextState string

	for i := 0; i < 20; i++ {
		prevState = "....." + prevState + "....."

		for j := 2; j < len(prevState)-2; j++ {
			nextState = nextState + checkRules(rules, prevState[j-2:j+3])
		}

		nextState = strings.TrimRight(nextState, ".")
		nextState = nextState[2:]
		//nextState = nextState[:len(nextState)-2]
		fmt.Println(nextState, len(nextState))
		prevState = nextState
		nextState = ""

	}

	val := -20
	sum := 0

	for _, c := range prevState {
		if c == '#' {
			sum += val
		}

		val++
	}

	fmt.Println("The sum is: ", sum)

}

func checkRules(rules []rule, in string) string {
	var comp string
	switch len(in) {
	case 3:
		comp = ".." + in

	case 4:
		comp = "." + in

	case 5:
		comp = in

	default:
		fmt.Println("Error, received string of invalid length!")
		return "?"
	}

	for _, r := range rules {
		if comp == r.in {
			return r.out
		}
	}

	return "."
}
