package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type dmgType string

const (
	b dmgType = "bludgeoning"
	c dmgType = "cold"
	f dmgType = "fire"
	s dmgType = "slashing"
	r dmgType = "radiation"
)

type group struct {
	unitCnt int
	hp      int
	init    int
	dmg     int
	dtype   dmgType
	weak    []dmgType
	imm     []dmgType
}

func main() {
	var inputImm, inputInf []string
	var immuneSystem = true

	file, err := os.Open("input_test.txt")

	if err != nil {
		fmt.Println("Error: Could not open file!")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//Split input into 2 separate string arrays, one for infection one for immune system
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Infection") {
			immuneSystem = false
		}

		elem := scanner.Text()

		if len(elem) > 0 {
			if immuneSystem {
				inputImm = append(inputImm, elem)
			} else {
				inputInf = append(inputInf, elem)
			}
		}

	}

	//Remove first element of each array, they contain no useful information
	inputImm = inputImm[1:]
	inputInf = inputInf[1:]

	var immune, infection []group

	for _, val := range inputImm {
		immune = append(immune, parseInput(val))
	}

	for _, val := range inputInf {
		infection = append(immune, parseInput(val))
	}

	fmt.Println(immune)
	fmt.Println(infection)

}

func parseInput(in string) (grp group) {
	res := strings.Split(in, "units")

	res[0] = strings.TrimSpace(res[0])

	units, err := strconv.Atoi(res[0])

	if err != nil {
		fmt.Println("Error converting units to number!")
		return grp
	}

	grp.unitCnt = units

	return grp
}
