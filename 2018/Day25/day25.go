package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type coords struct {
	w int
	x int
	y int
	z int
}

type constellation []coords

func main() {
	var stars map[int]coords
	var cons []constellation
	var input []string
	var starIdx int

	stars = make(map[int]coords)

	//file, err := os.Open("input1.txt")
	//file, err := os.Open("input2.txt")
	//file, err := os.Open("input3.txt")
	//file, err := os.Open("input4.txt")
	file, err := os.Open("input5.txt")

	if err != nil {
		fmt.Println("Error: Could not open file!")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//Split input into 2 separate string arrays, one for infection one for immune system
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	for _, l := range input {
		c := strings.Split(l, ",")

		if len(c) != 4 {
			fmt.Println("Error: Lenth of line is not 4!")
			return
		}

		c1, e1 := strconv.Atoi(c[0])
		c2, e2 := strconv.Atoi(c[1])
		c3, e3 := strconv.Atoi(c[2])
		c4, e4 := strconv.Atoi(c[3])

		if e1 != nil || e2 != nil || e3 != nil || e4 != nil {
			fmt.Println("Error converting coordinates")
			return
		}

		stars[starIdx] = coords{c1, c2, c3, c4}
		starIdx++
	}

	for _, s := range stars {
		fmt.Println(s)
	}

	for len(stars) > 0 {
		var curr constellation
		addedElem := true

		for i, v := range stars {
			curr = append(curr, v)
			delete(stars, i)
			break
		}

		for addedElem {
			addedElem = false
			for _, s1 := range curr {

				for i1, s2 := range stars {
					md := calcManhattanDistance(s2, s1)

					if md < 4 {
						curr = append(curr, s2)
						delete(stars, i1)
						addedElem = true
					}
				}
			}
		}

		cons = append(cons, curr)

		curr = nil
	}

	fmt.Println("Constellations:")
	fmt.Println(cons)
	fmt.Printf("Number of constellations: %d\n", len(cons))
}

func calcManhattanDistance(s1 coords, s2 coords) int {

	d1 := math.Abs(float64(s1.w - s2.w))
	d2 := math.Abs(float64(s1.x - s2.x))
	d3 := math.Abs(float64(s1.y - s2.y))
	d4 := math.Abs(float64(s1.z - s2.z))

	return int(d1 + d2 + d3 + d4)
}
