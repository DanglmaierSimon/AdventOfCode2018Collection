package main

import (
	"bufio"
	"fmt"
	"os"
)

type goblin struct {
	hp int
	x  int
	y  int
}

type elf struct {
	dmg int
	x   int
	y   int
}

type action int

const (
	endTurn action = 0
	attack  action = 1
	move    action = 2
)

func main() {
	var input []string
	//var completedRounds int
	var goblins []goblin
	var elves []elf
	var totalHP int

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

	for i := range input {
		fmt.Println(input[i])
	}

	elves, goblins = initGame(input)

	for _, gob := range goblins {
		totalHP += gob.hp
	}

	fmt.Println(elves)
	fmt.Println(goblins)

	for len(goblins) > 0 && len(elves) > 0 {

	}
}

func initGame(in []string) (elves []elf, goblins []goblin) {

	for x, line := range in {
		for y, val := range line {
			if val == 'G' {
				goblins = append(goblins, goblin{200, x, y})
			} else if val == 'E' {
				elves = append(elves, elf{3, x, y})
			}
		}
	}

	return elves, goblins
}

func doAction(in *[]string, x int, y int, elves *[]elf, goblins *[]goblin) {
	if (*in)[x][y] == 'E' {

	} else if (*in)[x][y] == 'G' {

	} else if ((*in)[x][y] != '.') || ((*in)[x][y] != '#') {
		fmt.Printf("Encountered unexpected symbol: %c", (*in)[x][y])
	}
}
